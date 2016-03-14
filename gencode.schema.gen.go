package gosercomp

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type GencodeColorGroup struct {
	Id     int32
	Name   string
	Colors []string
}

func (d *GencodeColorGroup) Size() (s uint64) {

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
		l := uint64(len(d.Colors))

		{

			t := l
			for t >= 0x80 {
				t <<= 7
				s++
			}
			s++

		}
		for k := range d.Colors {

			{
				l := uint64(len(d.Colors[k]))

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

		}
	}
	s += 4
	return
}
func (d *GencodeColorGroup) Marshal(buf []byte) ([]byte, error) {
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

		buf[i+0+0] = byte(d.Id >> 0)

		buf[i+1+0] = byte(d.Id >> 8)

		buf[i+2+0] = byte(d.Id >> 16)

		buf[i+3+0] = byte(d.Id >> 24)

	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+4] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+4] = byte(t)
			i++

		}
		copy(buf[i+4:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Colors))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+4] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+4] = byte(t)
			i++

		}
		for k := range d.Colors {

			{
				l := uint64(len(d.Colors[k]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+4] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+4] = byte(t)
					i++

				}
				copy(buf[i+4:], d.Colors[k])
				i += l
			}

		}
	}
	return buf[:i+4], nil
}

func (d *GencodeColorGroup) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Id = 0 | (int32(buf[i+0+0]) << 0) | (int32(buf[i+1+0]) << 8) | (int32(buf[i+2+0]) << 16) | (int32(buf[i+3+0]) << 24)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+4] & 0x7F)
			for buf[i+4]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+4]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+4 : i+4+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+4] & 0x7F)
			for buf[i+4]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+4]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Colors)) >= l {
			d.Colors = d.Colors[:l]
		} else {
			d.Colors = make([]string, l)
		}
		for k := range d.Colors {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+4] & 0x7F)
					for buf[i+4]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+4]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.Colors[k] = string(buf[i+4 : i+4+l])
				i += l
			}

		}
	}
	return i + 4, nil
}
