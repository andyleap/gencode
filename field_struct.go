package main

import (
	"io"
	"text/template"
)

var (
	StructTemps *template.Template
)

func init() {
	StructTemps = template.Must(template.New("serialize").Parse(`
	{
		err := d.{{.Name}}.Serialize(w)
		if err != nil {
			return nil
		}
	}`))
	template.Must(StructTemps.New("deserialize").Parse(`
	{
		err := d.{{.Name}}.Deserialize(r)
		if err != nil {
			return nil
		}
	}`))
	template.Must(StructTemps.New("field").Parse(`
	{{.Name}} {{.Struct}}`))
}

type StructField struct {
	Name   string
	Struct string
}

func (i StructField) GenerateSerialize(w io.Writer) {
	StructTemps.ExecuteTemplate(w, "serialize", i)
}

func (i StructField) GenerateDeserialize(w io.Writer) {
	StructTemps.ExecuteTemplate(w, "deserialize", i)
}

func (i StructField) GenerateField(w io.Writer) {
	StructTemps.ExecuteTemplate(w, "field", i)
}

func (i *StructField) SetName(name string) {
	i.Name = name
}
