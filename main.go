// gencode project main.go
package main

import (
	"flag"
	"log"
	"os"

	"github.com/kr/pretty"
)

var (
	SchemaFile = flag.String("schema", "", "Schema to generate gencode for")
	Package    = flag.String("package", "main", "Package to place code in")
	Verbose    = flag.Bool("verbose", false, "Pretty print the schema for debugging")
)

func main() {
	flag.Parse()

	if *SchemaFile == "" {
		log.Fatal("No file specified")
	}

	file, err := os.Open(*SchemaFile)

	if err != nil {
		log.Fatalf("Error opening schema file: %s", err)
	}

	if file == nil {
		log.Fatalf("error opening file %s", file)
	}

	s, err := ParseSchema(file)

	if *Verbose {
		pretty.Print(s)
	}

	if err != nil {
		log.Fatalf("Error parsing schema: %s", err)
	}

	outfile, err := os.OpenFile(*SchemaFile+".gen.go", os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatalf("Error opening output file: %s", err)
	}

	s.Generate(outfile, *Package)
}
