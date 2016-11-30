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
			{{if gt .SubOffset 0 }}i += {{.SubOffset}};{{end}}
		}
	}`))
	template.Must(ArrayTemps.New("unmarshal").Parse(`
	{
		for (uint64_t k = 0; k < {{.Count}}; k++) {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}i += {{.SubOffset}};{{end}}
		}
	}`))
	template.Must(ArrayTemps.New("bytemarshal").Parse(`
	{
		memcpy(&buf[{{if $.W.IAdjusted}}i + {{end}}{{$.W.Offset}}], &{{.Target}}[0], {{.Count}});
		i += {{.Count}};
	}`))
	template.Must(ArrayTemps.New("byteunmarshal").Parse(`
	{
		memcpy(&{{.Target}}[0], &buf[{{if $.W.IAdjusted}}i + {{end}}{{$.W.Offset}}], {{.Count}});
		i += {{.Count}};
	}`))
	template.Must(ArrayTemps.New("size").Parse(`
	{
		for (uint64_t k = 0; k < {{.Count}}; k++) {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}s += {{.SubOffset}};{{end}}
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
	SubOffset   int
	Target      string
	SubTypeCode string
	SubField    string
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
		err = parts.AddTemplate(ArrayTemps, "bytesize", ArrayTemp{at, w, SubOffset, target, subtypecode.String(), ""})
	} else {
		err = parts.AddTemplate(ArrayTemps, "size", ArrayTemp{at, w, SubOffset, target, subtypecode.String(), ""})
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
		err = parts.AddTemplate(ArrayTemps, "bytemarshal", ArrayTemp{at, w, SubOffset, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "marshal", ArrayTemp{at, w, SubOffset, target, subtypecode.String(), subfield.String()})
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
		err = parts.AddTemplate(ArrayTemps, "byteunmarshal", ArrayTemp{at, w, SubOffset, target, subtypecode.String(), subfield.String()})
	} else {
		err = parts.AddTemplate(ArrayTemps, "unmarshal", ArrayTemp{at, w, SubOffset, target, subtypecode.String(), subfield.String()})
	}
	return
}
