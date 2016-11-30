package cpp

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
		uint64_t l = {{.Target}}.length();
		{{.VarIntCode}}
		memcpy(&buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}], {{.Target}}.c_str(), l);
		i += l;
	}`))
	template.Must(StringTemps.New("unmarshal").Parse(`
	{
		uint64_t l = 0;
		{{.VarIntCode}}
		{{.Target}}.assign((const char*)&buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}], l);
		i += l;
	}`))
	template.Must(StringTemps.New("size").Parse(`
	{
		uint64_t l = {{.Target}}.length();
		{{.VarIntCode}}
		s += l;
	}`))
	template.Must(StringTemps.New("field").Parse(`std::string`))
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
