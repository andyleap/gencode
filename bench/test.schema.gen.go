package main
import (
	"math"
	"io"
)

var (
	_ = io.ReadFull
	_ = math.Float64frombits
)

type Person struct {
	Name string
	Age uint8
	Height float64
}

func (d *Person) Serialize(w io.Writer) error {
	{
		t := uint64(len(d.Name))
		buf := make([]byte, 10)
		i := 0
		for t >= 0x80 {
			buf[i] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i] = byte(t)
		i++
		_, err := w.Write(buf[:i])
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(d.Name))
		if err != nil {
			return err
		}
	}
	{
		buf := make([]byte, 1)
		
		buf[0] = byte(d.Age >> 0)
		
		
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
	{
		v := math.Float64bits(d.Height)
		
	{
		buf := make([]byte, 8)
		
		buf[0] = byte(v >> 0)
		
		buf[1] = byte(v >> 8)
		
		buf[2] = byte(v >> 16)
		
		buf[3] = byte(v >> 24)
		
		buf[4] = byte(v >> 32)
		
		buf[5] = byte(v >> 40)
		
		buf[6] = byte(v >> 48)
		
		buf[7] = byte(v >> 56)
		
		
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
	}
	return nil
}

func (d *Person) Deserialize(r io.Reader) error {
	{
		buf := make([]byte, 1)
		buf[0] = 0x80
		t := uint64(0)
		for buf[0] & 0x80 == 0x80 {
			t <<= 7
			_, err := io.ReadFull(r, buf)
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F)
		}
		buf = make([]byte, t)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		d.Name = string(buf)
	}
	{
		buf := make([]byte, 1)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		
		d.Age |= uint8(buf[0]) << 0
		
	}
	{
		var v uint64
		
	{
		buf := make([]byte, 8)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		
		v |= uint64(buf[0]) << 0
		
		v |= uint64(buf[1]) << 8
		
		v |= uint64(buf[2]) << 16
		
		v |= uint64(buf[3]) << 24
		
		v |= uint64(buf[4]) << 32
		
		v |= uint64(buf[5]) << 40
		
		v |= uint64(buf[6]) << 48
		
		v |= uint64(buf[7]) << 56
		
	}
		d.Height = math.Float64frombits(v)
	}
	return nil
}

type Group struct {
	Name string
	Members []interface{}
}

func (d *Group) Serialize(w io.Writer) error {
	{
		t := uint64(len(d.Name))
		buf := make([]byte, 10)
		i := 0
		for t >= 0x80 {
			buf[i] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i] = byte(t)
		i++
		_, err := w.Write(buf[:i])
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(d.Name))
		if err != nil {
			return err
		}
	}
	{
		t := uint64(len(d.Members))
		buf := make([]byte, 10)
		i := 0
		for t >= 0x80 {
			buf[i] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i] = byte(t)
		i++
		_, err := w.Write(buf[:i])
		if err != nil {
			return err
		}
		for _, s := range d.Members {
			
	{
		var t uint64
		switch s.(type) {
			
		case Person:
			t = 0
			
		case Group:
			t = 1
			
		}
		buf := make([]byte, 10)
		i := 0
		for t >= 0x80 {
			buf[i] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i] = byte(t)
		i++
		_, err := w.Write(buf[:i])
		if err != nil {
			return err
		}
		switch tt := s.(type) {
			
		case Person:
			err = tt.Serialize(w)
			
		case Group:
			err = tt.Serialize(w)
			
		}
		if err != nil {
			return err
		}
	}
		}
	}
	return nil
}

func (d *Group) Deserialize(r io.Reader) error {
	{
		buf := make([]byte, 1)
		buf[0] = 0x80
		t := uint64(0)
		for buf[0] & 0x80 == 0x80 {
			t <<= 7
			_, err := io.ReadFull(r, buf)
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F)
		}
		buf = make([]byte, t)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		d.Name = string(buf)
	}
	{
		buf := make([]byte, 1)
		buf[0] = 0x80
		t := uint64(0)
		for buf[0] & 0x80 == 0x80 {
			t <<= 7
			_, err := io.ReadFull(r, buf)
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F)
		}
		d.Members = make([]interface{}, t)
		for k := range d.Members {
			var s interface{}
			
	{
		buf := make([]byte, 1)
		buf[0] = 0x80
		t := uint64(0)
		for buf[0] & 0x80 == 0x80 {
			t <<= 7
			_, err := io.ReadFull(r, buf)
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F)
		}
		switch t {
			
		case 0:
			tt := Person{}
			err := tt.Deserialize(r)
			if err != nil {
				return err
			}
			s = tt
			
		case 1:
			tt := Group{}
			err := tt.Deserialize(r)
			if err != nil {
				return err
			}
			s = tt
			
		}
		
	}
			d.Members[k] = s
		}
	}
	return nil
}

