package cpp

import (
	"flag"

	"github.com/eyrie-io/gencode/schema"
)

type CppBackend struct {
	Namespace string
}

func (cb *CppBackend) Generate(s *schema.Schema) (string, error) {
	w := &Walker{}
	def, err := w.WalkSchema(s, cb.Namespace)
	if err != nil {
		return "", err
	}
	return def.String(), nil
}

func (cb *CppBackend) Flags() *flag.FlagSet {
	flags := flag.NewFlagSet("Cpp", flag.ExitOnError)
	flags.StringVar(&cb.Namespace, "namespace", "main", "namespace to build the gencode system for")
	return flags
}

func (cb *CppBackend) GeneratedFilename(filename string) string {
	return filename + ".gen.hpp"
}

func init() {
	schema.Register("cpp", &CppBackend{})
}
