package main

import (
	"io"
	"text/template"
)

var (
	IntTemps *template.Template
)

func init() {
	IntTemps = template.Must(template.New("serialize").Parse(`
	{
		err := binary.Write(w, binary.LittleEndian, d.{{.Name}})
		if err != nil {
			return err
		}
	}`))
	template.Must(IntTemps.New("deserialize").Parse(`
	{
		err := binary.Read(r, binary.LittleEndian, &d.{{.Name}})
		if err != nil {
			return err
		}
	}`))
	template.Must(IntTemps.New("field").Parse(`
	{{.Name}} {{if not .Signed}}u{{end}}int{{.Bits}}`))
}

type IntField struct {
	Name   string
	Bits   int
	Signed bool
}

func (i IntField) GenerateSerialize(w io.Writer) {
	IntTemps.ExecuteTemplate(w, "serialize", i)
}

func (i IntField) GenerateDeserialize(w io.Writer) {
	IntTemps.ExecuteTemplate(w, "deserialize", i)
}

func (i IntField) GenerateField(w io.Writer) {
	IntTemps.ExecuteTemplate(w, "field", i)
}

func (i *IntField) SetName(name string) {
	i.Name = name
}
