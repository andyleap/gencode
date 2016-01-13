package main

import (
	"fmt"
	"io"
)

type DeferField struct {
	Name     string
	Defer    string
	Resolved Field
}

type ResolveError struct {
	Defer string
}

func (r ResolveError) Error() string {
	return fmt.Sprintf("Can't resolve %s", r.Defer)
}

func (d DeferField) GenerateSerialize(w io.Writer) {
	d.Resolved.GenerateSerialize(w)
}

func (d DeferField) GenerateDeserialize(w io.Writer) {
	d.Resolved.GenerateDeserialize(w)
}

func (d DeferField) GenerateField(w io.Writer) {
	d.Resolved.GenerateField(w)
}

func (d *DeferField) SetName(name string) {
	d.Name = name
	if d.Resolved != nil {
		d.Resolved.SetName(name)
	}
}

func (d *DeferField) Resolve(s *Schema) error {
	for _, v := range s.Structs {
		if v.Name == d.Defer {
			d.Resolved = &StructField{
				Name:   d.Name,
				Struct: v.Name,
			}
			return nil
		}
	}
	return ResolveError{
		Defer: d.Defer,
	}
}
