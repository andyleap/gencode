package main
import (
	"math"
	"io"
	"reflect"
	"unsafe"
)

var (
	_ = io.ReadFull
	_ = math.Float64frombits
	_ = reflect.ValueOf
	_ = unsafe.Sizeof(0)
)

type Person struct {
	Name string
	Age uint8
	Height float64
}

func (d *Person) Serialize(w io.Writer) error {
	buf := []byte{0,0,0,0,0,0,0,0,0,0}
	{
		l := uint64(len(d.Name))
		
	{
		t := uint64(l)
		
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
	}
		var err error
		if sw, ok := w.(interface{WriteString(s string) (n int, err error);}); ok {
			_, err = sw.WriteString(d.Name)
		} else {
			_, err = w.Write([]byte(d.Name))
		}
		if err != nil {
			return err
		}
	}
	{
		
		buf[0] = byte(d.Age >> 0)
		
		
		_, err := w.Write(buf[:1])
		if err != nil {
			return err
		}
	}
	{
		v := math.Float64bits(d.Height)
		
	{
		
		buf[0] = byte(v >> 0)
		
		buf[1] = byte(v >> 8)
		
		buf[2] = byte(v >> 16)
		
		buf[3] = byte(v >> 24)
		
		buf[4] = byte(v >> 32)
		
		buf[5] = byte(v >> 40)
		
		buf[6] = byte(v >> 48)
		
		buf[7] = byte(v >> 56)
		
		
		_, err := w.Write(buf[:8])
		if err != nil {
			return err
		}
	}
	}
	return nil
}

func (d *Person) Deserialize(r io.Reader) error {
	buf := []byte{0,0,0,0,0,0,0,0,0,0}
	{
		l := uint64(0)
		
	{
		buf[0] = 0x80
		t := uint64(0)
		i := uint(0)
		for buf[0] & 0x80 == 0x80 {
			_, err := r.Read(buf[0:1])
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F) << i
			i += 7
		}
		
		l = t
		
	}
		sbuf := make([]byte, l)
		_, err := io.ReadFull(r, sbuf)
		if err != nil {
			return err
		}
		d.Name = *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: uintptr(unsafe.Pointer(&sbuf[0])), Len: int(l)}))
	}
	{
		_, err := io.ReadFull(r, buf[:1])
		if err != nil {
			return err
		}
		
		d.Age |= uint8(buf[0]) << 0
		
	}
	{
		var v uint64
		
	{
		_, err := io.ReadFull(r, buf[:8])
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
	Members []Person
}

func (d *Group) Serialize(w io.Writer) error {
	buf := []byte{0,0,0,0,0,0,0,0,0,0}
	{
		l := uint64(len(d.Name))
		
	{
		t := uint64(l)
		
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
	}
		var err error
		if sw, ok := w.(interface{WriteString(s string) (n int, err error);}); ok {
			_, err = sw.WriteString(d.Name)
		} else {
			_, err = w.Write([]byte(d.Name))
		}
		if err != nil {
			return err
		}
	}
	{
		l := uint64(len(d.Members))
		
	{
		t := uint64(l)
		
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
	}
		for k := range d.Members {
			
	{
		err := d.Members[k].Serialize(w)
		if err != nil {
			return err
		}
	}
		}
	}
	return nil
}

func (d *Group) Deserialize(r io.Reader) error {
	buf := []byte{0,0,0,0,0,0,0,0,0,0}
	{
		l := uint64(0)
		
	{
		buf[0] = 0x80
		t := uint64(0)
		i := uint(0)
		for buf[0] & 0x80 == 0x80 {
			_, err := r.Read(buf[0:1])
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F) << i
			i += 7
		}
		
		l = t
		
	}
		sbuf := make([]byte, l)
		_, err := io.ReadFull(r, sbuf)
		if err != nil {
			return err
		}
		d.Name = *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: uintptr(unsafe.Pointer(&sbuf[0])), Len: int(l)}))
	}
	{
		l := uint64(0)
		
	{
		buf[0] = 0x80
		t := uint64(0)
		i := uint(0)
		for buf[0] & 0x80 == 0x80 {
			_, err := r.Read(buf[0:1])
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F) << i
			i += 7
		}
		
		l = t
		
	}
		if uint64(cap(d.Members)) >= l {
			d.Members = d.Members[:l]
		} else {
			d.Members = make([]Person, l)
		}
		for k := range d.Members {
			
	{
		err := d.Members[k].Deserialize(r)
		if err != nil {
			return err
		}
	}
		}
	}
	return nil
}

type A struct {
	Name string
	BirthDay int64
	Phone string
	Siblings int64
	Spouse uint8
	Money float64
}

func (d *A) Serialize(w io.Writer) error {
	buf := []byte{0,0,0,0,0,0,0,0,0,0}
	{
		l := uint64(len(d.Name))
		
	{
		t := uint64(l)
		
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
	}
		var err error
		if sw, ok := w.(interface{WriteString(s string) (n int, err error);}); ok {
			_, err = sw.WriteString(d.Name)
		} else {
			_, err = w.Write([]byte(d.Name))
		}
		if err != nil {
			return err
		}
	}
	{
		
		buf[0] = byte(d.BirthDay >> 0)
		
		buf[1] = byte(d.BirthDay >> 8)
		
		buf[2] = byte(d.BirthDay >> 16)
		
		buf[3] = byte(d.BirthDay >> 24)
		
		buf[4] = byte(d.BirthDay >> 32)
		
		buf[5] = byte(d.BirthDay >> 40)
		
		buf[6] = byte(d.BirthDay >> 48)
		
		buf[7] = byte(d.BirthDay >> 56)
		
		
		_, err := w.Write(buf[:8])
		if err != nil {
			return err
		}
	}
	{
		l := uint64(len(d.Phone))
		
	{
		t := uint64(l)
		
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
	}
		var err error
		if sw, ok := w.(interface{WriteString(s string) (n int, err error);}); ok {
			_, err = sw.WriteString(d.Phone)
		} else {
			_, err = w.Write([]byte(d.Phone))
		}
		if err != nil {
			return err
		}
	}
	{
		
		buf[0] = byte(d.Siblings >> 0)
		
		buf[1] = byte(d.Siblings >> 8)
		
		buf[2] = byte(d.Siblings >> 16)
		
		buf[3] = byte(d.Siblings >> 24)
		
		buf[4] = byte(d.Siblings >> 32)
		
		buf[5] = byte(d.Siblings >> 40)
		
		buf[6] = byte(d.Siblings >> 48)
		
		buf[7] = byte(d.Siblings >> 56)
		
		
		_, err := w.Write(buf[:8])
		if err != nil {
			return err
		}
	}
	{
		
		buf[0] = byte(d.Spouse >> 0)
		
		
		_, err := w.Write(buf[:1])
		if err != nil {
			return err
		}
	}
	{
		v := math.Float64bits(d.Money)
		
	{
		
		buf[0] = byte(v >> 0)
		
		buf[1] = byte(v >> 8)
		
		buf[2] = byte(v >> 16)
		
		buf[3] = byte(v >> 24)
		
		buf[4] = byte(v >> 32)
		
		buf[5] = byte(v >> 40)
		
		buf[6] = byte(v >> 48)
		
		buf[7] = byte(v >> 56)
		
		
		_, err := w.Write(buf[:8])
		if err != nil {
			return err
		}
	}
	}
	return nil
}

func (d *A) Deserialize(r io.Reader) error {
	buf := []byte{0,0,0,0,0,0,0,0,0,0}
	{
		l := uint64(0)
		
	{
		buf[0] = 0x80
		t := uint64(0)
		i := uint(0)
		for buf[0] & 0x80 == 0x80 {
			_, err := r.Read(buf[0:1])
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F) << i
			i += 7
		}
		
		l = t
		
	}
		sbuf := make([]byte, l)
		_, err := io.ReadFull(r, sbuf)
		if err != nil {
			return err
		}
		d.Name = *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: uintptr(unsafe.Pointer(&sbuf[0])), Len: int(l)}))
	}
	{
		_, err := io.ReadFull(r, buf[:8])
		if err != nil {
			return err
		}
		
		d.BirthDay |= int64(buf[0]) << 0
		
		d.BirthDay |= int64(buf[1]) << 8
		
		d.BirthDay |= int64(buf[2]) << 16
		
		d.BirthDay |= int64(buf[3]) << 24
		
		d.BirthDay |= int64(buf[4]) << 32
		
		d.BirthDay |= int64(buf[5]) << 40
		
		d.BirthDay |= int64(buf[6]) << 48
		
		d.BirthDay |= int64(buf[7]) << 56
		
	}
	{
		l := uint64(0)
		
	{
		buf[0] = 0x80
		t := uint64(0)
		i := uint(0)
		for buf[0] & 0x80 == 0x80 {
			_, err := r.Read(buf[0:1])
			if err != nil {
				return err
			}
			t |= uint64(buf[0]&0x7F) << i
			i += 7
		}
		
		l = t
		
	}
		sbuf := make([]byte, l)
		_, err := io.ReadFull(r, sbuf)
		if err != nil {
			return err
		}
		d.Phone = *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: uintptr(unsafe.Pointer(&sbuf[0])), Len: int(l)}))
	}
	{
		_, err := io.ReadFull(r, buf[:8])
		if err != nil {
			return err
		}
		
		d.Siblings |= int64(buf[0]) << 0
		
		d.Siblings |= int64(buf[1]) << 8
		
		d.Siblings |= int64(buf[2]) << 16
		
		d.Siblings |= int64(buf[3]) << 24
		
		d.Siblings |= int64(buf[4]) << 32
		
		d.Siblings |= int64(buf[5]) << 40
		
		d.Siblings |= int64(buf[6]) << 48
		
		d.Siblings |= int64(buf[7]) << 56
		
	}
	{
		_, err := io.ReadFull(r, buf[:1])
		if err != nil {
			return err
		}
		
		d.Spouse |= uint8(buf[0]) << 0
		
	}
	{
		var v uint64
		
	{
		_, err := io.ReadFull(r, buf[:8])
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
		d.Money = math.Float64frombits(v)
	}
	return nil
}

