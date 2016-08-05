// This file includes tests for generated types in int.schema.golden.go.

package testdata

import (
	"reflect"
	"testing"
)

var tests = []int64{
	-1 << 63,
	-1<<63 + 1,
	-1,
	0,
	1,
	2,
	10,
	20,
	63,
	64,
	65,
	127,
	128,
	129,
	255,
	256,
	257,
	1<<63 - 1,
}

func testInt64(t *testing.T, x int64) {
	want := Ints{Vint64: x, Int64: x}
	buf, err := want.Marshal(nil)
	if err != nil {
		t.Errorf("%d: Marshal failed: %v", x, err)
		return
	}
	got := Ints{}
	n, err := got.Unmarshal(buf)
	if err != nil {
		t.Errorf("%d: Unmarshal failed: %v", x, err)
		return
	}
	if int(n) != len(buf) {
		t.Errorf("%d: Unmarshal and Marshal report different # bytes: %d vs. %d", x, n, len(buf))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%d: Unmarshal = %d; want %d", got, want)
	}
}

func TestInt64(t *testing.T) {
	for _, x := range tests {
		testInt64(t, x)
		testInt64(t, -x)
	}
}

func testUint64(t *testing.T, x uint64) {
	want := Ints{Vuint64: x, Uint64: x}
	buf, err := want.Marshal(nil)
	if err != nil {
		t.Errorf("%d: Marshal failed: %v", x, err)
		return
	}
	got := Ints{}
	n, err := got.Unmarshal(buf)
	if err != nil {
		t.Errorf("%d: Unmarshal failed: %v", x, err)
		return
	}
	if int(n) != len(buf) {
		t.Errorf("%d: Unmarshal and Marshal report different # bytes: %d vs. %d", x, n, len(buf))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%d: Unmarshal = %d; want %d", got, want)
	}
}

func TestUint64(t *testing.T) {
	for _, x := range tests {
		testUint64(t, uint64(x))
	}
	for x := uint64(0x7); x != 0; x <<= 1 {
		testUint64(t, x)
	}
}
