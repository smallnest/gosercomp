package gosercomp

// This file was generated by colf(1); DO NOT EDIT

import (
	"fmt"
	"io"
)

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

type ColferColorGroup struct {
	Id     int32
	Name   string
	Colors []string
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *ColferColorGroup) MarshalTo(buf []byte) int {
	var i int

	if v := o.Id; v != 0 {
		x := uint32(v)
		if v >= 0 {
			buf[i] = 0
		} else {
			x = ^x + 1
			buf[i] = 0 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if l := len(o.Name); l != 0 {
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
		copy(buf[i:], o.Name)
		i += l
	}

	if l := len(o.Colors); l != 0 {
		buf[i] = 2
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		for _, a := range o.Colors {
			l = len(a)
			x = uint(l)
			for x >= 0x80 {
				buf[i] = byte(x | 0x80)
				x >>= 7
				i++
			}
			buf[i] = byte(x)
			i++
			copy(buf[i:], a)
			i += l
		}
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is gosercomp.ColferMax.
func (o *ColferColorGroup) MarshalLen() (int, error) {
	l := 1

	if v := o.Id; v != 0 {
		l += 2
		x := uint32(v)
		if v < 0 {
			x = ^x + 1
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
	}

	if x := len(o.Name); x != 0 {
		l += x
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
	}

	if x := len(o.Colors); x != 0 {
		if x > ColferListMax {
			return -1, ColferMax(fmt.Sprintf("colfer: field gosercomp.ColferColorGroup.colors exceeds %d elements", ColferListMax))
		}
		for x >= 0x80 {
			x >>= 7
			l++
		}
		l += 2
		for _, a := range o.Colors {
			x = len(a)
			l += x
			for x >= 0x80 {
				x >>= 7
				l++
			}
			l++
		}
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct gosercomp.ColferColorGroup exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is gosercomp.ColferMax.
func (o *ColferColorGroup) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, gosercomp.ColferError and gosercomp.ColferMax.
func (o *ColferColorGroup) Unmarshal(data []byte) (int, error) {
	if len(data) > ColferSizeMax {
		n, err := o.Unmarshal(data[:ColferSizeMax])
		if err == io.EOF {
			return 0, ColferMax(fmt.Sprintf("colfer: struct gosercomp.ColferColorGroup exceeds %d bytes", ColferSizeMax))
		}
		return n, err
	}

	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		o.Id = int32(x)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	} else if header == 0|0x80 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		o.Id = int32(^x + 1)

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	}

	if header == 1 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		to := i + int(x)
		if to >= len(data) {
			return 0, io.EOF
		}
		o.Name = string(data[i:to])

		header = data[to]
		i = to + 1
	}

	if header == 2 {
		var x uint32
		for shift := uint(0); ; shift += 7 {
			if i >= len(data) {
				return 0, io.EOF
			}
			b := data[i]
			i++
			if b < 0x80 {
				x |= uint32(b) << shift
				break
			}
			x |= (uint32(b) & 0x7f) << shift
		}
		l := int(x)
		if l > ColferListMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field gosercomp.ColferColorGroup.colors length %d exceeds %d elements", l, ColferListMax))
		}
		a := make([]string, l)
		o.Colors = a
		for ai := range a {
			var x uint32
			for shift := uint(0); ; shift += 7 {
				if i >= len(data) {
					return 0, io.EOF
				}
				b := data[i]
				i++
				if b < 0x80 {
					x |= uint32(b) << shift
					break
				}
				x |= (uint32(b) & 0x7f) << shift
			}
			to := i + int(x)
			if to >= len(data) {
				return 0, io.EOF
			}
			a[ai] = string(data[i:to])
			i = to
		}

		if i >= len(data) {
			return 0, io.EOF
		}
		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	return i, nil
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, gosercomp.ColferError, gosercomp.ColferTail and gosercomp.ColferMax.
func (o *ColferColorGroup) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if err != nil {
		return err
	}
	if i != len(data) {
		return ColferTail(i)
	}
	return nil
}