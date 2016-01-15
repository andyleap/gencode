package main

import (
	"bytes"
	"io"
	"text/template"
)

var (
	FloatTemps *template.Template
)

func init() {
	FloatTemps = template.Must(template.New("serialize").Parse(`
	{
		v := math.Float{{.Bits}}bits({{.Target}})
		{{.IntCode}}
	}`))
	template.Must(FloatTemps.New("deserialize").Parse(`
	{
		var v uint{{.Bits}}
		{{.IntCode}}
		{{.Target}} = math.Float{{.Bits}}frombits(v)
	}`))
	template.Must(FloatTemps.New("field").Parse(`float{{.Bits}}`))
}

type FloatType struct {
	Bits int
}

type FloatTemp struct {
	FloatType
	Target  string
	IntCode string
}

func (i FloatType) GenerateSerialize(w io.Writer, target string) error {
	intHandler := &IntType{
		Bits:   i.Bits,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateSerialize(intcode, "v")
	if err != nil {
		return err
	}
	err = FloatTemps.ExecuteTemplate(w, "serialize", FloatTemp{i, target, string(intcode.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (i FloatType) GenerateDeserialize(w io.Writer, target string) error {
	intHandler := &IntType{
		Bits:   i.Bits,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateDeserialize(intcode, "v")
	if err != nil {
		return err
	}
	err = FloatTemps.ExecuteTemplate(w, "deserialize", FloatTemp{i, target, string(intcode.Bytes())})
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
