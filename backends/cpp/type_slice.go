package cpp

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/eyrie-io/gencode/schema"
)

var (
	SliceTemps *template.Template
)

func init() {
	SliceTemps = template.New("SliceTemps")
	template.Must(SliceTemps.New("marshal").Parse(`
	{
		uint64_t l{{.Depth}} = {{.Target}}.size();
		{{.VarIntCode}}
		for (uint64_t k{{.Depth}} = 0; k{{.Depth}} < l{{.Depth}}; k{{.Depth}}++) {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}i += {{.SubOffset}};{{end}}
		}
	}`))
	template.Must(SliceTemps.New("unmarshal").Parse(`
	{
		uint64_t l{{.Depth}} = 0;
		{{.VarIntCode}}
		{{.Target}}.resize(l{{.Depth}});
		for (uint64_t k{{.Depth}} = 0; k{{.Depth}} < l{{.Depth}}; k{{.Depth}}++) {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}i += {{.SubOffset}};{{end}}
		}
	}`))
	template.Must(SliceTemps.New("bytemarshal").Parse(`
	{
		uint64_t l{{.Depth}} = {{.Target}}.size();
		{{.VarIntCode}}
		memcpy(&buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}], &{{.Target}}[0], l{{.Depth}});
		i += l{{.Depth}};
	}`))
	template.Must(SliceTemps.New("byteunmarshal").Parse(`
	{
		uint64_t l{{.Depth}} = 0;
		{{.VarIntCode}}
		{{.Target}}.resize(l{{.Depth}});
		memcpy(&{{.Target}}[0], &buf[{{if .W.IAdjusted}}i + {{end}}{{.W.Offset}}], l);
		i += l{{.Depth}};
	}`))
	template.Must(SliceTemps.New("size").Parse(`
	{
		uint64_t l{{.Depth}} = {{.Target}}.size();
		{{.VarIntCode}}
		{{if eq .SubTypeCode "" }}
		{{if gt .SubOffset 0 }}
		s += {{.SubOffset}}*l{{.Depth}};
		{{end}}
		{{else}}
		for (uint64_t k{{.Depth}} = 0; k{{.Depth}} < l{{.Depth}}; k{{.Depth}}++) {
			{{.SubTypeCode}}
			{{if gt .SubOffset 0 }}s += {{.SubOffset}};{{end}}
		}
		{{end}}
	}`))
	template.Must(SliceTemps.New("bytesize").Parse(`
	{
		uint64_t l{{.Depth}} = {{.Target}}.size();
		{{.VarIntCode}}
		s += l{{.Depth}};
	}`))
	template.Must(SliceTemps.New("field").Parse(`[]`))
}

type SliceTemp struct {
	*schema.SliceType
	W           *Walker
	SubOffset   int
	Target      string
	SubTypeCode string
	SubField    string
	VarIntCode  string
	Depth       int
}

func (w *Walker) WalkSliceDef(st *schema.SliceType) (parts *StringBuilder, err error) {
	sub, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	parts = &StringBuilder{}
	x := sub.String()
	if strings.HasSuffix(x, ">") {
		x = x + " "
	}
	parts.Append(fmt.Sprintf("std::vector<%s>", x))
	return
}

func (w *Walker) WalkSliceSize(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntSize(intHandler, "l"+strconv.Itoa(st.Depth))
	if err != nil {
		return nil, err
	}
	offset := w.Offset
	subtypecode, err := w.WalkTypeSize(st.SubType, target+"[k"+strconv.Itoa(st.Depth)+"]")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytesize", SliceTemp{st, w, SubOffset, target, subtypecode.String(), "", intcode.String(), st.Depth})
	} else {
		err = parts.AddTemplate(SliceTemps, "size", SliceTemp{st, w, SubOffset, target, subtypecode.String(), "", intcode.String(), st.Depth})
	}
	return
}

func (w *Walker) WalkSliceMarshal(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntMarshal(intHandler, "l"+strconv.Itoa(st.Depth))
	if err != nil {
		return nil, err
	}
	offset := w.Offset
	subtypecode, err := w.WalkTypeMarshal(st.SubType, target+"[k"+strconv.Itoa(st.Depth)+"]")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	subfield, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "bytemarshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), st.Depth})
	} else {
		err = parts.AddTemplate(SliceTemps, "marshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), st.Depth})
	}
	return
}

func (w *Walker) WalkSliceUnmarshal(st *schema.SliceType, target string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	intHandler := &schema.IntType{
		Bits:   64,
		Signed: false,
		VarInt: true,
	}
	intcode, err := w.WalkIntUnmarshal(intHandler, "l"+strconv.Itoa(st.Depth))
	if err != nil {
		return nil, err
	}
	offset := w.Offset
	subtypecode, err := w.WalkTypeUnmarshal(st.SubType, target+"[k"+strconv.Itoa(st.Depth)+"]")
	if err != nil {
		return nil, err
	}
	SubOffset := w.Offset - offset
	w.Offset = offset
	subfield, err := w.WalkTypeDef(st.SubType)
	if err != nil {
		return nil, err
	}
	if _, ok := st.SubType.(*schema.ByteType); ok {
		err = parts.AddTemplate(SliceTemps, "byteunmarshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), st.Depth})
	} else {
		err = parts.AddTemplate(SliceTemps, "unmarshal", SliceTemp{st, w, SubOffset, target, subtypecode.String(), subfield.String(), intcode.String(), st.Depth})
	}
	return
}
