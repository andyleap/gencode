package main

import (
	"io"
	"text/template"
)

var (
	FloatTemps *template.Template
)

func init() {
	FloatTemps = template.Must(template.New("serialize").Parse(`
	{
		err := binary.Write(w, binary.LittleEndian, {{.Target}})
		if err != nil {
			return err
		}
	}`))
	template.Must(FloatTemps.New("deserialize").Parse(`
	{
		err := binary.Read(r, binary.LittleEndian, &{{.Target}})
		if err != nil {
			return err
		}
	}`))
	template.Must(FloatTemps.New("field").Parse(`float{{.Bits}}`))
}

type FloatType struct {
	Bits int
}

type FloatTemp struct {
	FloatType
	Target string
}

func (i FloatType) GenerateSerialize(w io.Writer, target string) error {
	err := FloatTemps.ExecuteTemplate(w, "serialize", FloatTemp{i, target})
	if err != nil {
		return err
	}
	return nil
}

func (i FloatType) GenerateDeserialize(w io.Writer, target string) error {
	err := FloatTemps.ExecuteTemplate(w, "deserialize", FloatTemp{i, target})
	if err != nil {
		return err
	}
	return nil
}

func (i FloatType) GenerateField(w io.Writer) error {
	err := FloatTemps.ExecuteTemplate(w, "field", i)
	if err != nil {
		return err
	}
	return nil
}
