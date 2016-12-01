package cpp

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	ArrayTemps *template.Template
)

func init() {
	ArrayTemps = template.New("ArrayTemps")
	template.Must(ArrayTemps.New("marshal").Parse(`
	{
		for (uint64_t k = 0; k < {{.Count}}; k++) {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(ArrayTemps.New("unmarshal").Parse(`
	{
		for (uint64_t k = 0; k < {{.Count}}; k++) {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(ArrayTemps.New("bytemarshal").Parse(`
	{
		memcpy(&buf[i], &{{.Target}}[0], {{.Count}});
		i += {{.Count}};
	}`))
	template.Must(ArrayTemps.New("byteunmarshal").Parse(`
	{
		memcpy(&{{.Target}}[0], &buf[i], {{.Count}});
		i += {{.Count}};
	}`))
	template.Must(ArrayTemps.New("size").Parse(`
	{
		for (uint64_t k = 0; k < {{.Count}}; k++) {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(ArrayTemps.New("bytesize").Parse(`
	{
		s += {{.Count}};
	}`))
	template.Must(ArrayTemps.New("field").Parse(`[{{.Count}}]`))
}

type ArrayTemp struct {
	*schema.ArrayType
	W           *Walker
	Target      string
	SubTypeCode string
	SubField    string
}

func (w *Walker) WalkArraySize(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := w.WalkTypeSize(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "bytesize", ArrayTemp{at, w, target, subtypecode.String(), ""})
	} else {
		err = parts.AddTemplate(ArrayTemps, "size", ArrayTemp{at, w, target, subtypecode.String(), ""})
	}
	return
}

func (w *Walker) WalkArrayMarshal(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := w.WalkTypeMarshal(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "bytemarshal", ArrayTemp{at, w, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "marshal", ArrayTemp{at, w, target, subtypecode.String(), subfield.String()})
	}
	return
}

func (w *Walker) WalkArrayUnmarshal(at *schema.ArrayType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := w.WalkTypeUnmarshal(at.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(at.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := at.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(ArrayTemps, "byteunmarshal", ArrayTemp{at, w, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "unmarshal", ArrayTemp{at, w, target, subtypecode.String(), subfield.String()})
	}
	return
}
