package cpp

import (
	"fmt"
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	PointerTemps *template.Template
)

func init() {
	PointerTemps = template.New("PointerTemps")
	template.Must(PointerTemps.New("marshal").Parse(`
	{
		if ({{.Target}} == NULL) {
			buf[i] = 0;
			i++;
		} else {
			buf[i] = 1;
			i++;
			{{.SubTypeCode}}
		}
	}`))
	template.Must(PointerTemps.New("unmarshal").Parse(`
	{
		if (buf[i] == 1) {
			i++;
			if ({{.Target}} == NULL) {
				{{.Target}} = new class {{.SubField}}();
			}
			{{.SubTypeCode}}
		} else {
			i++;
			{{.Target}} = NULL;
		}
	}`))
	template.Must(PointerTemps.New("size").Parse(`
	{
		s++;
		if ({{.Target}} != NULL) {
			{{.SubTypeCode}}
		}
	}`))

	template.Must(PointerTemps.New("field").Parse(`*`))
}

type PointerTemp struct {
	*schema.PointerType
	W           *Walker
	Target      string
	SubTypeCode string
	SubField    string
}

func (w *Walker) WalkPointerDef(pt *schema.PointerType) (parts *StringBuilder, err error) {
	sub, err := w.WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf("%s*", sub.String()))
	return
}

func (w *Walker) WalkPointerSize(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := w.WalkTypeSize(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(PointerTemps, "size", PointerTemp{pt, w, target, subtypecode.String(), ""})
	return
}

func (w *Walker) WalkPointerMarshal(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := w.WalkTypeMarshal(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(PointerTemps, "marshal", PointerTemp{pt, w, target, subtypecode.String(), subfield.String()})
	return
}

func (w *Walker) WalkPointerUnmarshal(pt *schema.PointerType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	subtypecode, err := w.WalkTypeUnmarshal(pt.SubType, "(*"+target+")")
	if err != nil {
		return nil, err
	}
	subfield, err := w.WalkTypeDef(pt.SubType)
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(PointerTemps, "unmarshal", PointerTemp{pt, w, target, subtypecode.String(), subfield.String()})
	return
}
