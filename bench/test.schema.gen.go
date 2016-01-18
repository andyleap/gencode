package main

import (
	"io"
	"math"
)

var (
	_ = math.Float64frombits
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
	{
		s += 1
	}
	{
		s += 8
	}
	return
}
func (d *Person) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= d.Size() {
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
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
			i++

		}
		copy(buf[i:], d.Name)
		i += l
	}
	{

		buf[i+0] = byte(d.Age >> 0)

		i += 1

	}
	{
		v := math.Float64bits(d.Height)

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

func (d *Person) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i : i+l])
		i += l
	}
	{

		d.Age = 0

		d.Age |= uint8(buf[i+0]) << 0

		i += 1

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
		d.Height = math.Float64frombits(v)
	}
	return i, nil
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
		if uint64(cap(buf)) >= d.Size() {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		t := uint64(usize)

		for t >= 0x80 {
			buf[i] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i] = byte(t)
		i++

	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
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
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
			i++

		}
		for k := range d.Members {

			{
				nbuf, err := d.Members[k].Marshal(buf[i:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
		}
	}
	return buf[:i], nil
}

func (d *Group) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)
	usize := uint64(0)

	{

		bs := uint8(7)
		t := uint64(buf[i] & 0x7F)
		for buf[i]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i]&0x7F) << bs
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
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
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
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
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
				ni, err := d.Members[k].Unmarshal(buf[i:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
		}
	}
	return i, nil
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
		s += 8
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
	{
		s += 8
	}
	{
		s += 1
	}
	{
		s += 8
	}
	return
}
func (d *A) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= d.Size() {
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
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
			i++

		}
		copy(buf[i:], d.Name)
		i += l
	}
	{

		buf[i+0] = byte(d.BirthDay >> 0)

		buf[i+1] = byte(d.BirthDay >> 8)

		buf[i+2] = byte(d.BirthDay >> 16)

		buf[i+3] = byte(d.BirthDay >> 24)

		buf[i+4] = byte(d.BirthDay >> 32)

		buf[i+5] = byte(d.BirthDay >> 40)

		buf[i+6] = byte(d.BirthDay >> 48)

		buf[i+7] = byte(d.BirthDay >> 56)

		i += 8

	}
	{
		l := uint64(len(d.Phone))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i] = byte(t)
			i++

		}
		copy(buf[i:], d.Phone)
		i += l
	}
	{

		buf[i+0] = byte(d.Siblings >> 0)

		buf[i+1] = byte(d.Siblings >> 8)

		buf[i+2] = byte(d.Siblings >> 16)

		buf[i+3] = byte(d.Siblings >> 24)

		buf[i+4] = byte(d.Siblings >> 32)

		buf[i+5] = byte(d.Siblings >> 40)

		buf[i+6] = byte(d.Siblings >> 48)

		buf[i+7] = byte(d.Siblings >> 56)

		i += 8

	}
	{

		buf[i+0] = byte(d.Spouse >> 0)

		i += 1

	}
	{
		v := math.Float64bits(d.Money)

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

func (d *A) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i : i+l])
		i += l
	}
	{

		d.BirthDay = 0

		d.BirthDay |= int64(buf[i+0]) << 0

		d.BirthDay |= int64(buf[i+1]) << 8

		d.BirthDay |= int64(buf[i+2]) << 16

		d.BirthDay |= int64(buf[i+3]) << 24

		d.BirthDay |= int64(buf[i+4]) << 32

		d.BirthDay |= int64(buf[i+5]) << 40

		d.BirthDay |= int64(buf[i+6]) << 48

		d.BirthDay |= int64(buf[i+7]) << 56

		i += 8

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i] & 0x7F)
			for buf[i]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Phone = string(buf[i : i+l])
		i += l
	}
	{

		d.Siblings = 0

		d.Siblings |= int64(buf[i+0]) << 0

		d.Siblings |= int64(buf[i+1]) << 8

		d.Siblings |= int64(buf[i+2]) << 16

		d.Siblings |= int64(buf[i+3]) << 24

		d.Siblings |= int64(buf[i+4]) << 32

		d.Siblings |= int64(buf[i+5]) << 40

		d.Siblings |= int64(buf[i+6]) << 48

		d.Siblings |= int64(buf[i+7]) << 56

		i += 8

	}
	{

		d.Spouse = 0

		d.Spouse |= uint8(buf[i+0]) << 0

		i += 1

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
		d.Money = math.Float64frombits(v)
	}
	return i, nil
}
