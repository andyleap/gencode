package main

import (
	"fmt"
	"io"
)

func (f *Field) GenerateSerialize(w io.Writer) error {
	target := fmt.Sprintf("d.%s", f.Name)
	err := f.Type.GenerateSerialize(w, target)
	if err != nil {
		return err
	}
	return nil
}

func (f *Field) GenerateDeserialize(w io.Writer) error {
	target := fmt.Sprintf("d.%s", f.Name)
	err := f.Type.GenerateDeserialize(w, target)
	if err != nil {
		return err
	}
	return nil
}

func (f *Field) GenerateField(w io.Writer) error {
	fmt.Fprintf(w, `
	%s `, f.Name)
	err := f.Type.GenerateField(w)
	if err != nil {
		return err
	}
	return nil
}

func (s *Struct) Generate(w io.Writer) error {
	fmt.Fprintf(w, `type %s struct {`, s.Name)
	for _, v := range s.Fields {
		err := v.GenerateField(w)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(w, `
}

func (d *%s) Serialize(w io.Writer) error {`, s.Name)
	for _, v := range s.Fields {
		err := v.GenerateSerialize(w)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(w, `
	return nil
}

func (d *%s) Deserialize(r io.Reader) error {`, s.Name)
	for _, v := range s.Fields {
		err := v.GenerateDeserialize(w)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(w, `
	return nil
}

`)
	return nil
}

func (s *Schema) Generate(w io.Writer, Package string) error {
	fmt.Fprintf(w, `package %s
import (
	"encoding/binary"
	"io"
)

var (
	_ = io.ReadFull
	_ = binary.Write
)

`, Package)
	for _, st := range s.Structs {
		err := st.Generate(w)
		if err != nil {
			return err
		}
	}
	return nil
}
