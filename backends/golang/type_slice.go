package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
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
		for k := range {{.Target}} {
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
		for k := range {{.Target}} {
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
		for k := range {{.Target}} {
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
	Target      string
	SubTypeCode string
	SubField    string
	VarIntCode  string
}

func WalkSliceDef(st *schema.SliceType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("[]")
	sub, err := WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	parts.Join(sub)
	return
}

func WalkSliceSize(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := WalkIntSize(intHandler, "l")
	if err != nil {
		return nil, err
	}
	subtypecode, err := WalkTypeSize(st.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytesize", SliceTemp{st, target, subtypecode.String(), "", intcode.String()})
	} else {
		err = parts.AddTemplate(SliceTemps, "size", SliceTemp{st, target, subtypecode.String(), "", intcode.String()})
	}
	return
}

func WalkSliceMarshal(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := WalkIntMarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	subtypecode, err := WalkTypeMarshal(st.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	subfield, err := WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytemarshal", SliceTemp{st, target, subtypecode.String(), subfield.String(), intcode.String()})
	} else {
		err = parts.AddTemplate(SliceTemps, "marshal", SliceTemp{st, target, subtypecode.String(), subfield.String(), intcode.String()})
	}
	return
}

func WalkSliceUnmarshal(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := WalkIntUnmarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	subtypecode, err := WalkTypeUnmarshal(st.SubType, target+"[k]")
	if err != nil {
		return nil, err
	}
	subfield, err := WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "byteunmarshal", SliceTemp{st, target, subtypecode.String(), subfield.String(), intcode.String()})
	} else {
		err = parts.AddTemplate(SliceTemps, "unmarshal", SliceTemp{st, target, subtypecode.String(), subfield.String(), intcode.String()})
	}
	return
}
