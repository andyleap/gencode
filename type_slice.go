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
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		for k := range {{.Target}} {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(SliceTemps.New("deserialize").Parse(`
	{
		l := uint64(0)
		{{.VarIntCode}}
		if uint64(cap({{.Target}})) >= l {
			{{.Target}} = {{.Target}}[:l]
		} else {
			{{.Target}} = make([]{{.SubField}}, l)
		}
		for k := range {{.Target}} {
			{{.SubTypeCode}}
		}
	}`))
	template.Must(SliceTemps.New("byteserialize").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		_, err := w.Write({{.Target}})
		if err != nil {
			return err
		}
	}`))
	template.Must(SliceTemps.New("bytedeserialize").Parse(`
	{
		l := uint64(0)
		{{.VarIntCode}}
		if uint64(cap({{.Target}})) >= l {
			{{.Target}} = {{.Target}}[:l]
		} else {
			{{.Target}} = make([]{{.SubField}}, l)
		}
		var err error
		n := uint64(0)
		for n < l && err == nil {
			var nn int
			nn, err = r.Read({{.Target}}[n:])
			n += uint64(nn)
		}
		if err != nil {
			return err
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
	VarIntCode  string
}

func (s SliceType) GenerateSerialize(w io.Writer, target string) error {
	intHandler := &VarIntType{
		Bits:   64,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateSerialize(intcode, "l")
	if err != nil {
		return err
	}
	subtype := &bytes.Buffer{}
	err = s.SubType.GenerateSerialize(subtype, target+"[k]")
	if err != nil {
		return err
	}
	if _, ok := s.SubType.(*ByteType); ok {
		err = SliceTemps.ExecuteTemplate(w, "byteserialize", SliceTemp{s, target, string(subtype.Bytes()), "", string(intcode.Bytes())})
	} else {
		err = SliceTemps.ExecuteTemplate(w, "serialize", SliceTemp{s, target, string(subtype.Bytes()), "", string(intcode.Bytes())})
	}
	if err != nil {
		return err
	}
	return nil
}

func (s SliceType) GenerateDeserialize(w io.Writer, target string) error {
	intHandler := &VarIntType{
		Bits:   64,
		Signed: false,
	}
	intcode := &bytes.Buffer{}
	err := intHandler.GenerateDeserialize(intcode, "l")
	if err != nil {
		return err
	}
	subtype := &bytes.Buffer{}
	err = s.SubType.GenerateDeserialize(subtype, target+"[k]")
	if err != nil {
		return err
	}
	subfield := &bytes.Buffer{}
	err = s.SubType.GenerateField(subfield)
	if err != nil {
		return err
	}
	if _, ok := s.SubType.(*ByteType); ok {
		err = SliceTemps.ExecuteTemplate(w, "bytedeserialize", SliceTemp{s, target, string(subtype.Bytes()), string(subfield.Bytes()), string(intcode.Bytes())})
	} else {
		err = SliceTemps.ExecuteTemplate(w, "deserialize", SliceTemp{s, target, string(subtype.Bytes()), string(subfield.Bytes()), string(intcode.Bytes())})
	}
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
