package main

import "testing"

func TestParser(t *testing.T) {
	g := MakeGrammar()

	s, err := g.ParseString(`
struct Person {
	name string
	age int32
}`)
	if err != nil {
		t.Fatal(err)
	}
	_ = s.(*Schema)

	//	def.Generate(os.Stdout, "test")
}

func TestParser2(t *testing.T) {
	g := MakeGrammar()

	s, err := g.ParseString(`
struct Person {
	name string
	age int32
}

struct PersonV2 {
	FirstName string
	LastName string
	age int32
}
`)
	if err != nil {
		t.Fatal(err)
	}
	_ = s.(*Schema)

	//	def.Generate(os.Stdout, "test")
}
