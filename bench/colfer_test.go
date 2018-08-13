package bench

import (
	"fmt"
	"testing"
)

func TestColferSize(t *testing.T) {
	p := ColferGroup{
		Name: "test",
		Members: []*ColferPerson{
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
	n, err := p.MarshalLen()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Colfer encoded size: %v\n", n)
}

func BenchmarkColferSerialize(b *testing.B) {
	p := ColferGroup{
		Name: "test",
		Members: []*ColferPerson{
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
		var buf [100]byte
		p.MarshalTo(buf[:])
	}
}

func BenchmarkColferDeserialize(b *testing.B) {
	p := ColferGroup{
		Name: "test",
		Members: []*ColferPerson{
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
	buf, err := p.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Unmarshal(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFixedColferSerialize(b *testing.B) {
	p := ColferFixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf, err := p.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.MarshalTo(buf)
	}
}

func BenchmarkFixedColferDeserialize(b *testing.B) {
	p := ColferFixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf, err := p.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Unmarshal(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}
