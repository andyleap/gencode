package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONSize(t *testing.T) {
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
	buf, _ := json.Marshal(p)
	fmt.Printf("JSON encoded size: %v\n", len(buf))
}

func BenchmarkJSONSerialize(b *testing.B) {
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(p)
	}
}

func BenchmarkJSONDeserialize(b *testing.B) {
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
	buf, _ := json.Marshal(p)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(buf, &p)
	}
}
