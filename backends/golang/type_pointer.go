package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	PointerTemps *template.Template
)

func init() {
	PointerTemps = template.New("PointerTemps")
	template.Must(PointerTemps.New("marshal").Parse(`
	{
		if {{.Target}} == nil {
			buf[{{if .W.IAdjusted}}i + {{end}}{{.Offset}}] = 0
		} else {
			buf[{{if .W.IAdjusted}}i + {{end}}{{.Offset}}] = 1
			{{.SubTypeCode}}
			i += {{.SubOffset}}
		}
	}`))
	template.Must(PointerTemps.New("unmarshal").Parse(`
	{
		if buf[{{if .W.IAdjusted}}i + {{end}}{{.Offset}}] == 1 {
			if {{.Target}} == nil {
				{{.Target}} = new({{.SubField}})
			}
			{{.SubTypeCode}}
			i += {{.SubOffset}}
		} else {
			{{.Target}} = nil
		}
	}`))
	template.Must(PointerTemps.New("size").Parse(`
	{
		if {{.Target}} != nil {
			{{.SubTypeCode}}
			s += {{.SubOffset}}
		}
	}`))

	template.Must(PointerTemps.New("field").Parse(`*`))
}

type PointerTemp struct {
	*schema.PointerType
	W           *Walker
	SubOffset   int
	Target      string
	SubTypeCode string
	SubField    string
}

func (w *Walker) WalkPointerDef(pt *schema.PointerType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("*")
	sub, err := w.WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	parts.Join(sub)
	return
}

func (w *Walker) WalkPointerSize(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	Offset := w.Offset
	w.Offset++
	subtypecode, err := w.WalkTypeSize(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - (Offset + 1)
	w.Offset = Offset
	err = parts.AddTemplate(PointerTemps, "size", PointerTemp{pt, w, SubOffset, target, subtypecode.String(), ""})
	w.Offset++
	w.IAdjusted = true
	return
}

func (w *Walker) WalkPointerMarshal(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	Offset := w.Offset
	w.Offset++
	subtypecode, err := w.WalkTypeMarshal(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - (Offset + 1)
	w.Offset = Offset
	err = parts.AddTemplate(PointerTemps, "marshal", PointerTemp{pt, w, SubOffset, target, subtypecode.String(), subfield.String()})
	w.IAdjusted = true
	w.Offset++
	return
}

func (w *Walker) WalkPointerUnmarshal(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	Offset := w.Offset
	w.Offset++
	subtypecode, err := w.WalkTypeUnmarshal(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - (Offset + 1)
	w.Offset = Offset
	err = parts.AddTemplate(PointerTemps, "unmarshal", PointerTemp{pt, w, SubOffset, target, subtypecode.String(), subfield.String()})
	w.IAdjusted = true
	w.Offset++
	return
}
