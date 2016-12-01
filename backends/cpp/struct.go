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
	parts.Append(fmt.Sprintf(`class C%s {
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
	~C%s();
	uint64_t MarshalSize();
	uint64_t Marshal(uint8_t* buf);
	uint64_t Unmarshal(uint8_t* buf);
};

C%s::~C%s() {`, s.Name, s.Name, s.Name))
	for _, f := range s.Fields {
		switch t := f.Type.(type) {
		case *schema.PointerType:
			parts.Append(fmt.Sprintf(`	delete this->%s;"`, s.Name))
		case *schema.SliceType:
			pt, ok := t.SubType.(*schema.PointerType)
			if ok {
				dt := pt.SubType.(*schema.DeferType)
				parts.Append(fmt.Sprintf(`
	for (std::vector<C%s*>::iterator i = %s.begin(); i != %s.end(); i++) {
		delete (*i);
	}`, dt.Defer, f.Name, f.Name))
			}
		}
	}
	parts.Append(fmt.Sprintf(`
}

uint64_t C%s::MarshalSize() {
	uint64_t s = 0;
`, s.Name))
	for _, f := range s.Fields {
		p, err := w.WalkFieldSize(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(fmt.Sprintf(`
	return s;
}

uint64_t C%s::Marshal(uint8_t* buf) {`, s.Name))
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
	return i;
}

uint64_t C%s::Unmarshal(uint8_t* buf) {
	uint64_t i = 0;
	`, s.Name))
	for _, f := range s.Fields {
		p, err := w.WalkFieldUnmarshal(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(fmt.Sprintf(`
	return i;
}

`))
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
