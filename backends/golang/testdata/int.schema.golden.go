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

type Ints struct {
	Vint8   int8
	Vint16  int16
	Vint32  int32
	Vint64  int64
	Vuint8  uint8
	Vuint16 uint16
	Vuint32 uint32
	Vuint64 uint64
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
}

func (d *Ints) MarshalSize() (s uint64) {

	{

		t := uint8(d.Vint8)
		t <<= 1
		if d.Vint8 < 0 {
			t = ^t
		}
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := uint16(d.Vint16)
		t <<= 1
		if d.Vint16 < 0 {
			t = ^t
		}
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := uint32(d.Vint32)
		t <<= 1
		if d.Vint32 < 0 {
			t = ^t
		}
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := uint64(d.Vint64)
		t <<= 1
		if d.Vint64 < 0 {
			t = ^t
		}
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Vuint8
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Vuint16
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Vuint32
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Vuint64
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	s += 30
	return
}
func (d *Ints) Marshal(buf []byte) ([]byte, error) {
	size := d.MarshalSize()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		t := uint8(d.Vint8)

		t <<= 1
		if d.Vint8 < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint16(d.Vint16)

		t <<= 1
		if d.Vint16 < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint32(d.Vint32)

		t <<= 1
		if d.Vint32 < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint64(d.Vint64)

		t <<= 1
		if d.Vint64 < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint8(d.Vuint8)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint16(d.Vuint16)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint32(d.Vuint32)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint64(d.Vuint64)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		buf[i+0+0] = byte(d.Int8 >> 0)

	}
	{

		buf[i+0+1] = byte(d.Int16 >> 0)

		buf[i+1+1] = byte(d.Int16 >> 8)

	}
	{

		buf[i+0+3] = byte(d.Int32 >> 0)

		buf[i+1+3] = byte(d.Int32 >> 8)

		buf[i+2+3] = byte(d.Int32 >> 16)

		buf[i+3+3] = byte(d.Int32 >> 24)

	}
	{

		buf[i+0+7] = byte(d.Int64 >> 0)

		buf[i+1+7] = byte(d.Int64 >> 8)

		buf[i+2+7] = byte(d.Int64 >> 16)

		buf[i+3+7] = byte(d.Int64 >> 24)

		buf[i+4+7] = byte(d.Int64 >> 32)

		buf[i+5+7] = byte(d.Int64 >> 40)

		buf[i+6+7] = byte(d.Int64 >> 48)

		buf[i+7+7] = byte(d.Int64 >> 56)

	}
	{

		buf[i+0+15] = byte(d.Uint8 >> 0)

	}
	{

		buf[i+0+16] = byte(d.Uint16 >> 0)

		buf[i+1+16] = byte(d.Uint16 >> 8)

	}
	{

		buf[i+0+18] = byte(d.Uint32 >> 0)

		buf[i+1+18] = byte(d.Uint32 >> 8)

		buf[i+2+18] = byte(d.Uint32 >> 16)

		buf[i+3+18] = byte(d.Uint32 >> 24)

	}
	{

		buf[i+0+22] = byte(d.Uint64 >> 0)

		buf[i+1+22] = byte(d.Uint64 >> 8)

		buf[i+2+22] = byte(d.Uint64 >> 16)

		buf[i+3+22] = byte(d.Uint64 >> 24)

		buf[i+4+22] = byte(d.Uint64 >> 32)

		buf[i+5+22] = byte(d.Uint64 >> 40)

		buf[i+6+22] = byte(d.Uint64 >> 48)

		buf[i+7+22] = byte(d.Uint64 >> 56)

	}
	return buf[:i+30], nil
}

func (d *Ints) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		bs := uint8(7)
		t := uint8(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint8(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vint8 = int8(t >> 1)
		if t&1 != 0 {
			d.Vint8 = ^d.Vint8
		}

	}
	{

		bs := uint8(7)
		t := uint16(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint16(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vint16 = int16(t >> 1)
		if t&1 != 0 {
			d.Vint16 = ^d.Vint16
		}

	}
	{

		bs := uint8(7)
		t := uint32(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint32(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vint32 = int32(t >> 1)
		if t&1 != 0 {
			d.Vint32 = ^d.Vint32
		}

	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vint64 = int64(t >> 1)
		if t&1 != 0 {
			d.Vint64 = ^d.Vint64
		}

	}
	{

		bs := uint8(7)
		t := uint8(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint8(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vuint8 = t

	}
	{

		bs := uint8(7)
		t := uint16(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint16(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vuint16 = t

	}
	{

		bs := uint8(7)
		t := uint32(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint32(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vuint32 = t

	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Vuint64 = t

	}
	{

		d.Int8 = 0 | (int8(buf[i+0+0]) << 0)

	}
	{

		d.Int16 = 0 | (int16(buf[i+0+1]) << 0) | (int16(buf[i+1+1]) << 8)

	}
	{

		d.Int32 = 0 | (int32(buf[i+0+3]) << 0) | (int32(buf[i+1+3]) << 8) | (int32(buf[i+2+3]) << 16) | (int32(buf[i+3+3]) << 24)

	}
	{

		d.Int64 = 0 | (int64(buf[i+0+7]) << 0) | (int64(buf[i+1+7]) << 8) | (int64(buf[i+2+7]) << 16) | (int64(buf[i+3+7]) << 24) | (int64(buf[i+4+7]) << 32) | (int64(buf[i+5+7]) << 40) | (int64(buf[i+6+7]) << 48) | (int64(buf[i+7+7]) << 56)

	}
	{

		d.Uint8 = 0 | (uint8(buf[i+0+15]) << 0)

	}
	{

		d.Uint16 = 0 | (uint16(buf[i+0+16]) << 0) | (uint16(buf[i+1+16]) << 8)

	}
	{

		d.Uint32 = 0 | (uint32(buf[i+0+18]) << 0) | (uint32(buf[i+1+18]) << 8) | (uint32(buf[i+2+18]) << 16) | (uint32(buf[i+3+18]) << 24)

	}
	{

		d.Uint64 = 0 | (uint64(buf[i+0+22]) << 0) | (uint64(buf[i+1+22]) << 8) | (uint64(buf[i+2+22]) << 16) | (uint64(buf[i+3+22]) << 24) | (uint64(buf[i+4+22]) << 32) | (uint64(buf[i+5+22]) << 40) | (uint64(buf[i+6+22]) << 48) | (uint64(buf[i+7+22]) << 56)

	}
	return i + 30, nil
}
