package golang

import (
	"text/template"

	"github.com/andyleap/gencode/schema"
)

var (
	BoolTemps *template.Template
)

func init() {
	BoolTemps = template.New("BoolTemps")

	template.Must(BoolTemps.New("marshal").Parse(`
	{
		if {{.Target}} {
			buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}] = 1
		} else {
			buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}] = 0
		}
	}`))
	template.Must(BoolTemps.New("unmarshal").Parse(`
	{
		{{.Target}} = buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}] == 1
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
	w.Offset++
	return
}

func (w *Walker) WalkBoolMarshal(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "marshal", BoolTemp{bt, w, target})
	w.Offset++
	return
}

func (w *Walker) WalkBoolUnmarshal(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "unmarshal", BoolTemp{bt, w, target})
	w.Offset++
	return
}
