package golang

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	StringTemps *template.Template
)

func init() {
	StringTemps = template.New("StringTemps")

	template.Must(StringTemps.New("marshal").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		copy(buf[i:], {{.Target}})
		i += l
	}`))
	template.Must(StringTemps.New("unmarshal").Parse(`
	{
		l := uint64(0)
		{{.VarIntCode}}
		{{.Target}} = string(buf[i:i+l])
		i += l
	}`))
	template.Must(StringTemps.New("size").Parse(`
	{
		l := uint64(len({{.Target}}))
		{{.VarIntCode}}
		s += l
	}`))
	template.Must(StringTemps.New("field").Parse(`string`))
}

type StringTemp struct {
	*schema.StringType
	W          *Walker
	Target     string
	VarIntCode string
}

func (w *Walker) WalkStringDef(st *schema.StringType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StringTemps, "field", st)
	return
}

func (w *Walker) WalkStringSize(st *schema.StringType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntSize(intHandler, "l")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(StringTemps, "size", StringTemp{st, w, target, intcode.String()})
	return
}

func (w *Walker) WalkStringMarshal(st *schema.StringType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntMarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(StringTemps, "marshal", StringTemp{st, w, target, intcode.String()})
	return
}

func (w *Walker) WalkStringUnmarshal(st *schema.StringType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntUnmarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(StringTemps, "unmarshal", StringTemp{st, w, target, intcode.String()})
	return
}
