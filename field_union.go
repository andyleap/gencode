package main

import (
	"io"
	"text/template"
)

var (
	UnionTemps *template.Template
)

func init() {
	UnionTemps = template.Must(template.New("serialize").Parse(`
	{
		var t uint64
		switch d.{{.Name}}.(type) {
			{{range $id, $struct := .Structs}}
		case {{$struct.Struct.Name}}:
			t = {{$id}}
			{{end}}
		}
		buf := make([]byte, 10)
		i := 0
		for t >= 0x80 {
			buf[i] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i] = byte(t)
		i++
		_, err := w.Write(buf[:i])
		if err != nil {
			return err
		}
		switch tt := d.{{.Name}}.(type) {
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
		buf := make([]byte, 1)
		buf[0] = 0x80
		t := uint64(0)
		for buf[0] & 0x80 == 0x80 {
			t <<= 7
			_, err := io.ReadFull(r, buf)
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F)
		}
		switch t {
			{{range $id, $struct := .Structs}}
		case {{$id}}:
			tt := {{$struct.Struct.Name}}{}
			err := tt.Deserialize(r)
			if err != nil {
				return err
			}
			d.{{$.Name}} = tt
			{{end}}
		}
		
	}`))
	template.Must(UnionTemps.New("field").Parse(`
	{{.Name}} {{if .Interface}}{{.Interface}}{{else}}interface{}{{end}}`))
}

type UnionField struct {
	Name      string
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

func (u UnionField) GenerateSerialize(w io.Writer) {
	UnionTemps.ExecuteTemplate(w, "serialize", u)
}

func (u UnionField) GenerateDeserialize(w io.Writer) {
	UnionTemps.ExecuteTemplate(w, "deserialize", u)
}

func (u UnionField) GenerateField(w io.Writer) {
	UnionTemps.ExecuteTemplate(w, "field", u)
}

func (u *UnionField) SetName(name string) {
	u.Name = name
}

func (u *UnionField) Resolve(s *Schema) error {
	for _, ud := range u.Structs {
		err := ud.Resolve(s)
		if err != nil {
			return err
		}
	}
	return nil
}
