package model

type ThriftIterColorGroup struct {
	ID     int32    `thrift:",1" db:"id" json:"id"`
	Name   string   `thrift:",2" db:"name" json:"name"`
	Colors []string `thrift:",3" db:"colors" json:"colors"`
}
