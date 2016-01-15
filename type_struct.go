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
		err := {{.Target}}.Serialize(w)
		if err != nil {
			return err
		}
	}`))
	template.Must(StructTemps.New("deserialize").Parse(`
	{
		err := {{.Target}}.Deserialize(r)
		if err != nil {
			return err
		}
	}`))
	template.Must(StructTemps.New("field").Parse(`{{.Struct}}`))
}

type StructType struct {
	Struct string
}

type StructTemp struct {
	StructType
	Target string
}

func (s StructType) GenerateSerialize(w io.Writer, target string) error {
	err := StructTemps.ExecuteTemplate(w, "serialize", StructTemp{s, target})
	if err != nil {
		return err
	}
	return nil
}

func (s StructType) GenerateDeserialize(w io.Writer, target string) error {
	err := StructTemps.ExecuteTemplate(w, "deserialize", StructTemp{s, target})
	if err != nil {
		return err
	}
	return nil
}

func (s StructType) GenerateField(w io.Writer) error {
	err := StructTemps.ExecuteTemplate(w, "field", s)
	if err != nil {
		return err
	}
	return nil
}
