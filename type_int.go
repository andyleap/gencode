package main

import (
	"io"
	"text/template"
)

var (
	IntTemps *template.Template
)

func init() {
	IntTemps = template.New("IntTemps").Funcs(template.FuncMap{
		"Bytes": func(bits int) int {
			return bits / 8
		},
		"BitRange": func(bits int) []int {
			return []int{0, 8, 16, 24, 32, 40, 48, 56, 64}[0:(bits / 8)]
		},
	})

	template.Must(IntTemps.New("serialize").Parse(`
	{
		buf := make([]byte, {{Bytes .Bits}})
		{{range BitRange .Bits}}
		buf[{{Bytes .}}] = byte({{$.Target}} >> {{.}})
		{{end}}
		
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}`))
	template.Must(IntTemps.New("deserialize").Parse(`
	{
		buf := make([]byte, {{Bytes .Bits}})
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		{{range BitRange .Bits}}
		{{$.Target}} |= {{if not $.Signed}}u{{end}}int{{$.Bits}}(buf[{{Bytes .}}]) << {{.}}
		{{end}}
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
