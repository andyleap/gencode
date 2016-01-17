package golang

import "github.com/andyleap/gencode/schema"

func WalkDeferDef(dt *schema.DeferType) (parts *StringBuilder, err error) {
	return WalkTypeDef(dt.Resolved)
}

func WalkDeferSize(dt *schema.DeferType, target string) (parts *StringBuilder, err error) {
	return WalkTypeSize(dt.Resolved, target)
}

func WalkDeferMarshal(dt *schema.DeferType, target string) (parts *StringBuilder, err error) {
	return WalkTypeMarshal(dt.Resolved, target)
}

func WalkDeferUnmarshal(dt *schema.DeferType, target string) (parts *StringBuilder, err error) {
	return WalkTypeUnmarshal(dt.Resolved, target)
}
