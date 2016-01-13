package main

import "io"

type Field interface {
	GenerateSerialize(w io.Writer)
	GenerateDeserialize(w io.Writer)
	GenerateField(w io.Writer)
	SetName(name string)
}

type Struct struct {
	Name   string
	Fields []Field
}

type Schema struct {
	Structs []Struct
}
