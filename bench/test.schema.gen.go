package main
import (
	"encoding/binary"
	"io"
)

var (
	_ = io.ReadFull
	_ = binary.Write
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
		err := binary.Write(w, binary.LittleEndian, d.Height)
		if err != nil {
			return err
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
		err := binary.Read(r, binary.LittleEndian, &d.Height)
		if err != nil {
			return err
		}
	}
	return nil
}

type Group struct {
	Name string
	Members []Person
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
		err := s.Serialize(w)
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
		d.Members = make([]Person, t)
		for k := range d.Members {
			var s Person
			
	{
		err := s.Deserialize(r)
		if err != nil {
			return err
		}
	}
			d.Members[k] = s
		}
	}
	return nil
}

