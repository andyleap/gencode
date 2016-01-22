package golang

import (
	"fmt"
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	ArrayTemps *template.Template
)

func init() {
	ArrayTemps = template.New("ArrayTemps")
	template.Must(ArrayTemps.New("marshal").Parse(`
	{
		for k := range {{.Target}} {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}
			i += {{.SubOffset}}
			{{end}}
		}
	}`))
	template.Must(ArrayTemps.New("unmarshal").Parse(`
	{
		for k := range {{.Target}} {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}
			i += {{.SubOffset}}
			{{end}}
		}
	}`))
	template.Must(ArrayTemps.New("bytemarshal").Parse(`
	{
		copy(buf[i:], {{.Target}}[:])
		i += {{.Count}}
	}`))
	template.Must(ArrayTemps.New("byteunmarshal").Parse(`
	{
		copy({{.Target}}[:], buf[i:])
		i += {{.Count}}
	}`))
	template.Must(ArrayTemps.New("size").Parse(`
	{
		for k := range {{.Target}} {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}
			s += {{.SubOffset}}
			{{end}}
		}
	}`))
	template.Must(ArrayTemps.New("bytesize").Parse(`
	{
		s += {{.Count}}
	}`))
	template.Must(ArrayTemps.New("field").Parse(`[{{.Count}}]`))
}

type ArrayTemp struct {
	*schema.ArrayType
	SubOffset   int
	Target      string
	SubTypeCode string
	SubField    string
}

func (w *Walker) WalkArrayDef(at *schema.ArrayType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf("[%d]", at.Count))
	sub, err := w.WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	parts.Join(sub)
	return
}

func (w *Walker) WalkArraySize(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	w.IAdjusted = true
	offset := w.Offset
	subtypecode, err := w.WalkTypeSize(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "bytesize", ArrayTemp{at, SubOffset, target, subtypecode.String(), ""})
	} else {
		err = parts.AddTemplate(ArrayTemps, "size", ArrayTemp{at, SubOffset, target, subtypecode.String(), ""})
	}
	return
}

func (w *Walker) WalkArrayMarshal(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	w.IAdjusted = true
	offset := w.Offset
	subtypecode, err := w.WalkTypeMarshal(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	subfield, err := w.WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "bytemarshal", ArrayTemp{at, SubOffset, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "marshal", ArrayTemp{at, SubOffset, target, subtypecode.String(), subfield.String()})
	}
	return
}

func (w *Walker) WalkArrayUnmarshal(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	w.IAdjusted = true
	offset := w.Offset
	subtypecode, err := w.WalkTypeUnmarshal(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	subfield, err := w.WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "byteunmarshal", ArrayTemp{at, SubOffset, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "unmarshal", ArrayTemp{at, SubOffset, target, subtypecode.String(), subfield.String()})
	}
	return
}
