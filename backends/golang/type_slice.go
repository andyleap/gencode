package golang

import (
	"strconv"
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	SliceTemps *template.Template
)

func init() {
	SliceTemps = template.New("SliceTemps")
	template.Must(SliceTemps.New("marshal").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		for {{.Variable}} := range {{.Target}} {
			{{.SubTypeCode}}
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
		for {{.Variable}} := range {{.Target}} {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(SliceTemps.New("bytemarshal").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		copy(buf[i:], {{.Target}})
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
		copy({{.Target}}, buf[i:])
		i += l
	}`))
	template.Must(SliceTemps.New("size").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		for {{.Variable}} := range {{.Target}} {
			_ = {{.Variable}}
			{{.SubTypeCode}}
		}
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
	Target      string
	SubTypeCode string
	SubField    string
	VarIntCode  string
	Variable    string
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
	variable := "k"
	if st.Depth > 0 {
		variable += strconv.Itoa(st.Depth)
	}
	subtypecode, err := w.WalkTypeSize(st.SubType, target+"["+variable+"]")
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytesize", SliceTemp{st, w, target, subtypecode.String(), "", intcode.String(), variable})
	} else {
		err = parts.AddTemplate(SliceTemps, "size", SliceTemp{st, w, target, subtypecode.String(), "", intcode.String(), variable})
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
	variable := "k"
	if st.Depth > 0 {
		variable += strconv.Itoa(st.Depth)
	}
	subtypecode, err := w.WalkTypeMarshal(st.SubType, target+"["+variable+"]")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytemarshal", SliceTemp{st, w, target, subtypecode.String(), subfield.String(), intcode.String(), variable})
	} else {
		err = parts.AddTemplate(SliceTemps, "marshal", SliceTemp{st, w, target, subtypecode.String(), subfield.String(), intcode.String(), variable})
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
	variable := "k"
	if st.Depth > 0 {
		variable += strconv.Itoa(st.Depth)
	}
	subtypecode, err := w.WalkTypeUnmarshal(st.SubType, target+"["+variable+"]")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "byteunmarshal", SliceTemp{st, w, target, subtypecode.String(), subfield.String(), intcode.String(), variable})
	} else {
		err = parts.AddTemplate(SliceTemps, "unmarshal", SliceTemp{st, w, target, subtypecode.String(), subfield.String(), intcode.String(), variable})
	}
	return
}
