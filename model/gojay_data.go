package model

import (
	"sync"

	"github.com/francoispqt/gojay"
)

type GojayColorGroup struct {
	Id     int      `json:"id" xml:"id,attr" msg:"id"`
	Name   string   `json:"name" xml:"name" msg:"name"`
	Colors []string `json:"colors" xml:"colors" msg:"colors"`
}

func init() {
	GojayColorGroupPool = &sync.Pool{
		New: func() interface{} {
			return &GojayColorGroup{}
		},
	}
}

var GojayColorGroupPool *sync.Pool

type Strings []string

// UnmarshalJSONArray decodes JSON array elements into slice
func (a *Strings) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value string
	if err := dec.String(&value); err != nil {
		return err
	}
	*a = append(*a, value)
	return nil
}

// MarshalJSONArray encodes arrays into JSON
func (a Strings) MarshalJSONArray(enc *gojay.Encoder) {
	for _, item := range a {
		enc.String(item)
	}
}

// IsNil checks if array is nil
func (a Strings) IsNil() bool {
	return len(a) == 0
}

// MarshalJSONObject implements MarshalerJSONObject
func (g *GojayColorGroup) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", g.Id)
	enc.StringKey("name", g.Name)
	colorsSlice := Strings(g.Colors)
	enc.ArrayKey("colors", colorsSlice)
}

// IsNil checks if instance is nil
func (g *GojayColorGroup) IsNil() bool {
	return g == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (g *GojayColorGroup) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "id":
		return dec.Int(&g.Id)

	case "name":
		return dec.String(&g.Name)

	case "colors":
		aSlice := Strings{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			g.Colors = []string(aSlice)
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (g *GojayColorGroup) NKeys() int { return 3 }

// Reset reset fields
func (g *GojayColorGroup) Reset() {
	g.Id = 0
	g.Name = ""
	g.Colors = nil
}
