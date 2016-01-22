package main

import (
	"fmt"
	"testing"
)

func TestGencodeSize(t *testing.T) {
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
	buf, _ := p.Marshal(nil)
	fmt.Printf("Gencode encoded size: %v\n", len(buf))
}

func BenchmarkGencodeSerialize(b *testing.B) {
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
		p.Marshal(nil)
	}
}

func BenchmarkGencodeDeserialize(b *testing.B) {
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
	buf, _ := p.Marshal(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Unmarshal(buf)
	}
}

func BenchmarkFixedGencodeSerialize(b *testing.B) {
	p := Fixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf, _ := p.Marshal(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Marshal(buf)
	}
}

func BenchmarkFixedGencodeDeserialize(b *testing.B) {
	p := Fixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf, _ := p.Marshal(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Unmarshal(buf)
	}
}
