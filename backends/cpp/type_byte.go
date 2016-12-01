package cpp

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	ByteTemps *template.Template
)

func init() {
	ByteTemps = template.New("ByteTemps")

	template.Must(ByteTemps.New("marshal").Parse(`
	{
		buf[i] = {{.Target}};
		i++;
	}`))
	template.Must(ByteTemps.New("unmarshal").Parse(`
	{
		{{.Target}} = buf[i];
		i++;
	}`))
	template.Must(ByteTemps.New("size").Parse(`
	{
		s++;
	}`))
}

type ByteTemp struct {
	*schema.ByteType
	W      *Walker
	Target string
}

func (w *Walker) WalkByteDef(bt *schema.ByteType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("uint8_t")
	return
}

func (w *Walker) WalkByteSize(bt *schema.ByteType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(ByteTemps, "size", ByteTemp{bt, w, target})
	return
}

func (w *Walker) WalkByteMarshal(bt *schema.ByteType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(ByteTemps, "marshal", ByteTemp{bt, w, target})
	return
}

func (w *Walker) WalkByteUnmarshal(bt *schema.ByteType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(ByteTemps, "unmarshal", ByteTemp{bt, w, target})
	return
}
