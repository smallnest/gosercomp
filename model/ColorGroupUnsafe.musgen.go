// Code generated by musgen. DO NOT EDIT.

package model

import (
	"reflect"
	"unsafe"

	"github.com/ymz-ncnk/musgo/errs"
)

// MarshalMUSUnsafe fills buf with the MUS encoding of v.
func (v ColorGroup) MarshalMUSUnsafe(buf []byte) int {
	i := 0
	{
		uv := uint64(v.Id)
		if v.Id < 0 {
			uv = ^(uv << 1)
		} else {
			uv = uv << 1
		}
		{
			for uv >= 0x80 {
				buf[i] = byte(uv) | 0x80
				uv >>= 7
				i++
			}
			buf[i] = byte(uv)
			i++
		}
	}
	{
		length := len(v.Name)
		{
			uv := uint64(length)
			if length < 0 {
				uv = ^(uv << 1)
			} else {
				uv = uv << 1
			}
			{
				for uv >= 0x80 {
					buf[i] = byte(uv) | 0x80
					uv >>= 7
					i++
				}
				buf[i] = byte(uv)
				i++
			}
		}
		if len(buf[i:]) < length {
			panic(errs.ErrSmallBuf)
		}
		i += copy(buf[i:], v.Name)
	}
	{
		length := len(v.Colors)
		{
			uv := uint64(length)
			if length < 0 {
				uv = ^(uv << 1)
			} else {
				uv = uv << 1
			}
			{
				for uv >= 0x80 {
					buf[i] = byte(uv) | 0x80
					uv >>= 7
					i++
				}
				buf[i] = byte(uv)
				i++
			}
		}
		for _, el := range v.Colors {
			{
				length := len(el)
				{
					uv := uint64(length)
					if length < 0 {
						uv = ^(uv << 1)
					} else {
						uv = uv << 1
					}
					{
						for uv >= 0x80 {
							buf[i] = byte(uv) | 0x80
							uv >>= 7
							i++
						}
						buf[i] = byte(uv)
						i++
					}
				}
				if len(buf[i:]) < length {
					panic(errs.ErrSmallBuf)
				}
				i += copy(buf[i:], el)
			}
		}
	}
	return i
}

// UnmarshalMUSUnsafe parses the MUS-encoded buf, and sets the result to *v.
func (v *ColorGroup) UnmarshalMUSUnsafe(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var uv uint64
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			}
			shift := 0
			done := false
			for l, b := range buf[i:] {
				if l == 9 && b > 1 {
					return i, errs.ErrOverflow
				}
				if b < 0x80 {
					uv = uv | uint64(b)<<shift
					done = true
					i += l + 1
					break
				}
				uv = uv | uint64(b&0x7F)<<shift
				shift += 7
			}
			if !done {
				return i, errs.ErrSmallBuf
			}
		}
		if uv&1 == 1 {
			uv = ^(uv >> 1)
		} else {
			uv = uv >> 1
		}
		v.Id = int(uv)
	}
	if err != nil {
		return i, errs.NewFieldError("Id", err)
	}
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				}
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l == 9 && b > 1 {
						return i, errs.ErrOverflow
					}
					if b < 0x80 {
						uv = uv | uint64(b)<<shift
						done = true
						i += l + 1
						break
					}
					uv = uv | uint64(b&0x7F)<<shift
					shift += 7
				}
				if !done {
					return i, errs.ErrSmallBuf
				}
			}
			if uv&1 == 1 {
				uv = ^(uv >> 1)
			} else {
				uv = uv >> 1
			}
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if len(buf) < i+length {
			return i, errs.ErrSmallBuf
		}
		content := buf[i : i+length]
		slcHeader := (*reflect.SliceHeader)(unsafe.Pointer(&content))
		strHeader := (*reflect.StringHeader)(unsafe.Pointer(&v.Name))
		strHeader.Data = slcHeader.Data
		strHeader.Len = slcHeader.Len
		i += length
	}
	if err != nil {
		return i, errs.NewFieldError("Name", err)
	}
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				}
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l == 9 && b > 1 {
						return i, errs.ErrOverflow
					}
					if b < 0x80 {
						uv = uv | uint64(b)<<shift
						done = true
						i += l + 1
						break
					}
					uv = uv | uint64(b&0x7F)<<shift
					shift += 7
				}
				if !done {
					return i, errs.ErrSmallBuf
				}
			}
			if uv&1 == 1 {
				uv = ^(uv >> 1)
			} else {
				uv = uv >> 1
			}
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		v.Colors = make([]string, length)
		for j := 0; j < length; j++ {
			{
				var length int
				{
					var uv uint64
					{
						if i > len(buf)-1 {
							return i, errs.ErrSmallBuf
						}
						shift := 0
						done := false
						for l, b := range buf[i:] {
							if l == 9 && b > 1 {
								return i, errs.ErrOverflow
							}
							if b < 0x80 {
								uv = uv | uint64(b)<<shift
								done = true
								i += l + 1
								break
							}
							uv = uv | uint64(b&0x7F)<<shift
							shift += 7
						}
						if !done {
							return i, errs.ErrSmallBuf
						}
					}
					if uv&1 == 1 {
						uv = ^(uv >> 1)
					} else {
						uv = uv >> 1
					}
					length = int(uv)
				}
				if length < 0 {
					return i, errs.ErrNegativeLength
				}
				if len(buf) < i+length {
					return i, errs.ErrSmallBuf
				}
				content := buf[i : i+length]
				slcHeader := (*reflect.SliceHeader)(unsafe.Pointer(&content))
				strHeader := (*reflect.StringHeader)(unsafe.Pointer(&v.Colors[j]))
				strHeader.Data = slcHeader.Data
				strHeader.Len = slcHeader.Len
				i += length
			}
			if err != nil {
				err = errs.NewSliceError(j, err)
				break
			}
		}
	}
	if err != nil {
		return i, errs.NewFieldError("Colors", err)
	}
	return i, err
}

// SizeMUSUnsafe returns the size of the MUS-encoded v.
func (v ColorGroup) SizeMUSUnsafe() int {
	size := 0
	{
		uv := uint64(v.Id<<1) ^ uint64(v.Id>>63)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}
	}
	{
		length := len(v.Name)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					uv >>= 7
					size++
				}
				size++
			}
		}
		size += len(v.Name)
	}
	{
		length := len(v.Colors)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					uv >>= 7
					size++
				}
				size++
			}
		}
		for _, el := range v.Colors {
			{
				length := len(el)
				{
					uv := uint64(length<<1) ^ uint64(length>>63)
					{
						for uv >= 0x80 {
							uv >>= 7
							size++
						}
						size++
					}
				}
				size += len(el)
			}
		}
	}
	return size
}