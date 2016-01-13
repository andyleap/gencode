package main

import (
	"io"
	"text/template"
)

var (
	VarIntTemps *template.Template
)

func init() {
	VarIntTemps = template.Must(template.New("serialize").Parse(`
	{
		t := uint{.Bits}(d.{{.Name}})
		{{if .Signed}}
		t <<= 1
   		if d.{{.Name}} < 0 {
   			t = ^t
   		}
		{{end}}
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
	}`))
	template.Must(VarIntTemps.New("deserialize").Parse(`
	{
		buf := make([]byte, 1)
		buf[0] = 0x80
		t := uint{{.Bits}}(0)
		for buf & 0x80 == 0x80 {
			t <<= 7
			_, err = io.ReadFull(r, buf)
			if err != nil {
				return err
			}
			t |= int(buf[0]&0x7F)
		}
		{{if .Signed}}
		d.{{.Name}} = int{{.Bits}}(t >> 1)
		if t&1 != 0 {
			d.{{.Name}} = ^d.{{.Name}}
		}
		{{else}}
		d.{{.Name}} = t
		{{end}}
	}`))
	template.Must(VarIntTemps.New("field").Parse(`
	{{.Name}} {{if not .Signed}}u{{end}}int{{.Bits}}`))
}

type VarIntField struct {
	Name   string
	Bits   int
	Signed bool
}

func (i VarIntField) GenerateSerialize(w io.Writer) {
	VarIntTemps.ExecuteTemplate(w, "serialize", i)
}

func (i VarIntField) GenerateDeserialize(w io.Writer) {
	VarIntTemps.ExecuteTemplate(w, "deserialize", i)
}

func (i VarIntField) GenerateField(w io.Writer) {
	VarIntTemps.ExecuteTemplate(w, "field", i)
}

func (i *VarIntField) SetName(name string) {
	i.Name = name
}
