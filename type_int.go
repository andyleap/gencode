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
		err := binary.Write(w, binary.LittleEndian, {{.Target}})
		if err != nil {
			return err
		}
	}`))
	template.Must(IntTemps.New("deserialize").Parse(`
	{
		err := binary.Read(r, binary.LittleEndian, &{{.Target}})
		if err != nil {
			return err
		}
	}`))
	template.Must(IntTemps.New("field").Parse(`{{if not .Signed}}u{{end}}int{{.Bits}}`))
}

type IntType struct {
	Bits   int
	Signed bool
}

type IntTemp struct {
	IntType
	Target string
}

func (i IntType) GenerateSerialize(w io.Writer, target string) error {
	err := IntTemps.ExecuteTemplate(w, "serialize", IntTemp{i, target})
	if err != nil {
		return err
	}
	return nil
}

func (i IntType) GenerateDeserialize(w io.Writer, target string) error {
	err := IntTemps.ExecuteTemplate(w, "deserialize", IntTemp{i, target})
	if err != nil {
		return err
	}
	return nil
}

func (i IntType) GenerateField(w io.Writer) error {
	err := IntTemps.ExecuteTemplate(w, "field", i)
	if err != nil {
		return err
	}
	return nil
}
