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
		}
	}`))
	template.Must(ArrayTemps.New("unmarshal").Parse(`
	{
		for k := range {{.Target}} {
			{{.SubTypeCode}}
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
	Target      string
	SubTypeCode string
	SubField    string
}

func WalkArrayDef(at *schema.ArrayType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf("[%d]", at.Count))
	sub, err := WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	parts.Join(sub)
	return
}

func WalkArraySize(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := WalkTypeSize(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "bytesize", ArrayTemp{at, target, subtypecode.String(), ""})
	} else {
		err = parts.AddTemplate(ArrayTemps, "size", ArrayTemp{at, target, subtypecode.String(), ""})
	}
	return
}

func WalkArrayMarshal(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := WalkTypeMarshal(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	subfield, err := WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "bytemarshal", ArrayTemp{at, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "marshal", ArrayTemp{at, target, subtypecode.String(), subfield.String()})
	}
	return
}

func WalkArrayUnmarshal(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := WalkTypeUnmarshal(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	subfield, err := WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "byteunmarshal", ArrayTemp{at, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "unmarshal", ArrayTemp{at, target, subtypecode.String(), subfield.String()})
	}
	return
}
