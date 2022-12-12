package main

import (
	"reflect"

	model "github.com/smallnest/gosercomp/model"
	"github.com/ymz-ncnk/musgo"
)

func main() {
	musGo, err := musgo.New()
	if err != nil {
		panic(err)
	}
	var v model.ColorGroup
	{
		err = musGo.Generate(reflect.TypeOf(v), false)
		if err != nil {
			panic(err)
		}
	}
	{
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Unsafe = true
		conf.Suffix = "MUSUnsafe"
		conf.Filename = conf.T.Name() + "Unsafe.musgen.go"
		err = musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
}
