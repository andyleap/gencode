package golang

import (
	"sort"
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	FloatTemps *template.Template
)

func init() {
	FloatTemps = template.New("FloatTemps").Funcs(template.FuncMap{
		"BitRange": func(bits int, bigEndian bool) []int {
			ar := []int{0, 8, 16, 24, 32, 40, 48, 56, 64}[0:(bits / 8)]
			if bigEndian {
				sort.Sort(sort.Reverse(sort.IntSlice(ar)))
			}
			return ar
		},
	})

	template.Must(FloatTemps.New("marshal").Parse(`
	{
		{{if .W.Unsafe}}
		*(*float{{.Bits}})(unsafe.Pointer(&buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}])) = {{.Target}}
		{{else}}
		v := *(*uint{{.Bits}})(unsafe.Pointer(&({{.Target}})))
		{{range $pos, $bits := BitRange .Bits .W.BigEndian}}
		buf[{{if $.W.IAdjusted}}i + {{end}}{{$pos}} + {{$.W.Offset}}] = byte(v >> {{$bits}})
		{{end}}
		{{end}}
	}`))
	template.Must(FloatTemps.New("unmarshal").Parse(`
	{
		{{if .W.Unsafe}}
		{{.Target}} = *(*float{{.Bits}})(unsafe.Pointer(&buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}]))
		{{else}}
		v := 0{{range $pos, $bits := BitRange .Bits .W.BigEndian}} | (uint{{$.Bits}}(buf[{{if $.W.IAdjusted}}i + {{end}}{{$pos}} + {{$.W.Offset}}]) << {{$bits}}){{end}}
		{{.Target}} = *(*float{{.Bits}})(unsafe.Pointer(&v))
		{{end}}
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
	w.Offset += ft.Bits / 8
	return
}

func (w *Walker) WalkFloatMarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "marshal", FloatTemp{ft, w, target})
	w.Offset += ft.Bits / 8
	return
}

func (w *Walker) WalkFloatUnmarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "unmarshal", FloatTemp{ft, w, target})
	w.Offset += ft.Bits / 8
	return
}
