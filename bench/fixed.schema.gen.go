package main
import (
	"math"
	"io"
)

var (
	_ = io.ReadFull
	_ = math.Float64frombits
)

type Fixed struct {
	A int64
	B uint32
	C float32
	D float64
}

func (d *Fixed) Serialize(w io.Writer) error {
	{
		buf := make([]byte, 8)
		
		buf[0] = byte(d.A >> 0)
		
		buf[1] = byte(d.A >> 8)
		
		buf[2] = byte(d.A >> 16)
		
		buf[3] = byte(d.A >> 24)
		
		buf[4] = byte(d.A >> 32)
		
		buf[5] = byte(d.A >> 40)
		
		buf[6] = byte(d.A >> 48)
		
		buf[7] = byte(d.A >> 56)
		
		
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
	{
		buf := make([]byte, 4)
		
		buf[0] = byte(d.B >> 0)
		
		buf[1] = byte(d.B >> 8)
		
		buf[2] = byte(d.B >> 16)
		
		buf[3] = byte(d.B >> 24)
		
		
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
	{
		v := math.Float32bits(d.C)
		
	{
		buf := make([]byte, 4)
		
		buf[0] = byte(v >> 0)
		
		buf[1] = byte(v >> 8)
		
		buf[2] = byte(v >> 16)
		
		buf[3] = byte(v >> 24)
		
		
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
	}
	{
		v := math.Float64bits(d.D)
		
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

func (d *Fixed) Deserialize(r io.Reader) error {
	{
		buf := make([]byte, 8)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		
		d.A |= int64(buf[0]) << 0
		
		d.A |= int64(buf[1]) << 8
		
		d.A |= int64(buf[2]) << 16
		
		d.A |= int64(buf[3]) << 24
		
		d.A |= int64(buf[4]) << 32
		
		d.A |= int64(buf[5]) << 40
		
		d.A |= int64(buf[6]) << 48
		
		d.A |= int64(buf[7]) << 56
		
	}
	{
		buf := make([]byte, 4)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		
		d.B |= uint32(buf[0]) << 0
		
		d.B |= uint32(buf[1]) << 8
		
		d.B |= uint32(buf[2]) << 16
		
		d.B |= uint32(buf[3]) << 24
		
	}
	{
		var v uint32
		
	{
		buf := make([]byte, 4)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return err
		}
		
		v |= uint32(buf[0]) << 0
		
		v |= uint32(buf[1]) << 8
		
		v |= uint32(buf[2]) << 16
		
		v |= uint32(buf[3]) << 24
		
	}
		d.C = math.Float32frombits(v)
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
		d.D = math.Float64frombits(v)
	}
	return nil
}

