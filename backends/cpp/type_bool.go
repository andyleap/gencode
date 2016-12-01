package cpp

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	BoolTemps *template.Template
)

func init() {
	BoolTemps = template.New("BoolTemps")

	template.Must(BoolTemps.New("marshal").Parse(`
	{
		if ({{.Target}}) {
			buf[i] = 1;
		} else {
			buf[i] = 0;
		}
		i++;
	}`))
	template.Must(BoolTemps.New("unmarshal").Parse(`
	{
		{{.Target}} = buf[i] == 1;
		i++;
	}`))
	template.Must(BoolTemps.New("size").Parse(`
	{
		s++;
	}`))
}

type BoolTemp struct {
	*schema.BoolType
	W      *Walker
	Target string
}

func (w *Walker) WalkBoolDef(bt *schema.BoolType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("bool")
	return
}

func (w *Walker) WalkBoolSize(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "size", BoolTemp{bt, w, target})
	return
}

func (w *Walker) WalkBoolMarshal(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "marshal", BoolTemp{bt, w, target})
	return
}

func (w *Walker) WalkBoolUnmarshal(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "unmarshal", BoolTemp{bt, w, target})
	return
}
