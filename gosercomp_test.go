package gosercomp

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	vitessbson "github.com/youtube/vitess/go/bson"
	"github.com/golang/protobuf/proto"
	goproto "github.com/gogo/protobuf/proto"
	flatbuffers "github.com/google/flatbuffers/go"
)

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var protobufGroup = ProtoColorGroup{
	Id:     proto.Int32(17),
	Name:   proto.String("Reds"),
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var gogoProtobufGroup = GogoProtoColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}


func BenchmarkMarshalByJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(group)
	}
}
func BenchmarkUnmarshalByJson(b *testing.B) {
	bytes, _ := json.Marshal(group)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByXml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xml.Marshal(group)
	}
}
func BenchmarkUnmarshalByXml(b *testing.B) {
	bytes, _ := xml.Marshal(group)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xml.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByBson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vitessbson.Marshal(group)
	}
}
func BenchmarkUnmarshalByBson(b *testing.B) {
	bytes, _ := vitessbson.Marshal(group)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vitessbson.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByMsgp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		group.MarshalMsg(nil)
	}
}
func BenchmarkUnmarshalByMsgp(b *testing.B) {
	bytes, _ := group.MarshalMsg(nil)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalMsg(bytes)
	}
}

func BenchmarkMarshalByProtoBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		proto.Marshal(&protobufGroup)
	}
}
func BenchmarkUnmarshalByProtoBuf(b *testing.B) {
	bytes, _ := proto.Marshal(&protobufGroup)
	result := ProtoColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByGogoProtoBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goproto.Marshal(&gogoProtobufGroup)
	}
}
func BenchmarkUnmarshalByGogoProtoBuf(b *testing.B) {
	bytes, _ := proto.Marshal(&gogoProtobufGroup)
	result := GogoProtoColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goproto.Unmarshal(bytes, &result)
	}
}

func serializeByFlatBuffers(builder *flatbuffers.Builder, cg *ColorGroup) []byte {
	builder.Reset()

	//prepare data
	name := builder.CreateString(cg.Name)
	//prepare colors array
	colorsLen := len(cg.Colors)
	FlatBufferColorGroupStartColorsVector(builder, colorsLen)
	offsets := make([]flatbuffers.UOffsetT, colorsLen)
	for i := colorsLen - 1; i >= 0; i-- {
		offsets[i] = builder.CreateString(cg.Colors[i])
	}
	for i := colorsLen - 1; i >= 0; i-- {
		builder.PrependUOffsetT(offsets[i])
	}
	offset := builder.EndVector(colorsLen)


	FlatBufferColorGroupStart(builder)
	FlatBufferColorGroupAddCgId(builder, int32(cg.Id))
	FlatBufferColorGroupAddName(builder, name)
	FlatBufferColorGroupAddColors(builder, offset)
	builder.Finish(FlatBufferColorGroupEnd(builder))
	return builder.Bytes[builder.Head():]
}

func BenchmarkMarshalByFlatBuffers(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serializeByFlatBuffers(builder, &group)
	}
}

func BenchmarkUnmarshalByFlatBuffers(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	bytes := serializeByFlatBuffers(builder, &group)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := GetRootAsFlatBufferColorGroup(bytes,0)
		result.CgId()
//		result.Name()
//		colorsLen := result.ColorsLength()
//		for j := 0; j < colorsLen; j++ {
//			result.Colors(j)
//		}
	}
}