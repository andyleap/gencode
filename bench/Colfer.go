package bench

// Code generated by colf(1); DO NOT EDIT.
// The compiler used schema file test.colf and fixed.colf.

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

var intconv = binary.BigEndian

// Colfer configuration attributes
var (
	// ColferSizeMax is the upper limit for serial byte sizes.
	ColferSizeMax = 16 * 1024 * 1024
	// ColferListMax is the upper limit for the number of elements in a list.
	ColferListMax = 64 * 1024
)

// ColferMax signals an upper limit breach.
type ColferMax string

// Error honors the error interface.
func (m ColferMax) Error() string { return string(m) }

// ColferError signals a data mismatch as as a byte index.
type ColferError int

// Error honors the error interface.
func (i ColferError) Error() string {
	return fmt.Sprintf("colfer: unknown header at byte %d", i)
}

// ColferTail signals data continuation as a byte index.
type ColferTail int

// Error honors the error interface.
func (i ColferTail) Error() string {
	return fmt.Sprintf("colfer: data continuation at byte %d", i)
}

type ColferFixed struct {
	A int64

	B uint32

	C float32

	D float64
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *ColferFixed) MarshalTo(buf []byte) int {
	var i int

	if v := o.A; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 0
		} else {
			x = ^x + 1
			buf[i] = 0 | 0x80
		}
		i++
		for n := 0; x >= 0x80 && n < 8; n++ {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if x := o.B; x >= 1<<21 {
		buf[i] = 1 | 0x80
		intconv.PutUint32(buf[i+1:], x)
		i += 5
	} else if x != 0 {
		buf[i] = 1
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if v := o.C; v != 0 {
		buf[i] = 2
		intconv.PutUint32(buf[i+1:], math.Float32bits(v))
		i += 5
	}

	if v := o.D; v != 0 {
		buf[i] = 3
		intconv.PutUint64(buf[i+1:], math.Float64bits(v))
		i += 9
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is bench.ColferMax.
func (o *ColferFixed) MarshalLen() (int, error) {
	l := 1

	if v := o.A; v != 0 {
		l += 2
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for n := 0; x >= 0x80 && n < 8; n++ {
			x >>= 7
			l++
		}
	}

	if x := o.B; x >= 1<<21 {
		l += 5
	} else if x != 0 {
		for l += 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if o.C != 0 {
		l += 5
	}

	if o.D != 0 {
		l += 9
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct bench.colferFixed exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is bench.ColferMax.
func (o *ColferFixed) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, bench.ColferError and bench.ColferMax.
func (o *ColferFixed) Unmarshal(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint64(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.A = int64(x)

		header = data[i]
		i++
	} else if header == 0|0x80 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint64(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.A = int64(^x + 1)

		header = data[i]
		i++
	}

	if header == 1 {
		start := i
		i++
		if i >= len(data) {
			goto eof
		}
		x := uint32(data[start])

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint32(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.B = x

		header = data[i]
		i++
	} else if header == 1|0x80 {
		start := i
		i += 4
		if i >= len(data) {
			goto eof
		}
		o.B = intconv.Uint32(data[start:])
		header = data[i]
		i++
	}

	if header == 2 {
		start := i
		i += 4
		if i >= len(data) {
			goto eof
		}
		o.C = math.Float32frombits(intconv.Uint32(data[start:]))
		header = data[i]
		i++
	}

	if header == 3 {
		start := i
		i += 8
		if i >= len(data) {
			goto eof
		}
		o.D = math.Float64frombits(intconv.Uint64(data[start:]))
		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	if i < ColferSizeMax {
		return i, nil
	}
eof:
	if i >= ColferSizeMax {
		return 0, ColferMax(fmt.Sprintf("colfer: struct bench.colferFixed size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, bench.ColferError, bench.ColferTail and bench.ColferMax.
func (o *ColferFixed) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}

type ColferPerson struct {
	Name string

	Age uint8

	Height float64
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *ColferPerson) MarshalTo(buf []byte) int {
	var i int

	if l := len(o.Name); l != 0 {
		buf[i] = 0
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.Name)
	}

	if x := o.Age; x != 0 {
		buf[i] = 1
		i++
		buf[i] = x
		i++
	}

	if v := o.Height; v != 0 {
		buf[i] = 2
		intconv.PutUint64(buf[i+1:], math.Float64bits(v))
		i += 9
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is bench.ColferMax.
func (o *ColferPerson) MarshalLen() (int, error) {
	l := 1

	if x := len(o.Name); x != 0 {
		if x > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field bench.colferPerson.name exceeds %d bytes", ColferSizeMax))
		}
		for l += x + 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if x := o.Age; x != 0 {
		l += 2
	}

	if o.Height != 0 {
		l += 9
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct bench.colferPerson exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is bench.ColferMax.
func (o *ColferPerson) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, bench.ColferError and bench.ColferMax.
func (o *ColferPerson) Unmarshal(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: bench.colferPerson.name size %d exceeds %d bytes", x, ColferSizeMax))
		}

		start := i
		i += int(x)
		if i >= len(data) {
			goto eof
		}
		o.Name = string(data[start:i])

		header = data[i]
		i++
	}

	if header == 1 {
		start := i
		i++
		if i >= len(data) {
			goto eof
		}
		o.Age = data[start]
		header = data[i]
		i++
	}

	if header == 2 {
		start := i
		i += 8
		if i >= len(data) {
			goto eof
		}
		o.Height = math.Float64frombits(intconv.Uint64(data[start:]))
		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	if i < ColferSizeMax {
		return i, nil
	}
eof:
	if i >= ColferSizeMax {
		return 0, ColferMax(fmt.Sprintf("colfer: struct bench.colferPerson size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, bench.ColferError, bench.ColferTail and bench.ColferMax.
func (o *ColferPerson) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}

type ColferGroup struct {
	Name string

	Members []*ColferPerson
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
// All nil entries in o.Members will be replaced with a new value.
func (o *ColferGroup) MarshalTo(buf []byte) int {
	var i int

	if l := len(o.Name); l != 0 {
		buf[i] = 0
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.Name)
	}

	if l := len(o.Members); l != 0 {
		buf[i] = 1
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for vi, v := range o.Members {
			if v == nil {
				v = new(ColferPerson)
				o.Members[vi] = v
			}
			i += v.MarshalTo(buf[i:])
		}
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is bench.ColferMax.
func (o *ColferGroup) MarshalLen() (int, error) {
	l := 1

	if x := len(o.Name); x != 0 {
		if x > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field bench.colferGroup.name exceeds %d bytes", ColferSizeMax))
		}
		for l += x + 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if x := len(o.Members); x != 0 {
		if x > ColferListMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field bench.colferGroup.members exceeds %d elements", ColferListMax))
		}
		for l += 2; x >= 0x80; l++ {
			x >>= 7
		}
		for _, v := range o.Members {
			if v == nil {
				l++
				continue
			}
			vl, err := v.MarshalLen()
			if err != nil {
				return 0, err
			}
			l += vl
		}
		if l > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: struct bench.colferGroup size exceeds %d bytes", ColferSizeMax))
		}
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct bench.colferGroup exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// All nil entries in o.Members will be replaced with a new value.
// The error return option is bench.ColferMax.
func (o *ColferGroup) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, bench.ColferError and bench.ColferMax.
func (o *ColferGroup) Unmarshal(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: bench.colferGroup.name size %d exceeds %d bytes", x, ColferSizeMax))
		}

		start := i
		i += int(x)
		if i >= len(data) {
			goto eof
		}
		o.Name = string(data[start:i])

		header = data[i]
		i++
	}

	if header == 1 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferListMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: bench.colferGroup.members length %d exceeds %d elements", x, ColferListMax))
		}

		l := int(x)
		a := make([]*ColferPerson, l)
		malloc := make([]ColferPerson, l)
		for ai := range a {
			v := &malloc[ai]
			a[ai] = v

			n, err := v.Unmarshal(data[i:])
			if err != nil {
				if err == io.EOF && len(data) >= ColferSizeMax {
					return 0, ColferMax(fmt.Sprintf("colfer: bench.colferGroup size exceeds %d bytes", ColferSizeMax))
				}
				return 0, err
			}
			i += n
		}
		o.Members = a

		if i >= len(data) {
			goto eof
		}
		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	if i < ColferSizeMax {
		return i, nil
	}
eof:
	if i >= ColferSizeMax {
		return 0, ColferMax(fmt.Sprintf("colfer: struct bench.colferGroup size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, bench.ColferError, bench.ColferTail and bench.ColferMax.
func (o *ColferGroup) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}