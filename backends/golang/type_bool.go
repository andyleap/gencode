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
			buf[i] = 1
		} else {
			buf[i] = 0
		}
		i++
	}`))
	template.Must(BoolTemps.New("unmarshal").Parse(`
	{
		{{.Target}} = buf[i] == 1
		i++
	}`))
}

type BoolTemp struct {
	*schema.BoolType
	Target string
}

func WalkBoolDef(bt *schema.BoolType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append("bool")
	return
}

func WalkBoolSize(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(`
	s += 1`)
	return
}

func WalkBoolMarshal(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "marshal", BoolTemp{bt, target})
	return
}

func WalkBoolUnmarshal(bt *schema.BoolType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(BoolTemps, "unmarshal", BoolTemp{bt, target})
	return
}
