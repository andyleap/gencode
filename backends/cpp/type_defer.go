package cpp

import "github.com/eyrie-io/gencode/schema"

func (w *Walker) WalkDeferDef(dt *schema.DeferType) (parts *StringBuilder, err error) {
	return w.WalkTypeDef(dt.Resolved)
}

func (w *Walker) WalkDeferSize(dt *schema.DeferType, target string) (parts *StringBuilder, err error) {
	return w.WalkTypeSize(dt.Resolved, target)
}

func (w *Walker) WalkDeferMarshal(dt *schema.DeferType, target string) (parts *StringBuilder, err error) {
	return w.WalkTypeMarshal(dt.Resolved, target)
}

func (w *Walker) WalkDeferUnmarshal(dt *schema.DeferType, target string) (parts *StringBuilder, err error) {
	return w.WalkTypeUnmarshal(dt.Resolved, target)
}
