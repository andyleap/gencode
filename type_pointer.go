package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"
)

var (
	PointerTemps *template.Template
)

func init() {
	PointerTemps = template.Must(template.New("serialize").Parse(`
	{
		if {{.Target}} == nil {
			_, err := w.Write([]byte{0})
			if err != nil {
				return err
			}
		} else {
			_, err := w.Write([]byte{1})
			if err != nil {
				return err
			}
			{{.SubTypeCode}}
		}
	}`))
	template.Must(PointerTemps.New("deserialize").Parse(`
	{
		buf := make([]byte, 1)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		if buf[0] == 1 {
			if {{.Target}} == nil {
				{{.Target}} = new({{.SubField}})
			}
			{{.SubTypeCode}}
		}
	}`))
	template.Must(PointerTemps.New("field").Parse(`*`))
}

type PointerType struct {
	SubType Type
}

type PointerTemp struct {
	PointerType
	Target      string
	SubTypeCode string
	SubField    string
}

func (p PointerType) GenerateSerialize(w io.Writer, target string) error {
	subtype := &bytes.Buffer{}
	err := p.SubType.GenerateSerialize(subtype, "(*"+target+")")
	if err != nil {
		return err
	}
	err = PointerTemps.ExecuteTemplate(w, "serialize", PointerTemp{p, target, string(subtype.Bytes()), ""})
	if err != nil {
		return err
	}
	return nil
}

func (p PointerType) GenerateDeserialize(w io.Writer, target string) error {
	subtype := &bytes.Buffer{}
	err := p.SubType.GenerateDeserialize(subtype, "(*"+target+")")
	if err != nil {
		return err
	}
	subfield := &bytes.Buffer{}
	err = p.SubType.GenerateField(subfield)
	if err != nil {
		return err
	}
	err = PointerTemps.ExecuteTemplate(w, "deserialize", PointerTemp{p, target, string(subtype.Bytes()), string(subfield.Bytes())})
	if err != nil {
		return err
	}
	return nil
}

func (p PointerType) GenerateField(w io.Writer) error {
	fmt.Fprint(w, "*")
	err := p.SubType.GenerateField(w)
	if err != nil {
		return err
	}
	return nil
}

func (p *PointerType) Resolve(s *Schema) error {
	if rt, ok := p.SubType.(ResolveType); ok {
		err := rt.Resolve(s)
		if err != nil {
			return err
		}
	}
	return nil
}
