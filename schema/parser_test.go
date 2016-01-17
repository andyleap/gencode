package schema

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
struct person {
	name string
	height float64
	age int8
}

struct committee {
	name string
	leader person
}

struct group {
	name string
	leader union {
		person
		committee
	}
}
`)
	if err != nil {
		t.Fatal(err)
	}
	_ = s.(*Schema)

	//	def.Generate(os.Stdout, "test")
}
