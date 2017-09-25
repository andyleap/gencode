package bench

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func BenchmarkFixedBinarySerialize(b *testing.B) {
	p := Fixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf := bytes.NewBuffer(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.Write(buf, binary.LittleEndian, p)
		buf.Reset()
	}
}

func BenchmarkFixedBinaryDeserialize(b *testing.B) {
	p := Fixed{
		A: -5,
		B: 6,
		C: 6.7,
		D: 12.65,
	}
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, p)
	rbuf := bytes.NewReader(buf.Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.Read(rbuf, binary.LittleEndian, &p)
		rbuf.Seek(0, 0)
	}
}
