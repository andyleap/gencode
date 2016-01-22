package main

import (
	"io"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
)

type Fixed struct {
	A int64
	B uint32
	C float32
	D float64
}

func (d *Fixed) Size() (s uint64) {

	s += 24
	return
}
func (d *Fixed) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.A >> 0)

		buf[1+0] = byte(d.A >> 8)

		buf[2+0] = byte(d.A >> 16)

		buf[3+0] = byte(d.A >> 24)

		buf[4+0] = byte(d.A >> 32)

		buf[5+0] = byte(d.A >> 40)

		buf[6+0] = byte(d.A >> 48)

		buf[7+0] = byte(d.A >> 56)

	}
	{

		buf[0+8] = byte(d.B >> 0)

		buf[1+8] = byte(d.B >> 8)

		buf[2+8] = byte(d.B >> 16)

		buf[3+8] = byte(d.B >> 24)

	}
	{

		v := *(*uint32)(unsafe.Pointer(&(d.C)))

		buf[0+12] = byte(v >> 0)

		buf[1+12] = byte(v >> 8)

		buf[2+12] = byte(v >> 16)

		buf[3+12] = byte(v >> 24)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.D)))

		buf[0+16] = byte(v >> 0)

		buf[1+16] = byte(v >> 8)

		buf[2+16] = byte(v >> 16)

		buf[3+16] = byte(v >> 24)

		buf[4+16] = byte(v >> 32)

		buf[5+16] = byte(v >> 40)

		buf[6+16] = byte(v >> 48)

		buf[7+16] = byte(v >> 56)

	}
	return buf[:i+24], nil
}

func (d *Fixed) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.A = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	{

		d.B = 0 | (uint32(buf[0+8]) << 0) | (uint32(buf[1+8]) << 8) | (uint32(buf[2+8]) << 16) | (uint32(buf[3+8]) << 24)

	}
	{

		v := 0 | (uint32(buf[0+12]) << 0) | (uint32(buf[1+12]) << 8) | (uint32(buf[2+12]) << 16) | (uint32(buf[3+12]) << 24)
		d.C = *(*float32)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[0+16]) << 0) | (uint64(buf[1+16]) << 8) | (uint64(buf[2+16]) << 16) | (uint64(buf[3+16]) << 24) | (uint64(buf[4+16]) << 32) | (uint64(buf[5+16]) << 40) | (uint64(buf[6+16]) << 48) | (uint64(buf[7+16]) << 56)
		d.D = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 24, nil
}
