package main

import (
	"bytes"
	"io"
	"text/template"
)

var (
	StringTemps *template.Template
)

func init() {
	StringTemps = template.Must(template.New("serialize").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		var err error
		if sw, ok := w.(interface{WriteString(s string) (n int, err error);}); ok {
			_, err = sw.WriteString({{.Target}})
		} else {
			_, err = w.Write([]byte({{.Target}}))
		}
		if err != nil {
			return err
		}
	}`))
	template.Must(StringTemps.New("deserialize").Parse(`
	{
		l := uint64(0)
		{{.VarIntCode}}
		sbuf := make([]byte, l)
		var err error
		n := uint64(0)
		for n < l && err == nil {
			var nn int
			nn, err = r.Read(sbuf[n:])
			n += uint64(nn)
		}
		if err != nil {
			return err
		}
		{{.Target}} = *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: uintptr(unsafe.Pointer(&sbuf[0])), Len: int(l)}))
	}`))
	template.Must(StringTemps.New("field").Parse(`string`))
}

type StringType struct {
}

type StringTemp struct {
	StringType
	Target     string
	VarIntCode string
}

func (s StringType) GenerateSerialize(w io.Writer, target string) error {
	intHandler := &VarIntType{
		Bits:   64,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateSerialize(intcode, "l")
	if err != nil {
		return err
	}
	err = StringTemps.ExecuteTemplate(w, "serialize", StringTemp{s, target, string(intcode.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (s StringType) GenerateDeserialize(w io.Writer, target string) error {
	intHandler := &VarIntType{
		Bits:   64,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateDeserialize(intcode, "l")
	if err != nil {
		return err
	}
	err = StringTemps.ExecuteTemplate(w, "deserialize", StringTemp{s, target, string(intcode.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (s StringType) GenerateField(w io.Writer) error {
	err := StringTemps.ExecuteTemplate(w, "field", s)
	if err != nil {
		return err
	}
	return nil
}
