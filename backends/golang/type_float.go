package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	FloatTemps *template.Template
)

func init() {
	FloatTemps = template.New("FloatTemps")

	template.Must(FloatTemps.New("marshal").Parse(`
	{
		v := math.Float{{.Bits}}bits({{.Target}})
		{{.IntCode}}
	}`))
	template.Must(FloatTemps.New("unmarshal").Parse(`
	{
		var v uint{{.Bits}}
		{{.IntCode}}
		{{.Target}} = math.Float{{.Bits}}frombits(v)
	}`))
	template.Must(FloatTemps.New("field").Parse(`float{{.Bits}}`))
}

type FloatTemp struct {
	*schema.FloatType
	Target  string
	IntCode string
}

func WalkFloatDef(ft *schema.FloatType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(FloatTemps, "field", ft)
	return
}

func WalkFloatSize(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	return WalkIntSize(&schema.IntType{Bits: ft.Bits}, "")
}

func WalkFloatMarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intcode, err := WalkIntMarshal(&schema.IntType{Bits: ft.Bits}, "v")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(FloatTemps, "marshal", FloatTemp{ft, target, intcode.String()})
	return
}

func WalkFloatUnmarshal(ft *schema.FloatType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intcode, err := WalkIntUnmarshal(&schema.IntType{Bits: ft.Bits}, "v")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(FloatTemps, "unmarshal", FloatTemp{ft, target, intcode.String()})
	return
}
