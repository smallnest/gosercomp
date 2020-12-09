package main

import (
	"github.com/golang/protobuf/proto"

	model "github.com/smallnest/gosercomp/model"
)

var group = model.ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var avroGroup = model.AvroColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var zgroup = model.ZColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var egroup = model.EColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var fgroup = model.FColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var thriftColorGroup = model.ThriftColorGroup{
	ID:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var protobufGroup = model.ProtoColorGroup{
	Id:     proto.Int32(17),
	Name:   proto.String("Reds"),
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var gogoProtobufGroup = model.GogoProtoColorGroup{
	Id:     proto.Int32(1),
	Name:   proto.String("Reds"),
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var colferGroup = model.ColferColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var avroSchema = `{"namespace": "gosercomp",
"type": "record",
"name": "ColorGroup",
"fields": [
	 {"name": "id", "type": "int"},
	 {"name": "name",  "type": "string"},
	 {"name": "colors", "type": {"type": "array", "items": "string"}}
]
}`

var rlpgroup = model.RlpColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}
