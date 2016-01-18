package golang

import (
	"go/format"

	"github.com/andyleap/gencode/schema"
)

var (
	GolangBackend = &schema.Backend{
		Generate: func(s *schema.Schema) (string, error) {
			def, err := WalkSchema(s)
			if err != nil {
				return "", err
			}
			out, err := format.Source([]byte(def.String()))
			if err != nil {
				return def.String(), nil
			}
			return string(out), nil
		},
	}
)

func init() {
	schema.Register("go", GolangBackend)
}
