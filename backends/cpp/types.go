package cpp

import (
	"fmt"

	"github.com/eyrie-io/gencode/schema"
)

func (w *Walker) WalkTypeDef(t schema.Type) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.BoolType:
		return w.WalkBoolDef(tt)
	case *schema.ByteType:
		return w.WalkByteDef(tt)
	case *schema.DeferType:
		return w.WalkDeferDef(tt)
	case *schema.FloatType:
		return w.WalkFloatDef(tt)
	case *schema.IntType:
		return w.WalkIntDef(tt)
	case *schema.PointerType:
		return w.WalkPointerDef(tt)
	case *schema.SliceType:
		return w.WalkSliceDef(tt)
	case *schema.MapType:
		return w.WalkMapDef(tt)
	case *schema.StringType:
		return w.WalkStringDef(tt)
	case *schema.StructType:
		return w.WalkStructDef(tt)
	}
	return nil, fmt.Errorf("No such type %T", t)
}

func (w *Walker) WalkTypeSize(t schema.Type, target string) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ArrayType:
		return w.WalkArraySize(tt, target)
	case *schema.BoolType:
		return w.WalkBoolSize(tt, target)
	case *schema.ByteType:
		return w.WalkByteSize(tt, target)
	case *schema.DeferType:
		return w.WalkDeferSize(tt, target)
	case *schema.FloatType:
		return w.WalkFloatSize(tt, target)
	case *schema.IntType:
		return w.WalkIntSize(tt, target)
	case *schema.PointerType:
		return w.WalkPointerSize(tt, target)
	case *schema.SliceType:
		return w.WalkSliceSize(tt, target)
	case *schema.MapType:
		return w.WalkMapSize(tt, target)
	case *schema.StringType:
		return w.WalkStringSize(tt, target)
	case *schema.StructType:
		return w.WalkStructSize(tt, target)
	}
	return nil, fmt.Errorf("No such type %T", t)
}

func (w *Walker) WalkTypeMarshal(t schema.Type, target string) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ArrayType:
		return w.WalkArrayMarshal(tt, target)
	case *schema.BoolType:
		return w.WalkBoolMarshal(tt, target)
	case *schema.ByteType:
		return w.WalkByteMarshal(tt, target)
	case *schema.DeferType:
		return w.WalkDeferMarshal(tt, target)
	case *schema.FloatType:
		return w.WalkFloatMarshal(tt, target)
	case *schema.IntType:
		return w.WalkIntMarshal(tt, target)
	case *schema.PointerType:
		return w.WalkPointerMarshal(tt, target)
	case *schema.SliceType:
		return w.WalkSliceMarshal(tt, target)
	case *schema.MapType:
		return w.WalkMapMarshal(tt, target)
	case *schema.StringType:
		return w.WalkStringMarshal(tt, target)
	case *schema.StructType:
		return w.WalkStructMarshal(tt, target)
	}
	return nil, fmt.Errorf("No such type %T", t)
}

func (w *Walker) WalkTypeUnmarshal(t schema.Type, target string) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ArrayType:
		return w.WalkArrayUnmarshal(tt, target)
	case *schema.BoolType:
		return w.WalkBoolUnmarshal(tt, target)
	case *schema.ByteType:
		return w.WalkByteUnmarshal(tt, target)
	case *schema.DeferType:
		return w.WalkDeferUnmarshal(tt, target)
	case *schema.FloatType:
		return w.WalkFloatUnmarshal(tt, target)
	case *schema.IntType:
		return w.WalkIntUnmarshal(tt, target)
	case *schema.PointerType:
		return w.WalkPointerUnmarshal(tt, target)
	case *schema.SliceType:
		return w.WalkSliceUnmarshal(tt, target)
	case *schema.MapType:
		return w.WalkMapUnmarshal(tt, target)
	case *schema.StringType:
		return w.WalkStringUnmarshal(tt, target)
	case *schema.StructType:
		return w.WalkStructUnmarshal(tt, target)
	}
	return nil, fmt.Errorf("No such type %T", t)
}
