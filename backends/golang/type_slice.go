package golang

import (
	"fmt"
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	SliceTemps *template.Template

	// need to make sure we never reuse the same indexer when nesting
	sliceIteratorDepth = 0
)

func init() {
	SliceTemps = template.New("SliceTemps")
	template.Must(SliceTemps.New("marshal").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		for k{{.Index}} := range {{.Target}} {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}
			i += {{.SubOffset}}
			{{end}}
		}
	}`))
	template.Must(SliceTemps.New("unmarshal").Parse(`
	{
		l := uint64(0)
		{{.VarIntCode}}
		if uint64(cap({{.Target}})) >= l {
			{{.Target}} = {{.Target}}[:l]
		} else {
			{{.Target}} = make([]{{.SubField}}, l)
		}
		for k{{.Index}} := range {{.Target}} {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}
			i += {{.SubOffset}}
			{{end}}
		}
	}`))
	template.Must(SliceTemps.New("bytemarshal").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		copy(buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}:], {{.Target}})
		i += l
	}`))
	template.Must(SliceTemps.New("byteunmarshal").Parse(`
	{
		l := uint64(0)
		{{.VarIntCode}}
		if uint64(cap({{.Target}})) >= l {
			{{.Target}} = {{.Target}}[:l]
		} else {
			{{.Target}} = make([]{{.SubField}}, l)
		}
		copy({{.Target}}, buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}:])
		i += l
	}`))
	template.Must(SliceTemps.New("size").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		{{if eq .SubTypeCode "" }}
		{{if gt .SubOffset 0 }}
		s += {{.SubOffset}}*l
		{{end}}
		{{else}}
		for k{{.Index}} := range {{.Target}} {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}
			s += {{.SubOffset}}
			{{end}}
		}
		{{end}}
	}`))
	template.Must(SliceTemps.New("bytesize").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		s += l
	}`))
	template.Must(SliceTemps.New("field").Parse(`[]`))
}

type SliceTemp struct {
	*schema.SliceType
	W           *Walker
	SubOffset   int
	Target      string
	SubTypeCode string
	SubField    string
	VarIntCode  string
	Index       int
}

func (w *Walker) WalkSliceDef(st *schema.SliceType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("[]")
	sub, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	parts.Join(sub)
	return
}

func (w *Walker) WalkSliceSize(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntSize(intHandler, "l")
	if err != nil {
		return nil, err
	}
	offset := w.Offset
	sliceIteratorDepth++
	subtypecode, err := w.WalkTypeSize(st.SubType, fmt.Sprintf("%s[k%d]", target, sliceIteratorDepth-1))
	sliceIteratorDepth--
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytesize", SliceTemp{st, w, SubOffset, target, subtypecode.String(), "", intcode.String(), sliceIteratorDepth})
	} else {
		err = parts.AddTemplate(SliceTemps, "size", SliceTemp{st, w, SubOffset, target, subtypecode.String(), "", intcode.String(), sliceIteratorDepth})
	}
	return
}

func (w *Walker) WalkSliceMarshal(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntMarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	offset := w.Offset
	sliceIteratorDepth++
	subtypecode, err := w.WalkTypeMarshal(st.SubType, fmt.Sprintf("%s[k%d]", target, sliceIteratorDepth-1))
	sliceIteratorDepth--
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	subfield, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytemarshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), sliceIteratorDepth})
	} else {
		err = parts.AddTemplate(SliceTemps, "marshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), sliceIteratorDepth})
	}
	return
}

func (w *Walker) WalkSliceUnmarshal(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntUnmarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	offset := w.Offset
	sliceIteratorDepth++
	subtypecode, err := w.WalkTypeUnmarshal(st.SubType, fmt.Sprintf("%s[k%d]", target, sliceIteratorDepth-1))
	sliceIteratorDepth--
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	subfield, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "byteunmarshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), sliceIteratorDepth})
	} else {
		err = parts.AddTemplate(SliceTemps, "unmarshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), sliceIteratorDepth})
	}
	return
}
