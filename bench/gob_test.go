package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

func TestGobSize(t *testing.T) {
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
	buf := &bytes.Buffer{}
	e := gob.NewEncoder(buf)
	e.Encode(p)
	fmt.Printf("GOB encoded size: %v\n", len(buf.Bytes()))
	e.Encode(p)
	e.Encode(p)
	e.Encode(p)
	pos := buf.Len()
	e.Encode(p)
	fmt.Printf("GOB Stream encoded size: %v\n", buf.Len()-pos)
}

func BenchmarkGobSerialize(b *testing.B) {
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
	buf := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		e := gob.NewEncoder(buf)
		e.Encode(p)
	}
}

func BenchmarkGobDeserialize(b *testing.B) {
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
	buf := &bytes.Buffer{}
	e := gob.NewEncoder(buf)
	e.Encode(p)
	rbuf := bytes.NewReader(buf.Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbuf.Seek(0, 0)
		d := gob.NewDecoder(rbuf)
		d.Decode(&p)
	}
}

func BenchmarkGobStreamSerialize(b *testing.B) {
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
	buf := &bytes.Buffer{}
	e := gob.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		e.Encode(p)
	}
}

func BenchmarkGobStreamDeserialize(b *testing.B) {
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
	buf := &bytes.Buffer{}
	e := gob.NewEncoder(buf)
	e.Encode(p)
	e.Encode(p)
	e.Encode(p)
	e.Encode(p)
	e.Encode(p)
	pos := int64(buf.Len())
	e.Encode(p)
	rbuf := bytes.NewReader(buf.Bytes())
	d := gob.NewDecoder(rbuf)
	d.Decode(&p)
	d.Decode(&p)
	d.Decode(&p)
	d.Decode(&p)
	d.Decode(&p)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbuf.Seek(pos, 0)
		err := d.Decode(&p)
		if err != nil {
			b.Log(i)
			b.Fatal(err)
		}
	}
}
