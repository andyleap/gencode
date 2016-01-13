package main

import "io"

type Field interface {
	GenerateSerialize(w io.Writer)
	GenerateDeserialize(w io.Writer)
	GenerateField(w io.Writer)
	SetName(name string)
}

type ResolveField interface {
	Resolve(s *Schema) error
}

type Struct struct {
	Name   string
	Fields []Field
}

type Schema struct {
	Structs []*Struct
}

func (s *Schema) ResolveAll() error {
	for _, st := range s.Structs {
		for _, f := range st.Fields {
			if rf, ok := f.(ResolveField); ok {
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
