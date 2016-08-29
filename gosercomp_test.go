package gosercomp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"testing"

	thrift "git.apache.org/thrift.git/lib/go/thrift"
	memdump "github.com/alexflint/go-memdump"
	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/linkedin/goavro"
	"github.com/ugorji/go/codec"
	//vitessbson "github.com/youtube/vitess/go/bson"
)

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var thriftColorGroup = ThriftColorGroup{
	ID:     1,
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

var avroSchema = `{"namespace": "gosercomp",
"type": "record",
"name": "ColorGroup",
"fields": [
	 {"name": "id", "type": "int"},
	 {"name": "name",  "type": "string"},
	 {"name": "colors", "type": {"type": "array", "items": "string"}}
]
}`

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

// func BenchmarkMarshalByBson(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		vitessbson.Marshal(group)
// 	}
// }
// func BenchmarkUnmarshalByBson(b *testing.B) {
// 	bytes, _ := vitessbson.Marshal(group)
// 	result := ColorGroup{}
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		vitessbson.Unmarshal(bytes, &result)
// 	}
// }

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
	offsets := make([]flatbuffers.UOffsetT, colorsLen)
	for i := colorsLen - 1; i >= 0; i-- {
		offsets[i] = builder.CreateString(cg.Colors[i])
	}

	FlatBufferColorGroupStartColorsVector(builder, colorsLen)
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
		_ = GetRootAsFlatBufferColorGroup(bytes, 0)
	}
}

func BenchmarkUnmarshalByFlatBuffers_withFields(b *testing.B) {
	builder := flatbuffers.NewBuilder(0)
	bytes := serializeByFlatBuffers(builder, &group)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := GetRootAsFlatBufferColorGroup(bytes, 0)
		result.CgId()
		result.Name()
		colorsLen := result.ColorsLength()
		for j := 0; j < colorsLen; j++ {
			result.Colors(j)
		}
	}
}

func BenchmarkMarshalByThrift(b *testing.B) {
	t := thrift.NewTSerializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault() //NewTCompactProtocolFactory() or NewTJSONProtocolFactory()
	t.Protocol = pf.GetProtocol(t.Transport)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = t.Write(&thriftColorGroup)
	}
}
func BenchmarkUnmarshalByThrift(b *testing.B) {
	t := thrift.NewTDeserializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = pf.GetProtocol(t.Transport)

	t0 := thrift.NewTSerializer()
	t0.Protocol = pf.GetProtocol(t0.Transport)
	s, _ := t0.Write(&thriftColorGroup)

	result := ThriftColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Read(&result, s)
	}
}

func BenchmarkMarshalByAvro(b *testing.B) {
	someRecord, err := goavro.NewRecord(goavro.RecordSchema(avroSchema))
	someRecord.Set("id", int32(1))
	someRecord.Set("name", "Reds")
	colors := []string{"Crimson", "Red", "Ruby", "Maroon"}
	s := make([]interface{}, len(colors))
	for i, v := range colors {
		s[i] = v
	}
	someRecord.Set("colors", s)

	codec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = codec.Encode(buf, someRecord)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkUnmarshalByAvro(b *testing.B) {
	codec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		panic(err)
	}
	someRecord, err := goavro.NewRecord(goavro.RecordSchema(avroSchema))
	someRecord.Set("id", int32(1))
	someRecord.Set("name", "Reds")
	colors := []string{"Crimson", "Red", "Ruby", "Maroon"}
	s := make([]interface{}, len(colors))
	for i, v := range colors {
		s[i] = v
	}
	someRecord.Set("colors", s)

	buf := new(bytes.Buffer)
	err = codec.Encode(buf, someRecord)
	if err != nil {
		panic(err)
	}
	objectBytes := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = codec.Decode(bytes.NewReader(objectBytes))
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMarshalByGencode(b *testing.B) {
	var group = GencodeColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	buf := make([]byte, group.Size())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		group.Marshal(buf)
	}
}
func BenchmarkUnmarshalByGencode(b *testing.B) {
	var group = GencodeColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	buf, _ := group.Marshal(nil)

	var groupResult GencodeColorGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupResult.Unmarshal(buf)
	}
}

func BenchmarkMarshalByCodecAndCbor(b *testing.B) {
	var buf bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf, &ch)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = enc.Encode(group)
	}
}
func BenchmarkUnmarshalByCodecAndCbor(b *testing.B) {
	var buf bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf, &ch)
	_ = enc.Encode(group)

	var g ColorGroup
	dec := codec.NewDecoder(&buf, &ch)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByCodecAndMsgp(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = enc.Encode(group)
	}
}
func BenchmarkUnmarshalByCodecAndMsgp(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	var g ColorGroup
	dec := codec.NewDecoder(&buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByGoMemdump(b *testing.B) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	for i := 0; i < b.N; i++ {
		memdump.Encode(w, &group)
	}
}
func BenchmarkUnmarshalByGoMemdump(b *testing.B) {
	result := &ColorGroup{}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	r := bufio.NewReader(&buf)
	memdump.Encode(w, &group)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memdump.Decode(r, &result)
	}
}
