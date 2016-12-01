package golang

import (
	"fmt"

	"github.com/eyrie-io/gencode/schema"
)

type Walker struct {
	Unsafe bool
}

func (w *Walker) WalkSchema(s *schema.Schema, Package string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`package %s

	import (
		"unsafe"
		"io"
		"time"
	)

	var (
		_ = unsafe.Sizeof(0)
		_ = io.ReadFull
		_ = time.Now()
	)
	`, Package))
	for _, st := range s.Structs {
		p, err := w.WalkStruct(st)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	return
}
