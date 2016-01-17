package golang

import (
	"fmt"

	"github.com/andyleap/gencode/schema"
)

func WalkTypeDef(t schema.Type) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ByteType:
		return WalkByteDef(tt)
	case *schema.DeferType:
		return WalkDeferDef(tt)
	case *schema.FloatType:
		return WalkFloatDef(tt)
	case *schema.IntType:
		return WalkIntDef(tt)
	case *schema.PointerType:
		return WalkPointerDef(tt)
	case *schema.SliceType:
		return WalkSliceDef(tt)
	case *schema.StringType:
		return WalkStringDef(tt)
	case *schema.StructType:
		return WalkStructDef(tt)
	case *schema.UnionType:
		return WalkUnionDef(tt)
	}
	return nil, fmt.Errorf("No such type %T", t)
}

func WalkTypeSize(t schema.Type, target string) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ByteType:
		return WalkByteSize(tt, target)
	case *schema.DeferType:
		return WalkDeferSize(tt, target)
	case *schema.FloatType:
		return WalkFloatSize(tt, target)
	case *schema.IntType:
		return WalkIntSize(tt, target)
	case *schema.PointerType:
		return WalkPointerSize(tt, target)
	case *schema.SliceType:
		return WalkSliceSize(tt, target)
	case *schema.StringType:
		return WalkStringSize(tt, target)
	case *schema.StructType:
		return WalkStructSize(tt, target)
	case *schema.UnionType:
		return WalkUnionSize(tt, target)
	}
	return nil, fmt.Errorf("No such type %T", t)
}

func WalkTypeMarshal(t schema.Type, target string) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ByteType:
		return WalkByteMarshal(tt, target)
	case *schema.DeferType:
		return WalkDeferMarshal(tt, target)
	case *schema.FloatType:
		return WalkFloatMarshal(tt, target)
	case *schema.IntType:
		return WalkIntMarshal(tt, target)
	case *schema.PointerType:
		return WalkPointerMarshal(tt, target)
	case *schema.SliceType:
		return WalkSliceMarshal(tt, target)
	case *schema.StringType:
		return WalkStringMarshal(tt, target)
	case *schema.StructType:
		return WalkStructMarshal(tt, target)
	case *schema.UnionType:
		return WalkUnionMarshal(tt, target)
	}
	return nil, fmt.Errorf("No such type %T", t)
}

func WalkTypeUnmarshal(t schema.Type, target string) (*StringBuilder, error) {
	switch tt := t.(type) {
	case *schema.ByteType:
		return WalkByteUnmarshal(tt, target)
	case *schema.DeferType:
		return WalkDeferUnmarshal(tt, target)
	case *schema.FloatType:
		return WalkFloatUnmarshal(tt, target)
	case *schema.IntType:
		return WalkIntUnmarshal(tt, target)
	case *schema.PointerType:
		return WalkPointerUnmarshal(tt, target)
	case *schema.SliceType:
		return WalkSliceUnmarshal(tt, target)
	case *schema.StringType:
		return WalkStringUnmarshal(tt, target)
	case *schema.StructType:
		return WalkStructUnmarshal(tt, target)
	case *schema.UnionType:
		return WalkUnionUnmarshal(tt, target)
	}
	return nil, fmt.Errorf("No such type %T", t)
}
