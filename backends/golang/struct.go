package golang

import (
	"fmt"

	"github.com/andyleap/gencode/schema"
)

func WalkStruct(s *schema.Struct) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`type %s struct {
	`, s.Name))
	for _, f := range s.Fields {
		p, err := WalkFieldDef(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
		parts.Append(`
	`)
	}
	parts.Append(fmt.Sprintf(`}
	
func (d *%s) Size() (s uint64) {
	`, s.Name))
	for _, f := range s.Fields {
		p, err := WalkFieldSize(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(fmt.Sprintf(`
	return
	}
	
func (d *%s) Marshal(buf []byte) ([]byte, error) {
	{
		size := d.Size()
		if uint64(cap(buf)) >= d.Size() {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)
	`, s.Name))
	for _, f := range s.Fields {
		p, err := WalkFieldMarshal(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(fmt.Sprintf(`
	return buf[:i], nil
}
	
func (d *%s) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	`, s.Name))
	for _, f := range s.Fields {
		p, err := WalkFieldUnmarshal(f)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	parts.Append(`
	return i, nil
	}
`)
	return
}

func WalkFieldDef(s *schema.Field) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`%s `, s.Name))
	subp, err := WalkTypeDef(s.Type)
	if err != nil {
		return nil, err
	}
	parts.Join(subp)
	return
}

func WalkFieldSize(s *schema.Field) (parts *StringBuilder, err error) {
	return WalkTypeSize(s.Type, "d."+s.Name)
}

func WalkFieldMarshal(s *schema.Field) (parts *StringBuilder, err error) {
	return WalkTypeMarshal(s.Type, "d."+s.Name)
}

func WalkFieldUnmarshal(s *schema.Field) (parts *StringBuilder, err error) {
	return WalkTypeUnmarshal(s.Type, "d."+s.Name)
}
