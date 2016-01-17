package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	StructTemps *template.Template
)

func init() {
	StructTemps = template.New("StructTemps")

	template.Must(StructTemps.New("marshal").Parse(`
	{
		nbuf, err := {{.Target}}.Marshal(buf[i:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}`))
	template.Must(StructTemps.New("unmarshal").Parse(`
	{
		ni, err := {{.Target}}.Unmarshal(buf[i:])
		if err != nil {
			return 0, err
		}
		i += ni
	}`))
	template.Must(StructTemps.New("size").Parse(`
	{
		s += {{.Target}}.Size()
	}`))
	template.Must(StructTemps.New("field").Parse(`{{.Struct}}`))
}

type StructTemp struct {
	*schema.StructType
	Target string
}

func WalkStructDef(st *schema.StructType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "field", st)
	return
}

func WalkStructSize(st *schema.StructType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "size", StructTemp{st, target})
	return
}

func WalkStructMarshal(st *schema.StructType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "marshal", StructTemp{st, target})
	return
}

func WalkStructUnmarshal(st *schema.StructType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(StructTemps, "unmarshal", StructTemp{st, target})
	return
}
