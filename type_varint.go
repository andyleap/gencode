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
		t := uint{.Bits}({{.Target}})
		{{if .Signed}}
		t <<= 1
   		if {{.Target}} < 0 {
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
		{{.Target}} = int{{.Bits}}(t >> 1)
		if t&1 != 0 {
			{{.Target}} = ^d.{{.Name}}
		}
		{{else}}
		{{.Target}} = t
		{{end}}
	}`))
	template.Must(VarIntTemps.New("field").Parse(`{{if not .Signed}}u{{end}}int{{.Bits}}`))
}

type VarIntType struct {
	Bits   int
	Signed bool
}

type VarIntTemp struct {
	VarIntType
	Target string
}

func (v VarIntType) GenerateSerialize(w io.Writer, target string) error {
	err := VarIntTemps.ExecuteTemplate(w, "serialize", VarIntTemp{v, target})
	if err != nil {
		return err
	}
	return nil
}

func (v VarIntType) GenerateDeserialize(w io.Writer, target string) error {
	err := VarIntTemps.ExecuteTemplate(w, "deserialize", VarIntTemp{v, target})
	if err != nil {
		return err
	}
	return nil
}

func (v VarIntType) GenerateField(w io.Writer) error {
	err := VarIntTemps.ExecuteTemplate(w, "field", v)
	if err != nil {
		return err
	}
	return nil
}
