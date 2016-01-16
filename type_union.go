package main

import (
	"bytes"
	"io"
	"text/template"
)

var (
	UnionTemps *template.Template
)

func init() {
	UnionTemps = template.Must(template.New("serialize").Parse(`
	{
		var v uint64
		switch {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			v = {{$id}}
			{{end}}
		}
		{{.VarIntCode}}
		var err error
		switch tt := {{.Target}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			err = tt.Serialize(w)
			{{end}}
		}
		if err != nil {
			return err
		}
	}`))
	template.Must(UnionTemps.New("deserialize").Parse(`
	{
		v := uint64(0)
		{{.VarIntCode}}
		switch v {
			{{range $id, $struct := .Structs}}
		case {{$id}}:
			tt := {{$struct.Struct.Name}}{}
			err := tt.Deserialize(r)
			if err != nil {
				return err
			}
			{{$.Target}} = tt
			{{end}}
		}
		
	}`))
	template.Must(UnionTemps.New("field").Parse(`{{if .Interface}}{{.Interface}}{{else}}interface{}{{end}}`))
}

type UnionType struct {
	Structs   []*UnionDefer
	Interface string
}

type UnionDefer struct {
	Defer  string
	Struct *Struct
}

func (ud *UnionDefer) Resolve(s *Schema) error {
	for _, v := range s.Structs {
		if v.Name == ud.Defer {
			ud.Struct = v
			return nil
		}
	}
	return ResolveError{
		Defer: ud.Defer,
	}
}

type UnionTemp struct {
	UnionType
	Target     string
	VarIntCode string
}

func (u UnionType) GenerateSerialize(w io.Writer, target string) error {
	intHandler := &VarIntType{
		Bits:   64,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateSerialize(intcode, "v")
	if err != nil {
		return err
	}
	err = UnionTemps.ExecuteTemplate(w, "serialize", UnionTemp{u, target, string(intcode.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (u UnionType) GenerateDeserialize(w io.Writer, target string) error {
	intHandler := &VarIntType{
		Bits:   64,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateDeserialize(intcode, "v")
	if err != nil {
		return err
	}
	err = UnionTemps.ExecuteTemplate(w, "deserialize", UnionTemp{u, target, string(intcode.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (u UnionType) GenerateField(w io.Writer) error {
	err := UnionTemps.ExecuteTemplate(w, "field", u)
	if err != nil {
		return err
	}
	return nil
}

func (u *UnionType) Resolve(s *Schema) error {
	for _, ud := range u.Structs {
		err := ud.Resolve(s)
		if err != nil {
			return err
		}
	}
	return nil
}
