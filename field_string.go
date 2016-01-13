package main

import (
	"io"
	"text/template"
)

var (
	StringTemps *template.Template
)

func init() {
	StringTemps = template.Must(template.New("serialize").Parse(`
	{
		t := uint64(len(d.{{.Name}}))
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
		_, err = w.Write([]byte(d.{{.Name}}))
		if err != nil {
			return err
		}
	}`))
	template.Must(StringTemps.New("deserialize").Parse(`
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
		buf = make([]byte, t)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		d.{{.Name}} = string(buf)
	}`))
	template.Must(StringTemps.New("field").Parse(`
	{{.Name}} string`))
}

type StringField struct {
	Name string
}

func (i StringField) GenerateSerialize(w io.Writer) {
	StringTemps.ExecuteTemplate(w, "serialize", i)
}

func (i StringField) GenerateDeserialize(w io.Writer) {
	StringTemps.ExecuteTemplate(w, "deserialize", i)
}

func (i StringField) GenerateField(w io.Writer) {
	StringTemps.ExecuteTemplate(w, "field", i)
}

func (i *StringField) SetName(name string) {
	i.Name = name
}
