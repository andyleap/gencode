package main

import (
	"io"
	"text/template"
)

var (
	ByteTemps *template.Template
)

func init() {
	ByteTemps = template.New("ByteTemps")

	template.Must(ByteTemps.New("serialize").Parse(`
	{
		_, err := w.Write([]byte{ {{.Target}} })
		if err != nil {
			return err
		}
	}`))
	template.Must(ByteTemps.New("deserialize").Parse(`
	{
		buf := make([]byte, 1)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		{{.Target}} = buf[0]
	}`))
	template.Must(ByteTemps.New("field").Parse(`byte`))
}

type ByteType struct {
}

type ByteTemp struct {
	ByteType
	Target string
}

func (b ByteType) GenerateSerialize(w io.Writer, target string) error {
	err := ByteTemps.ExecuteTemplate(w, "serialize", ByteTemp{b, target})
	if err != nil {
		return err
	}
	return nil
}

func (b ByteType) GenerateDeserialize(w io.Writer, target string) error {
	err := ByteTemps.ExecuteTemplate(w, "deserialize", ByteTemp{b, target})
	if err != nil {
		return err
	}
	return nil
}

func (b ByteType) GenerateField(w io.Writer) error {
	err := ByteTemps.ExecuteTemplate(w, "field", b)
	if err != nil {
		return err
	}
	return nil
}
