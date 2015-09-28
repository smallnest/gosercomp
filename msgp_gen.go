package gosercomp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *ColorGroup) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "id"
	o = append(o, 0x83, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt(o, z.Id)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "colors"
	o = append(o, 0xa6, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Colors)))
	for xvk := range z.Colors {
		o = msgp.AppendString(o, z.Colors[xvk])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ColorGroup) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.Id, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "colors":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Colors) >= int(xsz) {
				z.Colors = z.Colors[:xsz]
			} else {
				z.Colors = make([]string, xsz)
			}
			for xvk := range z.Colors {
				z.Colors[xvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *ColorGroup) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + msgp.ArrayHeaderSize
	for xvk := range z.Colors {
		s += msgp.StringPrefixSize + len(z.Colors[xvk])
	}
	return
}
