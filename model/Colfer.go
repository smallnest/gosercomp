package model

// Code generated by colf(1); DO NOT EDIT.
// The compiler used schema file colorgroup.colf.

import (
	"encoding/binary"
	"fmt"
	"io"
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

type ColferColorGroup struct {
	Id int32

	Name string

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
		i += copy(buf[i:], o.Name)
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
			x = uint(len(a))
			for x >= 0x80 {
				buf[i] = byte(x | 0x80)
				x >>= 7
				i++
			}
			buf[i] = byte(x)
			i++
			i += copy(buf[i:], a)
		}
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is model.ColferMax.
func (o *ColferColorGroup) MarshalLen() (int, error) {
	l := 1

	if v := o.Id; v != 0 {
		x := uint32(v)
		if v < 0 {
			x = ^x + 1
		}
		for l += 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if x := len(o.Name); x != 0 {
		if x > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field model.ColferColorGroup.name exceeds %d bytes", ColferSizeMax))
		}
		for l += x + 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if x := len(o.Colors); x != 0 {
		if x > ColferListMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field model.ColferColorGroup.colors exceeds %d elements", ColferListMax))
		}
		for l += 2; x >= 0x80; l++ {
			x >>= 7
		}
		for _, a := range o.Colors {
			x = len(a)
			if x > ColferSizeMax {
				return 0, ColferMax(fmt.Sprintf("colfer: field model.ColferColorGroup.colors exceeds %d bytes", ColferSizeMax))
			}
			for l += x + 1; x >= 0x80; l++ {
				x >>= 7
			}
		}
		if l >= ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: struct model.ColferColorGroup size exceeds %d bytes", ColferSizeMax))
		}
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct model.ColferColorGroup exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is model.ColferMax.
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
// The error return options are io.EOF, model.ColferError and model.ColferMax.
func (o *ColferColorGroup) Unmarshal(data []byte) (int, error) {
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
		x := uint32(data[i])
		i++

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
		o.Id = int32(x)

		header = data[i]
		i++
	} else if header == 0|0x80 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint32(data[i])
		i++

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
		o.Id = int32(^x + 1)

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

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: model.ColferColorGroup.name size %d exceeds %d bytes", x, ColferSizeMax))
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

	if header == 2 {
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
			return 0, ColferMax(fmt.Sprintf("colfer: model.ColferColorGroup.colors length %d exceeds %d elements", x, ColferListMax))
		}
		a := make([]string, int(x))
		o.Colors = a

		for ai := range a {
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
				return 0, ColferMax(fmt.Sprintf("colfer: model.ColferColorGroup.colors element %d size %d exceeds %d bytes", ai, x, ColferSizeMax))
			}

			start := i
			i += int(x)
			if i >= len(data) {
				goto eof
			}
			a[ai] = string(data[start:i])
		}

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
		return 0, ColferMax(fmt.Sprintf("colfer: struct model.ColferColorGroup size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, model.ColferError, model.ColferTail and model.ColferMax.
func (o *ColferColorGroup) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}
