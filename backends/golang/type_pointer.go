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
			buf[i] = 0
			i++
		} else {
			buf[i] = 1
			i++
			{{.SubTypeCode}}
		}
	}`))
	template.Must(PointerTemps.New("unmarshal").Parse(`
	{
		if buf[i] == 1 {
			if {{.Target}} == nil {
				{{.Target}} = new({{.SubField}})
			}
			i++
			{{.SubTypeCode}}
		}
	}`))
	template.Must(PointerTemps.New("size").Parse(`
	{
		s++
		if {{.Target}} != nil {
			{{.SubTypeCode}}
		}
	}`))

	template.Must(PointerTemps.New("field").Parse(`*`))
}

type PointerTemp struct {
	*schema.PointerType
	Target      string
	SubTypeCode string
	SubField    string
}

func WalkPointerDef(pt *schema.PointerType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("*")
	sub, err := WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	parts.Join(sub)
	return
}

func WalkPointerSize(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := WalkTypeSize(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(PointerTemps, "size", PointerTemp{pt, target, subtypecode.String(), ""})
	return
}

func WalkPointerMarshal(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := WalkTypeMarshal(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	subfield, err := WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(PointerTemps, "marshal", PointerTemp{pt, target, subtypecode.String(), subfield.String()})
	return
}

func WalkPointerUnmarshal(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := WalkTypeUnmarshal(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	subfield, err := WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(PointerTemps, "unmarshal", PointerTemp{pt, target, subtypecode.String(), subfield.String()})
	return
}
