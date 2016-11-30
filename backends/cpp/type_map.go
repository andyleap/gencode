package cpp

import (
	"fmt"
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	MapTemps *template.Template
)

func init() {
	MapTemps = template.New("MapTemps")
	template.Must(MapTemps.New("marshal").Parse(`
	{
		uint64_t l = {{.Target}}.size();
		{{.VarIntCode}}
		for (std::map<{{.KeySubField}}, {{.ValueSubField}}>::iterator it = {{.Target}}.begin(); it != {{.Target}}.end(); it++) {
			const {{.KeySubField}}& k = it->first;
			{{.KeySubTypeCode}}
			const {{.ValueSubField}}& v = it->second;
			{{.ValueSubTypeCode}}
		}
	}`))
	template.Must(MapTemps.New("unmarshal").Parse(`
	{
		uint64_t l = 0;
		{{.VarIntCode}}
		for (uint64_t x = 0; x < l; x++) {
			{{.KeySubField}} k;
			{{.KeySubTypeCode}}
			{{.ValueSubField}} v;
			{{.ValueSubTypeCode}}
			{{.Target}}[k] = v;
		}
	}`))
	template.Must(MapTemps.New("size").Parse(`
	{
		uint64_t l = {{.Target}}.size();
		{{.VarIntCode}}
		for (std::map<{{.KeySubField}}, {{.ValueSubField}}>::iterator it = {{.Target}}.begin(); it != {{.Target}}.end(); it++) {
			const {{.KeySubField}}& k = it->first;
			{{.KeySubTypeCode}}
			const {{.ValueSubField}}& v = it->second;
			{{.ValueSubTypeCode}}
		}
	}`))
}

type MapTemp struct {
	*schema.MapType
	W                *Walker
	Target           string
	KeySubTypeCode   string
	KeySubField      string
	ValueSubTypeCode string
	ValueSubField    string
	VarIntCode       string
}

func (w *Walker) WalkMapDef(st *schema.MapType) (parts *StringBuilder, err error) {
	ksub, err := w.WalkTypeDef(st.KeySubType)
	if err != nil {
		return nil, err
	}
	vsub, err := w.WalkTypeDef(st.ValueSubType)
	if err != nil {
		return nil, err
	}
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf("std::map<%s, %s>", ksub.String(), vsub.String()))
	return
}

func (w *Walker) WalkMapSize(st *schema.MapType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntSize(intHandler, "l")
	if err != nil {
		return nil, err
	}
	keysubtype, err := w.WalkTypeDef(st.KeySubType)
	if err != nil {
		return nil, err
	}
	keysubtypecode, err := w.WalkTypeSize(st.KeySubType, "k")
	if err != nil {
		return nil, err
	}
	valuesubtype, err := w.WalkTypeDef(st.ValueSubType)
	if err != nil {
		return nil, err
	}
	valuesubtypecode, err := w.WalkTypeSize(st.ValueSubType, "v")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(MapTemps, "size", MapTemp{st, w, target, keysubtypecode.String(), keysubtype.String(), valuesubtypecode.String(), valuesubtype.String(), intcode.String()})
	return
}

func (w *Walker) WalkMapMarshal(st *schema.MapType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntMarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	keysubtype, err := w.WalkTypeDef(st.KeySubType)
	if err != nil {
		return nil, err
	}
	keysubtypecode, err := w.WalkTypeMarshal(st.KeySubType, "k")
	if err != nil {
		return nil, err
	}
	valuesubtype, err := w.WalkTypeDef(st.ValueSubType)
	if err != nil {
		return nil, err
	}
	valuesubtypecode, err := w.WalkTypeMarshal(st.ValueSubType, "v")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(MapTemps, "marshal", MapTemp{st, w, target, keysubtypecode.String(), keysubtype.String(), valuesubtypecode.String(), valuesubtype.String(), intcode.String()})
	return
}

func (w *Walker) WalkMapUnmarshal(st *schema.MapType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntUnmarshal(intHandler, "l")
	if err != nil {
		return nil, err
	}
	keysubtype, err := w.WalkTypeDef(st.KeySubType)
	if err != nil {
		return nil, err
	}
	keysubtypecode, err := w.WalkTypeUnmarshal(st.KeySubType, "k")
	if err != nil {
		return nil, err
	}
	valuesubtype, err := w.WalkTypeDef(st.ValueSubType)
	if err != nil {
		return nil, err
	}
	valuesubtypecode, err := w.WalkTypeUnmarshal(st.ValueSubType, "v")
	if err != nil {
		return nil, err
	}
	err = parts.AddTemplate(MapTemps, "unmarshal", MapTemp{st, w, target, keysubtypecode.String(), keysubtype.String(), valuesubtypecode.String(), valuesubtype.String(), intcode.String()})
	return
}
