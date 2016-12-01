package cpp

import (
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	IntTemps *template.Template
)

func init() {
	IntTemps = template.New("IntTemps").Funcs(template.FuncMap{
		"Bytes": func(bits int) int {
			return bits / 8
		},
		"BitRange": func(bits int) []int {
			return []int{0, 8, 16, 24, 32, 40, 48, 56, 64}[0:(bits / 8)]
		},
	})

	template.Must(IntTemps.New("marshal").Parse(`
	{
		{{if .VarInt}}uint{{.Bits}}_t t = {{.Target}};
		{{if .Signed}}t <<= 1;
		if ({{.Target}} < 0) { t = ^t; }{{end}}
   		while (t >= 0x80) {
			buf[i] = uint8_t(t) | 0x80;
			t >>= 7;
			i++;
		}
		buf[i] = uint8_t(t);
		i++;
		{{else}}
		memcpy(&buf[i], &{{.Target}}, {{.Bits}}/8);
		i += {{.Bits}}/8;
		{{end}}
	}`))
	template.Must(IntTemps.New("unmarshal").Parse(`
	{
		{{if .VarInt}}
		uint8_t bs = 7;
		uint{{.Bits}}_t t = buf[i] & 0x7F;
		while ((buf[i] & 0x80) == 0x80) {
			i++;
			t |= ((uint{{.Bits}}_t)buf[i] & 0x7F) << bs;
			bs += 7;
		}
		i++;
		{{if .Signed}}
		{{.Target}} = (int{{.Bits}}_t)(t >> 1);
		if (t&1 != 0) {
			{{.Target}} = ^{{.Target}};
		}
		{{else}}
		{{.Target}} = t;
		{{end}}
		{{else}}
		memcpy(&{{.Target}}, &buf[i], {{.Bits}}/8);
		i += {{.Bits}}/8;
		{{end}}
	}`))
	template.Must(IntTemps.New("field").Parse(`{{if not .Signed}}u{{end}}int{{.Bits}}_t`))
	template.Must(IntTemps.New("size").Parse(`
	{
		{{if .VarInt}}
		{{if .Signed}}
		uint{{.Bits}}_t t := {{.Target}};
		t <<= 1;
		if ({{.Target}}) < 0 {
			t = ^t;
		}
		while (t >= 0x80) {
			t >>= 7;
			s++;
		}
		s++;
		{{else}}
		uint{{.Bits}}_t t = {{.Target}};
		while (t >= 0x80) {
			t >>= 7;
			s++;
		}
		s++;
		{{end}}
		{{else}}
		s += {{.Bits}}/8;
		{{end}}
	}`))
}

type IntTemp struct {
	*schema.IntType
	W      *Walker
	Target string
}

func (w *Walker) WalkIntDef(it *schema.IntType) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(IntTemps, "field", it)
	return
}

func (w *Walker) WalkIntSize(it *schema.IntType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(IntTemps, "size", IntTemp{it, w, target})
	return
}

func (w *Walker) WalkIntMarshal(it *schema.IntType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(IntTemps, "marshal", IntTemp{it, w, target})
	return
}

func (w *Walker) WalkIntUnmarshal(it *schema.IntType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	err = parts.AddTemplate(IntTemps, "unmarshal", IntTemp{it, w, target})
	return
}
