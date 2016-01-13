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
		err := binary.Write(w, binary.LittleEndian, d.{{.Name}})
		if err != nil {
			return err
		}
	}`))
	template.Must(FloatTemps.New("deserialize").Parse(`
	{
		err := binary.Read(r, binary.LittleEndian, &d.{{.Name}})
		if err != nil {
			return err
		}
	}`))
	template.Must(FloatTemps.New("field").Parse(`
	{{.Name}} float{{.Bits}}`))
}

type FloatField struct {
	Name string
	Bits int
}

func (i FloatField) GenerateSerialize(w io.Writer) {
	FloatTemps.ExecuteTemplate(w, "serialize", i)
}

func (i FloatField) GenerateDeserialize(w io.Writer) {
	FloatTemps.ExecuteTemplate(w, "deserialize", i)
}

func (i FloatField) GenerateField(w io.Writer) {
	FloatTemps.ExecuteTemplate(w, "field", i)
}

func (i *FloatField) SetName(name string) {
	i.Name = name
}
