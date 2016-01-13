package main

import (
	"fmt"
	"io"
)

func (s *Struct) Generate(w io.Writer) {
	fmt.Fprintf(w, `type %s struct {`, s.Name)
	for _, v := range s.Fields {
		v.GenerateField(w)
	}
	fmt.Fprintf(w, `
}

func (d *%s) Serialize(w io.Writer) error {`, s.Name)
	for _, v := range s.Fields {
		v.GenerateSerialize(w)
	}
	fmt.Fprintf(w, `
	return nil
}

func (d *%s) Deserialize(r io.Reader) error {`, s.Name)
	for _, v := range s.Fields {
		v.GenerateDeserialize(w)
	}
	fmt.Fprintf(w, `
	return nil
}

`)
}

func (s *Schema) Generate(w io.Writer, Package string) {
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
		st.Generate(w)
	}

}
