package testdata

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Array struct {
	A [3]int32
	B [3]Nested
	C []int32
	D []Nested
}

func (d *Array) Size() (s uint64) {

	{
		for k := range d.A {
			_ = k // make compiler happy in case k is unused

			s += 4

		}
	}
	{
		for k := range d.B {
			_ = k // make compiler happy in case k is unused

			{
				s += d.B[k].Size()
			}

		}
	}
	{
		l := uint64(len(d.C))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 4 * l

	}
	{
		l := uint64(len(d.D))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k := range d.D {

			{
				s += d.D[k].Size()
			}

		}

	}
	return
}
func (d *Array) Marshal(buf []byte) ([]byte, error) {
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
		for k := range d.A {

			{

				buf[i+0+0] = byte(d.A[k] >> 0)

				buf[i+1+0] = byte(d.A[k] >> 8)

				buf[i+2+0] = byte(d.A[k] >> 16)

				buf[i+3+0] = byte(d.A[k] >> 24)

			}

			i += 4

		}
	}
	{
		for k := range d.B {

			{
				nbuf, err := d.B[k].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	{
		l := uint64(len(d.C))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k := range d.C {

			{

				buf[i+0+0] = byte(d.C[k] >> 0)

				buf[i+1+0] = byte(d.C[k] >> 8)

				buf[i+2+0] = byte(d.C[k] >> 16)

				buf[i+3+0] = byte(d.C[k] >> 24)

			}

			i += 4

		}
	}
	{
		l := uint64(len(d.D))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k := range d.D {

			{
				nbuf, err := d.D[k].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *Array) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		for k := range d.A {

			{

				d.A[k] = 0 | (int32(buf[i+0+0]) << 0) | (int32(buf[i+1+0]) << 8) | (int32(buf[i+2+0]) << 16) | (int32(buf[i+3+0]) << 24)

			}

			i += 4

		}
	}
	{
		for k := range d.B {

			{
				ni, err := d.B[k].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.C)) >= l {
			d.C = d.C[:l]
		} else {
			d.C = make([]int32, l)
		}
		for k := range d.C {

			{

				d.C[k] = 0 | (int32(buf[i+0+0]) << 0) | (int32(buf[i+1+0]) << 8) | (int32(buf[i+2+0]) << 16) | (int32(buf[i+3+0]) << 24)

			}

			i += 4

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.D)) >= l {
			d.D = d.D[:l]
		} else {
			d.D = make([]Nested, l)
		}
		for k := range d.D {

			{
				ni, err := d.D[k].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

type Nested struct {
	A [16]string
	B []string
}

func (d *Nested) Size() (s uint64) {

	{
		for k := range d.A {
			_ = k // make compiler happy in case k is unused

			{
				l := uint64(len(d.A[k]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}
	}
	{
		l := uint64(len(d.B))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k := range d.B {

			{
				l := uint64(len(d.B[k]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	return
}
func (d *Nested) Marshal(buf []byte) ([]byte, error) {
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
		for k := range d.A {

			{
				l := uint64(len(d.A[k]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				copy(buf[i+0:], d.A[k])
				i += l
			}

		}
	}
	{
		l := uint64(len(d.B))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k := range d.B {

			{
				l := uint64(len(d.B[k]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				copy(buf[i+0:], d.B[k])
				i += l
			}

		}
	}
	return buf[:i+0], nil
}

func (d *Nested) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		for k := range d.A {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.A[k] = string(buf[i+0 : i+0+l])
				i += l
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.B)) >= l {
			d.B = d.B[:l]
		} else {
			d.B = make([]string, l)
		}
		for k := range d.B {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.B[k] = string(buf[i+0 : i+0+l])
				i += l
			}

		}
	}
	return i + 0, nil
}
