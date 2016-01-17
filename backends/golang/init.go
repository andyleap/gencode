package golang

import "github.com/andyleap/gencode/schema"

var (
	GolangBackend = &schema.Backend{
		Generate: func(s *schema.Schema) (string, error) {
			def, err := WalkSchema(s)
			if err != nil {
				return "", err
			}
			return def.String(), nil
		},
	}
)

func init() {
	schema.Register("go", GolangBackend)
}
