package cpp

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	StructTemps *template.Template
)

func init() {
	StructTemps = template.New("StructTemps")

	template.Must(StructTemps.New("marshal").Parse(`
	{
		uint64_t nbuf = {{.Target}}.Marshal(&buf[i]);
		i += nbuf;
	}`))
	template.Must(StructTemps.New("unmarshal").Parse(`
	{
		uint64_t ni = {{.Target}}.Unmarshal(&buf[i]);
		i += ni;
	}`))
	template.Must(StructTemps.New("size").Parse(`
	{
		s += {{.Target}}.MarshalSize();
	}`))
	template.Must(StructTemps.New("field").Parse(`C{{.Struct}}`))
}

type StructTemp struct {
	*schema.StructType
	W      *Walker
	Target string
}

func (w *Walker) WalkStructDef(st *schema.StructType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "field", st)
	return
}

func (w *Walker) WalkStructSize(st *schema.StructType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "size", StructTemp{st, w, target})
	return
}

func (w *Walker) WalkStructMarshal(st *schema.StructType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "marshal", StructTemp{st, w, target})
	return
}

func (w *Walker) WalkStructUnmarshal(st *schema.StructType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "unmarshal", StructTemp{st, w, target})
	return
}
