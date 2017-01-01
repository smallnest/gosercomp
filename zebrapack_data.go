package gosercomp

//go:generate zebrapack -fast-strings -no-rtti
type ZColorGroup struct {
	Id     int      `json:"id" xml:"id,attr" msg:"id" zid:"0"`
	Name   string   `json:"name" xml:"name" msg:"name" zid:"1"`
	Colors []string `json:"colors" xml:"colors" msg:"colors" zid:"2"`
}
