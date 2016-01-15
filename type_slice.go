package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

var (
	SliceTemps *template.Template
)

func init() {
	SliceTemps = template.Must(template.New("serialize").Parse(`
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
		for _, s := range {{.Target}} {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(SliceTemps.New("deserialize").Parse(`
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
		{{.Target}} = make([]{{.SubField}}, t)
		for k := range {{.Target}} {
			var s {{.SubField}}
			{{.SubTypeCode}}
			{{.Target}}[k] = s
		}
	}`))
	template.Must(SliceTemps.New("field").Parse(`[]`))
}

type SliceType struct {
	SubType Type
}

type SliceTemp struct {
	SliceType
	Target      string
	SubTypeCode string
	SubField    string
}

func (s SliceType) GenerateSerialize(w io.Writer, target string) error {
	subtype := &bytes.Buffer{}
	err := s.SubType.GenerateSerialize(subtype, "s")
	if err != nil {
		return err
	}
	err = SliceTemps.ExecuteTemplate(w, "serialize", SliceTemp{s, target, string(subtype.Bytes()), ""})
	if err != nil {
		return err
	}
	return nil
}

func (s SliceType) GenerateDeserialize(w io.Writer, target string) error {
	subtype := &bytes.Buffer{}
	err := s.SubType.GenerateDeserialize(subtype, "s")
	if err != nil {
		return err
	}
	subfield := &bytes.Buffer{}
	err = s.SubType.GenerateField(subfield)
	if err != nil {
		return err
	}
	err = SliceTemps.ExecuteTemplate(w, "deserialize", SliceTemp{s, target, string(subtype.Bytes()), string(subfield.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (s SliceType) GenerateField(w io.Writer) error {
	fmt.Fprint(w, "[]")
	err := s.SubType.GenerateField(w)
	if err != nil {
		return err
	}
	return nil
}

func (st *SliceType) Resolve(s *Schema) error {
	if rt, ok := st.SubType.(ResolveType); ok {
		err := rt.Resolve(s)
		if err != nil {
			return err
		}
	}
	return nil
}
