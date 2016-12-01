package golang

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	FloatTemps *template.Template
)

func init() {
	FloatTemps = template.New("FloatTemps").Funcs(template.FuncMap{
		"Bytes": func(bits int) int {
			return bits / 8
		},
		"BitRange": func(bits int) []int {
			return []int{0, 8, 16, 24, 32, 40, 48, 56, 64}[0:(bits / 8)]
		},
	})

	template.Must(FloatTemps.New("marshal").Parse(`
	{
		{{if .W.Unsafe}}
		*(*float{{.Bits}})(unsafe.Pointer(&buf[i])) = {{.Target}}
		{{else}}
		v := *(*uint{{.Bits}})(unsafe.Pointer(&({{.Target}})))
		{{range BitRange .Bits}}
		buf[i + {{Bytes .}}] = byte(v >> {{.}})
		{{end}}
		{{end}}
		i += {{.Bits}}/8
	}`))
	template.Must(FloatTemps.New("unmarshal").Parse(`
	{
		{{if .W.Unsafe}}
		{{.Target}} = *(*float{{.Bits}})(unsafe.Pointer(&buf[i]))
		{{else}}
		v := 0{{range BitRange .Bits}} | (uint{{$.Bits}}(buf[i + {{Bytes .}}]) << {{.}}){{end}}
		{{.Target}} = *(*float{{.Bits}})(unsafe.Pointer(&v))
		{{end}}
		i += {{.Bits}}/8
	}`))
	template.Must(FloatTemps.New("size").Parse(`
	{
		s += {{.Bits}}/8
	}`))
	template.Must(FloatTemps.New("field").Parse(`float{{.Bits}}`))
}

type FloatTemp struct {
	*schema.FloatType
	W      *Walker
	Target string
}

func (w *Walker) WalkFloatDef(ft *schema.FloatType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "field", ft)
	return
}

func (w *Walker) WalkFloatSize(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "size", FloatTemp{ft, w, target})
	return
}

func (w *Walker) WalkFloatMarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "marshal", FloatTemp{ft, w, target})
	return
}

func (w *Walker) WalkFloatUnmarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "unmarshal", FloatTemp{ft, w, target})
	return
}
