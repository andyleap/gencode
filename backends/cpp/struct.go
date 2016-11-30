package cpp

import (
	"fmt"

	"github.com/eyrie-io/gencode/schema"
)

func (w *Walker) WalkStruct(s *schema.Struct) (parts *StringBuilder, err error) {
	if s.Framed {
		err = fmt.Errorf("Framed structs are not supported in cpp")
		return
	}
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`class %s {
public:
`, s.Name))
	for _, f := range s.Fields {
		p, err := w.WalkFieldDef(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
		parts.Append(`
`)
	}
	parts.Append(fmt.Sprintf(`
public:
	uint64_t MarshalSize();
	uint64_t Marshal(uint8_t* buf);
	uint64_t Unmarshal(uint8_t* buf);
};

uint64_t %s::MarshalSize() {
	uint64_t s = 0;
`, s.Name))
	for _, f := range s.Fields {
		p, err := w.WalkFieldSize(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	if w.Offset > 0 {
		parts.Append(fmt.Sprintf(`
	s += %d;`, w.Offset))
		w.Offset = 0
	}
	w.IAdjusted = false

	parts.Append(fmt.Sprintf(`
	return s;
}

uint64_t %s::Marshal(uint8_t* buf) {`, s.Name))
	parts.Append(`
	uint64_t size = this->MarshalSize();`)
	parts.Append(`
	uint64_t i = 0;
	`)
	for _, f := range s.Fields {
		p, err := w.WalkFieldMarshal(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(fmt.Sprintf(`
	return i+%d;
}

uint64_t %s::Unmarshal(uint8_t* buf) {
	uint64_t i = 0;
	`, w.Offset, s.Name))
	w.Offset = 0
	for _, f := range s.Fields {
		p, err := w.WalkFieldUnmarshal(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(fmt.Sprintf(`
	return i + %d;
}

`, w.Offset))
	w.Offset = 0
	w.IAdjusted = false
	return
}

func (w *Walker) WalkFieldDef(s *schema.Field) (parts *StringBuilder, err error) {
	switch t := s.Type.(type) {
	case *schema.ArrayType:
		subp, err := w.WalkTypeDef(t.SubType)
		if err != nil {
			return nil, err
		}
		parts = &StringBuilder{}
		parts.Append(fmt.Sprintf(`	%s %s[%d];`, subp.String(), s.Name, t.Count))
	default:
		subp, err := w.WalkTypeDef(s.Type)
		if err != nil {
			return nil, err
		}
		parts = &StringBuilder{}
		parts.Append(fmt.Sprintf(`	%s %s;`, subp.String(), s.Name))
	}
	return
}

func (w *Walker) WalkFieldSize(s *schema.Field) (parts *StringBuilder, err error) {
	return w.WalkTypeSize(s.Type, "this->"+s.Name)
}

func (w *Walker) WalkFieldMarshal(s *schema.Field) (parts *StringBuilder, err error) {
	return w.WalkTypeMarshal(s.Type, "this->"+s.Name)
}

func (w *Walker) WalkFieldUnmarshal(s *schema.Field) (parts *StringBuilder, err error) {
	return w.WalkTypeUnmarshal(s.Type, "this->"+s.Name)
}
