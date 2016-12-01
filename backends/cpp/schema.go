package cpp

import (
	"fmt"

	"github.com/eyrie-io/gencode/schema"
)

type Walker struct {
	Needs     []string
	Offset    int
	IAdjusted bool
	Unsafe    bool
}

func (w *Walker) WalkSchema(s *schema.Schema, Namespace string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`#pragma once

#include <vector>
#include <string>
#include <map>
#include <stdint.h>

namespace %s {

`, Namespace))
	for _, st := range s.Structs {
		p, err := w.WalkStruct(st)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(`
}
`)
	return
}
