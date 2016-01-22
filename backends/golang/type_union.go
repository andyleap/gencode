package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	UnionTemps *template.Template
)

func init() {
	UnionTemps = template.New("UnionTemps")

	template.Must(UnionTemps.New("marshal").Parse(`
	{
		var v uint64
		switch {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			v = {{$id}} + 1
			{{end}}
		}
		{{.VarIntCode}}
		switch tt := {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			{{index $.SubTypeCode $id}}
			{{if gt (index $.SubTypeOffset $id) 0 }}
			i += {{index $.SubTypeOffset $id}}
			{{end}}
			{{end}}
		}
	}`))
	template.Must(UnionTemps.New("unmarshal").Parse(`
	{
		v := uint64(0)
		{{.VarIntCode}}
		switch v {
			{{range $id, $struct := .Structs}}
		case {{$id}} + 1:
			var tt {{index $.SubTypeField $id}}
			{{index $.SubTypeCode $id}}
			{{if gt (index $.SubTypeOffset $id) 0 }}
			i += {{index $.SubTypeOffset $id}}
			{{end}}
			{{$.Target}} = tt
			{{end}}
		default:
			{{.Target}} = nil
		}
	}`))
	template.Must(UnionTemps.New("size").Parse(`
	{
		var v uint64
		switch {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			v = {{$id}} + 1
			{{end}}
		}
		{{.VarIntCode}}
		switch tt := {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			{{index $.SubTypeCode $id}}
			{{if gt (index $.SubTypeOffset $id) 0 }}
			s += {{index $.SubTypeOffset $id}}
			{{end}}
			{{end}}
		}
	}`))
	template.Must(UnionTemps.New("field").Parse(`{{if .Interface}}{{.Interface}}{{else}}interface{}{{end}}`))
}

type UnionTemp struct {
	*schema.UnionType
	Target        string
	VarIntCode    string
	SubTypeCode   []string
	SubTypeField  []string
	SubTypeOffset []int
}

func (w *Walker) WalkUnionDef(ut *schema.UnionType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(UnionTemps, "field", ut)
	return
}

func (w *Walker) WalkUnionSize(ut *schema.UnionType, target string) (parts *StringBuilder, err error) {
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
	subtypecodes := []string{}
	subtypeoffsets := []int{}
	for _, st := range ut.Types {
		offset := w.Offset
		subType, err := w.WalkTypeSize(st, "t")
		if err != nil {
			return nil, err
		}
		SubOffset := w.Offset - offset
		w.Offset = offset
		subtypeoffsets = append(subtypeoffsets, SubOffset)
		subtypecodes = append(subtypecodes, subType.String())
	}

	err = parts.AddTemplate(UnionTemps, "size", UnionTemp{ut, target, intcode.String(), subtypecodes, nil, subtypeoffsets})
	return
}

func (w *Walker) WalkUnionMarshal(ut *schema.UnionType, target string) (parts *StringBuilder, err error) {
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
	subtypecodes := []string{}
	subtypeoffsets := []int{}
	for _, st := range ut.Types {
		offset := w.Offset
		subType, err := w.WalkTypeMarshal(st, "t")
		if err != nil {
			return nil, err
		}
		SubOffset := w.Offset - offset
		w.Offset = offset
		subtypeoffsets = append(subtypeoffsets, SubOffset)
		subtypecodes = append(subtypecodes, subType.String())
	}

	err = parts.AddTemplate(UnionTemps, "marshal", UnionTemp{ut, target, intcode.String(), subtypecodes, nil, subtypeoffsets})
	return
}

func (w *Walker) WalkUnionUnmarshal(ut *schema.UnionType, target string) (parts *StringBuilder, err error) {
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
	subtypecodes := []string{}
	subtypeoffsets := []int{}
	for _, st := range ut.Types {
		offset := w.Offset
		subType, err := w.WalkTypeUnmarshal(st, "t")
		if err != nil {
			return nil, err
		}
		SubOffset := w.Offset - offset
		w.Offset = offset
		subtypeoffsets = append(subtypeoffsets, SubOffset)
		subtypecodes = append(subtypecodes, subType.String())
	}
	subtypefields := []string{}
	for _, st := range ut.Types {
		subType, err := w.WalkTypeDef(st)
		if err != nil {
			return nil, err
		}
		subtypefields = append(subtypefields, subType.String())
	}
	err = parts.AddTemplate(UnionTemps, "unmarshal", UnionTemp{ut, target, intcode.String(), subtypecodes, subtypefields, subtypeoffsets})
	return
}
