package golang

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	TimeTemps *template.Template
)

func init() {
	TimeTemps = template.New("TimeTemps")

	template.Must(TimeTemps.New("marshal").Parse(`
	{
		b, err := {{.Target}}.MarshalBinary()
		if err != nil {
			return nil, err
		}
		copy(buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}:], b)
	}`))
	template.Must(TimeTemps.New("unmarshal").Parse(`
	{
		{{.Target}}.UnmarshalBinary(buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}:{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}} + 15])
	}`))
	template.Must(TimeTemps.New("field").Parse(`time.Time`))
}

type TimeTemp struct {
	*schema.TimeType
	W      *Walker
	Target string
}

func (w *Walker) WalkTimeDef(tt *schema.TimeType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(TimeTemps, "field", tt)
	return
}

func (w *Walker) WalkTimeSize(tt *schema.TimeType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	w.Offset += 15
	return
}

func (w *Walker) WalkTimeMarshal(tt *schema.TimeType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(TimeTemps, "marshal", TimeTemp{tt, w, target})
	w.Offset += 15
	return
}

func (w *Walker) WalkTimeUnmarshal(tt *schema.TimeType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(TimeTemps, "unmarshal", TimeTemp{tt, w, target})
	w.Offset += 15
	return
}
