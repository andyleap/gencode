package main

import (
	"io"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
)

type Person struct {
	Name   string
	Age    uint8
	Height float64
}

func (d *Person) Size() (s uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 9
	return
}
func (d *Person) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.Name))

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
		copy(buf[i:], d.Name)
		i += l
	}
	{

		buf[i+0+0] = byte(d.Age >> 0)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Height)))

		buf[i+0+1] = byte(v >> 0)

		buf[i+1+1] = byte(v >> 8)

		buf[i+2+1] = byte(v >> 16)

		buf[i+3+1] = byte(v >> 24)

		buf[i+4+1] = byte(v >> 32)

		buf[i+5+1] = byte(v >> 40)

		buf[i+6+1] = byte(v >> 48)

		buf[i+7+1] = byte(v >> 56)

	}
	return buf[:i+9], nil
}

func (d *Person) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i : i+l])
		i += l
	}
	{

		d.Age = 0 | (uint8(buf[i+0+0]) << 0)

	}
	{

		v := 0 | (uint64(buf[i+0+1]) << 0) | (uint64(buf[i+1+1]) << 8) | (uint64(buf[i+2+1]) << 16) | (uint64(buf[i+3+1]) << 24) | (uint64(buf[i+4+1]) << 32) | (uint64(buf[i+5+1]) << 40) | (uint64(buf[i+6+1]) << 48) | (uint64(buf[i+7+1]) << 56)
		d.Height = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 9, nil
}

type Group struct {
	Name    string
	Members []Person
}

func (d *Group) FramedSize() (s uint64, us uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Members))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		for k := range d.Members {

			{
				s += d.Members[k].Size()
			}

		}
	}
	l := s
	us = s

	{

		t := l
		for t >= 0x80 {
			t <<= 7
			s++
		}
		s++

	}
	return
}
func (d *Group) Size() (s uint64) {
	s, _ = d.FramedSize()
	return
}

func (d *Group) Marshal(buf []byte) ([]byte, error) {
	size, usize := d.FramedSize()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		t := uint64(usize)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{
		l := uint64(len(d.Name))

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
		copy(buf[i:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Members))

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
		for k := range d.Members {

			{
				nbuf, err := d.Members[k].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *Group) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	usize := uint64(0)

	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		usize = t

	}
	if usize > uint64(len(buf))+i {
		return 0, io.EOF
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i : i+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Members)) >= l {
			d.Members = d.Members[:l]
		} else {
			d.Members = make([]Person, l)
		}
		for k := range d.Members {

			{
				ni, err := d.Members[k].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

func (d *Group) Serialize(w io.Writer) error {
	buf, err := d.Marshal(nil)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (d *Group) Deserialize(r io.Reader) error {
	size := uint64(0)
	sbuf := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	bs := uint8(0)
	i := uint64(0)
	for sbuf[i]&0x80 == 0x80 {
		_, err := r.Read(sbuf[i : i+1])
		if err != nil {
			return err
		}
		size |= uint64(sbuf[i]&0x7F) << bs
		bs += 7
		i++
	}
	buf := make([]byte, size+i)
	copy(buf, sbuf[0:i])
	n := uint64(i)
	size += i
	var err error
	for n < size && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += uint64(nn)
	}
	if err != nil {
		return err
	}
	_, err = d.Unmarshal(buf)
	if err != nil {
		return err
	}
	return nil
}

type A struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int64
	Spouse   uint8
	Money    float64
}

func (d *A) Size() (s uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Phone))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 25
	return
}
func (d *A) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.Name))

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
		copy(buf[i:], d.Name)
		i += l
	}
	{

		buf[i+0+0] = byte(d.BirthDay >> 0)

		buf[i+1+0] = byte(d.BirthDay >> 8)

		buf[i+2+0] = byte(d.BirthDay >> 16)

		buf[i+3+0] = byte(d.BirthDay >> 24)

		buf[i+4+0] = byte(d.BirthDay >> 32)

		buf[i+5+0] = byte(d.BirthDay >> 40)

		buf[i+6+0] = byte(d.BirthDay >> 48)

		buf[i+7+0] = byte(d.BirthDay >> 56)

	}
	{
		l := uint64(len(d.Phone))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i:], d.Phone)
		i += l
	}
	{

		buf[i+0+8] = byte(d.Siblings >> 0)

		buf[i+1+8] = byte(d.Siblings >> 8)

		buf[i+2+8] = byte(d.Siblings >> 16)

		buf[i+3+8] = byte(d.Siblings >> 24)

		buf[i+4+8] = byte(d.Siblings >> 32)

		buf[i+5+8] = byte(d.Siblings >> 40)

		buf[i+6+8] = byte(d.Siblings >> 48)

		buf[i+7+8] = byte(d.Siblings >> 56)

	}
	{

		buf[i+0+16] = byte(d.Spouse >> 0)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Money)))

		buf[i+0+17] = byte(v >> 0)

		buf[i+1+17] = byte(v >> 8)

		buf[i+2+17] = byte(v >> 16)

		buf[i+3+17] = byte(v >> 24)

		buf[i+4+17] = byte(v >> 32)

		buf[i+5+17] = byte(v >> 40)

		buf[i+6+17] = byte(v >> 48)

		buf[i+7+17] = byte(v >> 56)

	}
	return buf[:i+25], nil
}

func (d *A) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i : i+l])
		i += l
	}
	{

		d.BirthDay = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Phone = string(buf[i : i+l])
		i += l
	}
	{

		d.Siblings = 0 | (int64(buf[i+0+8]) << 0) | (int64(buf[i+1+8]) << 8) | (int64(buf[i+2+8]) << 16) | (int64(buf[i+3+8]) << 24) | (int64(buf[i+4+8]) << 32) | (int64(buf[i+5+8]) << 40) | (int64(buf[i+6+8]) << 48) | (int64(buf[i+7+8]) << 56)

	}
	{

		d.Spouse = 0 | (uint8(buf[i+0+16]) << 0)

	}
	{

		v := 0 | (uint64(buf[i+0+17]) << 0) | (uint64(buf[i+1+17]) << 8) | (uint64(buf[i+2+17]) << 16) | (uint64(buf[i+3+17]) << 24) | (uint64(buf[i+4+17]) << 32) | (uint64(buf[i+5+17]) << 40) | (uint64(buf[i+6+17]) << 48) | (uint64(buf[i+7+17]) << 56)
		d.Money = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 25, nil
}
