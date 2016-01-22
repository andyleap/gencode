// gencode project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/andyleap/gencode/schema"
	"github.com/kr/pretty"

	_ "github.com/andyleap/gencode/backends/golang"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	backend, ok := schema.Backends[os.Args[1]]
	if !ok {
		fmt.Fprintln(os.Stderr, "No such backend, available backends are:")
		for name := range schema.Backends {
			fmt.Fprintln(os.Stderr, name)
		}
		os.Exit(1)
	}

	flags := backend.Flags()

	SchemaFile := ""
	Debug := false

	flags.StringVar(&SchemaFile, "schema", "", "Schema file to process")
	flags.BoolVar(&Debug, "debug", false, "Pretty print the resulting schema defs")

	flags.Parse(os.Args[2:])

	if SchemaFile == "" {
		log.Fatal("No file specified")
	}

	file, err := os.Open(SchemaFile)

	if err != nil {
		log.Fatalf("Error opening schema file: %s", err)
	}

	if file == nil {
		log.Fatalf("error opening file %s", file)
	}

	s, err := schema.ParseSchema(file)

	if Debug {
		pretty.Print(s)
	}

	if err != nil {
		log.Fatalf("Error parsing schema: %s", err)
	}

	code, err := backend.Generate(s)

	if err != nil {
		log.Fatalf("Error generating output: %s", err)
	}

	err = ioutil.WriteFile(backend.GeneratedFilename(SchemaFile), []byte(code), 0666)

	if err != nil {
		log.Fatalf("Error writing output file: %s", err)
	}

}
