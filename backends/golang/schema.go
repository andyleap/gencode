package golang

import (
	"fmt"

	"github.com/andyleap/gencode/schema"
)

func WalkSchema(s *schema.Schema) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`package main
	
	import (
		"math"
	)
	
	var (
		_ = math.Float64frombits
	)
	`))
	for _, st := range s.Structs {
		p, err := WalkStruct(st)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	return
}
