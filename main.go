// gencode project main.go
package main

import (
	"flag"
	"log"
	"os"
)

var (
	SchemaFile = flag.String("schema", "", "Schema to generate gencode for")
	Package    = flag.String("package", "main", "Package to place code in")
)

func main() {
	g := MakeGrammar()

	flag.Parse()

	if *SchemaFile == "" {
		log.Fatal("No file specified")
	}

	file, err := os.Open(*SchemaFile)

	if err != nil {
		log.Fatalf("Error opening schema file: %s", err)
	}

	outfile, err := os.OpenFile(*SchemaFile+".gen.go", os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatalf("Error opening output file file: %s", err)
	}

	s, err := g.Parse(file)

	if err != nil {
		log.Fatalf("Error parsing schema file: %s", err)
	}

	def := s.(*Schema)

	def.Generate(outfile, *Package)
}
