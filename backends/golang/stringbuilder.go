package golang

import (
	"bytes"
	"strings"
	"text/template"
)

type StringBuilder []string

func (sb *StringBuilder) Append(s string) {
	*sb = append(*sb, s)
}

func (sb *StringBuilder) Join(s *StringBuilder) {
	*sb = append(*sb, *s...)
}

func (sb *StringBuilder) String() string {
	return strings.Join(*sb, "")
}

func (sb *StringBuilder) AddTemplate(t *template.Template, name string, data interface{}) error {
	buf := &bytes.Buffer{}
	err := t.ExecuteTemplate(buf, name, data)
	if err != nil {
		return err
	}
	sb.Append(buf.String())
	return nil
}
