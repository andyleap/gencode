package schema

import (
	"fmt"
	"io"
)

type Backend struct {
	Generate func(*Schema) (string, error)
}

var (
	Backends = make(map[string]*Backend)
)

func Register(name string, backend *Backend) {
	Backends[name] = backend
}

type Type interface{}

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
	Framed bool
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

type ResolveError struct {
	Defer string
}

func (r ResolveError) Error() string {
	return fmt.Sprintf("Can't resolve %s", r.Defer)
}

type ArrayType struct {
	SubType Type
	Count   uint64
}

func (at *ArrayType) Resolve(s *Schema) error {
	if rt, ok := at.SubType.(ResolveType); ok {
		err := rt.Resolve(s)
		if err != nil {
			return err
		}
	}
	return nil
}

type BoolType struct {
}

type ByteType struct {
}

type IntType struct {
	Bits   int
	Signed bool
	VarInt bool
}

type DeferType struct {
	Defer    string
	Resolved Type
}

func (d *DeferType) Resolve(s *Schema) error {
	for _, v := range s.Structs {
		if v.Name == d.Defer {
			d.Resolved = &StructType{
				Struct: v.Name,
			}
			return nil
		}
	}
	return ResolveError{
		Defer: d.Defer,
	}
}

type PointerType struct {
	SubType Type
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

type SliceType struct {
	SubType Type
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

type StringType struct {
}

type StructType struct {
	Struct string
}

type UnionType struct {
	Types     []Type
	Interface string
}

func (u *UnionType) Resolve(s *Schema) error {
	for _, ud := range u.Types {
		if rt, ok := ud.(ResolveType); ok {
			err := rt.Resolve(s)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type FloatType struct {
	Bits int
}
