package main

import (
	"math"
)

var (
	_ = math.Float64frombits
)

type Fixed struct {
	A int64
	B uint32
	C float32
	D float64
}

func (d *Fixed) Size() (s uint64) {

	{
		s += 8
	}
	{
		s += 4
	}
	{
		s += 4
	}
	{
		s += 8
	}
	return
}

func (d *Fixed) Marshal(buf []byte) ([]byte, error) {
	{
		size := d.Size()
		if uint64(cap(buf)) >= d.Size() {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[i+0] = byte(d.A >> 0)

		buf[i+1] = byte(d.A >> 8)

		buf[i+2] = byte(d.A >> 16)

		buf[i+3] = byte(d.A >> 24)

		buf[i+4] = byte(d.A >> 32)

		buf[i+5] = byte(d.A >> 40)

		buf[i+6] = byte(d.A >> 48)

		buf[i+7] = byte(d.A >> 56)

		i += 8

	}
	{

		buf[i+0] = byte(d.B >> 0)

		buf[i+1] = byte(d.B >> 8)

		buf[i+2] = byte(d.B >> 16)

		buf[i+3] = byte(d.B >> 24)

		i += 4

	}
	{
		v := math.Float32bits(d.C)

		{

			buf[i+0] = byte(v >> 0)

			buf[i+1] = byte(v >> 8)

			buf[i+2] = byte(v >> 16)

			buf[i+3] = byte(v >> 24)

			i += 4

		}
	}
	{
		v := math.Float64bits(d.D)

		{

			buf[i+0] = byte(v >> 0)

			buf[i+1] = byte(v >> 8)

			buf[i+2] = byte(v >> 16)

			buf[i+3] = byte(v >> 24)

			buf[i+4] = byte(v >> 32)

			buf[i+5] = byte(v >> 40)

			buf[i+6] = byte(v >> 48)

			buf[i+7] = byte(v >> 56)

			i += 8

		}
	}
	return buf[:i], nil
}

func (d *Fixed) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.A = 0

		d.A |= int64(buf[i+0]) << 0

		d.A |= int64(buf[i+1]) << 8

		d.A |= int64(buf[i+2]) << 16

		d.A |= int64(buf[i+3]) << 24

		d.A |= int64(buf[i+4]) << 32

		d.A |= int64(buf[i+5]) << 40

		d.A |= int64(buf[i+6]) << 48

		d.A |= int64(buf[i+7]) << 56

		i += 8

	}
	{

		d.B = 0

		d.B |= uint32(buf[i+0]) << 0

		d.B |= uint32(buf[i+1]) << 8

		d.B |= uint32(buf[i+2]) << 16

		d.B |= uint32(buf[i+3]) << 24

		i += 4

	}
	{
		var v uint32

		{

			v = 0

			v |= uint32(buf[i+0]) << 0

			v |= uint32(buf[i+1]) << 8

			v |= uint32(buf[i+2]) << 16

			v |= uint32(buf[i+3]) << 24

			i += 4

		}
		d.C = math.Float32frombits(v)
	}
	{
		var v uint64

		{

			v = 0

			v |= uint64(buf[i+0]) << 0

			v |= uint64(buf[i+1]) << 8

			v |= uint64(buf[i+2]) << 16

			v |= uint64(buf[i+3]) << 24

			v |= uint64(buf[i+4]) << 32

			v |= uint64(buf[i+5]) << 40

			v |= uint64(buf[i+6]) << 48

			v |= uint64(buf[i+7]) << 56

			i += 8

		}
		d.D = math.Float64frombits(v)
	}
	return i, nil
}
