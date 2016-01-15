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
		t := uint64(len({{.Target}}))
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
		_, err = w.Write([]byte({{.Target}}))
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
		{{.Target}} = string(buf)
	}`))
	template.Must(StringTemps.New("field").Parse(`string`))
}

type StringType struct {
}

type StringTemp struct {
	StringType
	Target string
}

func (s StringType) GenerateSerialize(w io.Writer, target string) error {
	err := StringTemps.ExecuteTemplate(w, "serialize", StringTemp{s, target})
	if err != nil {
		return err
	}
	return nil
}

func (s StringType) GenerateDeserialize(w io.Writer, target string) error {
	err := StringTemps.ExecuteTemplate(w, "deserialize", StringTemp{s, target})
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
