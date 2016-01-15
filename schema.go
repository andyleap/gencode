package main

import "io"

type Type interface {
	GenerateSerialize(w io.Writer, target string) error
	GenerateDeserialize(w io.Writer, target string) error
	GenerateField(w io.Writer) error
}

type ResolveType interface {
	Resolve(s *Schema) error
}

type Field struct {
	Name string
	Type Type
}

type Struct struct {
	Name   string
	Fields []*Field
}

type Schema struct {
	Structs []*Struct
}

func (s *Schema) ResolveAll() error {
	for _, st := range s.Structs {
		for _, f := range st.Fields {
			if rf, ok := f.Type.(ResolveType); ok {
				err := rf.Resolve(s)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

var (
	grammar = MakeGrammar()
)

func ParseSchema(rs io.ReadSeeker) (*Schema, error) {
	s, err := grammar.Parse(rs)
	if err != nil {
		return nil, err
	}
	schema := s.(*Schema)
	err = schema.ResolveAll()
	if err != nil {
		return nil, err
	}
	return schema, nil
}
