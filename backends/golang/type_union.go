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
			v = {{$id}}
			{{end}}
		}
		{{.VarIntCode}}
		switch tt := {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			{{index $.SubTypeCode $id}}
			{{end}}
		}
	}`))
	template.Must(UnionTemps.New("unmarshal").Parse(`
	{
		v := uint64(0)
		{{.VarIntCode}}
		switch v {
			{{range $id, $struct := .Structs}}
		case {{$id}}:
			var tt {{index $.SubTypeField $id}}
			{{index $.SubTypeCode $id}}
			{{$.Target}} = tt
			{{end}}
		}
	}`))
	template.Must(UnionTemps.New("size").Parse(`
	{
		var v uint64
		switch {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			v = {{$id}}
			{{end}}
		}
		{{.VarIntCode}}
		switch tt := {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			{{index $.SubTypeCode $id}}
			{{end}}
		}
	}`))
	template.Must(UnionTemps.New("field").Parse(`{{if .Interface}}{{.Interface}}{{else}}interface{}{{end}}`))
}

type UnionTemp struct {
	*schema.UnionType
	Target       string
	VarIntCode   string
	SubTypeCode  []string
	SubTypeField []string
}

func WalkUnionDef(ut *schema.UnionType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(UnionTemps, "field", ut)
	return
}

func WalkUnionSize(ut *schema.UnionType, target string) (parts *StringBuilder, err error) {
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
	subtypecodes := []string{}
	for _, st := range ut.Types {
		subType, err := WalkTypeSize(st, "t")
		if err != nil {
			return nil, err
		}
		subtypecodes = append(subtypecodes, subType.String())
	}

	err = parts.AddTemplate(UnionTemps, "size", UnionTemp{ut, target, intcode.String(), subtypecodes, nil})
	return
}

func WalkUnionMarshal(ut *schema.UnionType, target string) (parts *StringBuilder, err error) {
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
	subtypecodes := []string{}
	for _, st := range ut.Types {
		subType, err := WalkTypeMarshal(st, "t")
		if err != nil {
			return nil, err
		}
		subtypecodes = append(subtypecodes, subType.String())
	}

	err = parts.AddTemplate(UnionTemps, "marshal", UnionTemp{ut, target, intcode.String(), subtypecodes, nil})
	return
}

func WalkUnionUnmarshal(ut *schema.UnionType, target string) (parts *StringBuilder, err error) {
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
	subtypecodes := []string{}
	for _, st := range ut.Types {
		subType, err := WalkTypeUnmarshal(st, "t")
		if err != nil {
			return nil, err
		}
		subtypecodes = append(subtypecodes, subType.String())
	}
	subtypefields := []string{}
	for _, st := range ut.Types {
		subType, err := WalkTypeDef(st)
		if err != nil {
			return nil, err
		}
		subtypefields = append(subtypefields, subType.String())
	}
	err = parts.AddTemplate(UnionTemps, "unmarshal", UnionTemp{ut, target, intcode.String(), subtypecodes, subtypefields})
	return
}
