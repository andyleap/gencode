package cpp

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
	memcpy(&buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}], &{{.Target}}, {{.Bits}}/8);`))
	template.Must(FloatTemps.New("unmarshal").Parse(`
	memcpy(&{{.Target}}, &buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}], {{.Bits}}/8);`))
	template.Must(FloatTemps.New("field").Parse(`{{if .IsFloat}}float{{else}}double{{end}}`))
}

type FloatTemp struct {
	*schema.FloatType
	W       *Walker
	Target  string
	IsFloat bool
}

func (w *Walker) WalkFloatDef(ft *schema.FloatType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "field", FloatTemp{ft, w, "", ft.Bits == 32})
	return
}

func (w *Walker) WalkFloatSize(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	w.Offset += ft.Bits / 8
	return
}

func (w *Walker) WalkFloatMarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "marshal", FloatTemp{ft, w, target, ft.Bits == 32})
	w.Offset += ft.Bits / 8
	return
}

func (w *Walker) WalkFloatUnmarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "unmarshal", FloatTemp{ft, w, target, ft.Bits == 32})
	w.Offset += ft.Bits / 8
	return
}
