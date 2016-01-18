package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	ByteTemps *template.Template
)

func init() {
	ByteTemps = template.New("ByteTemps")

	template.Must(ByteTemps.New("marshal").Parse(`
	{
		buf[i] = {{.Target}}
		i++
	}`))
	template.Must(ByteTemps.New("unmarshal").Parse(`
	{
		{{.Target}} = buf[i]
		i++
	}`))
}

type ByteTemp struct {
	*schema.ByteType
	Target string
}

func WalkByteDef(bt *schema.ByteType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("byte")
	return
}

func WalkByteSize(bt *schema.ByteType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(`
	s += 1`)
	return
}

func WalkByteMarshal(bt *schema.ByteType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(ByteTemps, "marshal", ByteTemp{bt, target})
	return
}

func WalkByteUnmarshal(bt *schema.ByteType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(ByteTemps, "unmarshal", ByteTemp{bt, target})
	return
}
