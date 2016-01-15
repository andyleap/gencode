package main

import (
	"fmt"
	"io"
)

type DeferType struct {
	Defer    string
	Resolved Type
}

type ResolveError struct {
	Defer string
}

func (r ResolveError) Error() string {
	return fmt.Sprintf("Can't resolve %s", r.Defer)
}

func (d DeferType) GenerateSerialize(w io.Writer, target string) error {
	err := d.Resolved.GenerateSerialize(w, target)
	if err != nil {
		return err
	}
	return nil
}

func (d DeferType) GenerateDeserialize(w io.Writer, target string) error {
	err := d.Resolved.GenerateDeserialize(w, target)
	if err != nil {
		return err
	}
	return nil
}

func (d DeferType) GenerateField(w io.Writer) error {
	err := d.Resolved.GenerateField(w)
	if err != nil {
		return err
	}
	return nil
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
