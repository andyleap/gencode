package main

import (
	"fmt"
	"testing"
)

func TestMSGPSize(t *testing.T) {
	p := Group{
		Name: "test",
		Members: []Person{
			{
				Name:   "John",
				Age:    21,
				Height: 5.9,
			},
			{
				Name:   "Tom",
				Age:    23,
				Height: 5.8,
			},
			{
				Name:   "Alan",
				Age:    24,
				Height: 6,
			},
		},
	}
	buf, _ := p.MarshalMsg(nil)
	fmt.Printf("MSGP encoded size: %v\n", len(buf))
}

func BenchmarkMSGPSerialize(b *testing.B) {
	p := Group{
		Name: "test",
		Members: []Person{
			{
				Name:   "John",
				Age:    21,
				Height: 5.9,
			},
			{
				Name:   "Tom",
				Age:    23,
				Height: 5.8,
			},
			{
				Name:   "Alan",
				Age:    24,
				Height: 6,
			},
		},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.MarshalMsg(nil)
	}
}

func BenchmarkMSGPDeserialize(b *testing.B) {
	p := Group{
		Name: "test",
		Members: []Person{
			{
				Name:   "John",
				Age:    21,
				Height: 5.9,
			},
			{
				Name:   "Tom",
				Age:    23,
				Height: 5.8,
			},
			{
				Name:   "Alan",
				Age:    24,
				Height: 6,
			},
		},
	}
	buf, _ := p.MarshalMsg(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.UnmarshalMsg(buf)
	}
}
