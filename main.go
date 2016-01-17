// gencode project main.go
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/andyleap/gencode/schema"
	"github.com/kr/pretty"

	_ "github.com/andyleap/gencode/backends/golang"
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

	s, err := schema.ParseSchema(file)

	if *Verbose {
		pretty.Print(s)
	}

	if err != nil {
		log.Fatalf("Error parsing schema: %s", err)
	}

	code, err := schema.Backends["go"].Generate(s)

	if err != nil {
		log.Fatalf("Error generating output: %s", err)
	}

	err = ioutil.WriteFile(*SchemaFile+".gen.go", []byte(code), 0666)

	if err != nil {
		log.Fatalf("Error writing output file: %s", err)
	}

}
