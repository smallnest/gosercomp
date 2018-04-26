package gosercomp

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"reflect"
	"testing"

	thrift "git.apache.org/thrift.git/lib/go/thrift"
	memdump "github.com/alexflint/go-memdump"
	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	hprose "github.com/hprose/hprose-golang/io"
	jsoniter "github.com/json-iterator/go"
	"github.com/linkedin/goavro"
	"github.com/tidwall/gjson"
	"github.com/ugorji/go/codec"
	//vitessbson "github.com/youtube/vitess/go/bson"

	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/niubaoshu/gotiny"
	msgpackv2 "gopkg.in/vmihailenco/msgpack.v2"
)

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var zgroup = ZColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var egroup = EColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var fgroup = FColorGroup{
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

var colferGroup = ColferColorGroup{
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

func TestMarshaledDataLen(t *testing.T) {
	log.SetFlags(log.LstdFlags)

	buf, _ := json.Marshal(group)
	t.Logf("json:\t\t\t\t %d bytes", len(buf))

	buf, _ = xml.Marshal(group)
	t.Logf("xml:\t\t\t\t %d bytes", len(buf))

	buf, _ = group.MarshalMsg(nil)
	t.Logf("msgp:\t\t\t\t %d bytes", len(buf))

	buf, _ = proto.Marshal(&protobufGroup)
	t.Logf("protobuf:\t\t\t\t %d bytes", len(buf))

	buf, _ = goproto.Marshal(&gogoProtobufGroup)
	t.Logf("gogoprotobuf:\t\t\t %d bytes", len(buf))

	builder := flatbuffers.NewBuilder(0)
	buf = serializeByFlatBuffers(builder, &group)
	t.Logf("flatbuffers:\t\t\t %d bytes", len(buf))

	ts := thrift.NewTSerializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault() //NewTCompactProtocolFactory() or NewTJSONProtocolFactory()
	ts.Protocol = pf.GetProtocol(ts.Transport)
	buf, _ = ts.Write(context.Background(), &thriftColorGroup)
	t.Logf("thrift:\t\t\t\t %d bytes", len(buf))

	someRecord, _ := goavro.NewRecord(goavro.RecordSchema(avroSchema))
	someRecord.Set("id", int32(1))
	someRecord.Set("name", "Reds")
	colors := []string{"Crimson", "Red", "Ruby", "Maroon"}
	s := make([]interface{}, len(colors))
	for i, v := range colors {
		s[i] = v
	}
	someRecord.Set("colors", s)
	acodec, _ := goavro.NewCodec(avroSchema)
	buff := new(bytes.Buffer)
	_ = acodec.Encode(buff, someRecord)
	t.Logf("avro:\t\t\t\t %d bytes", len(buff.Bytes()))

	var group1 = GencodeColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	buff2 := make([]byte, group1.Size())
	b, _ := group1.Marshal(buff2)
	t.Logf("gencode:\t\t\t\t %d bytes", len(b))

	var buf3 bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf3, &ch)
	_ = enc.Encode(group)
	t.Logf("UgorjiCodec_Cbor:\t\t %d bytes", len(buf3.Bytes()))

	var buf4 bytes.Buffer
	var mh codec.MsgpackHandle
	enc4 := codec.NewEncoder(&buf4, &mh)
	enc4.Encode(group)
	t.Logf("UgorjiCodec_Msgp:\t\t %d bytes", len(buf4.Bytes()))

	var buf5 bytes.Buffer
	var mh5 codec.BincHandle
	enc5 := codec.NewEncoder(&buf5, &mh5)
	_ = enc5.Encode(group)
	t.Logf("UgorjiCodec_Bin:\t\t\t %d bytes", len(buf5.Bytes()))
	_ = enc5.Encode(group)
	t.Logf("UgorjiCodec_Json:\t\t %d bytes", len(buf5.Bytes()))

	buf, _ = egroup.MarshalJSON()
	t.Logf("easyjson:\t\t\t %d bytes", len(buf))

	buf, _ = fgroup.MarshalJSON()
	t.Logf("ffjson:\t\t\t\t %d bytes", len(buf))

	buf, _ = jsoniter.Marshal(&group)
	t.Logf("jsoniter:\t\t\t %d bytes", len(buf))

	var buf6 bytes.Buffer
	memdump.Encode(&buf6, &group)
	t.Logf("memdump:\t\t\t\t %d bytes", len(buf6.Bytes()))

	l, _ := colferGroup.MarshalLen()
	t.Logf("colfer:\t\t\t\t %d bytes", l)

	buf, _ = zgroup.MarshalMsg(buf[:0])
	t.Logf("zebrapack:\t\t\t %d bytes", len(buf))

	buf = gotiny.Encodes(&group)
	t.Logf("gotiny:\t\t\t\t %d bytes", len(buf))

	writer := hprose.NewWriter(true)
	ss := &HproseSerializer{writer: writer}
	_ = ss.Marshal(&group)
	t.Logf("hprose:\t\t\t\t %d bytes", len(buf))

	encoder := sereal.NewEncoderV3()
	buf, _ = encoder.Marshal(&group)
	t.Logf("sereal:\t\t\t\t %d bytes", len(buf))

	buf, _ = msgpackv2.Marshal(&group)
	t.Logf("msgpackv2:\t\t\t %d bytes", len(buf))
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
		_, _ = t.Write(context.Background(), &thriftColorGroup)
	}
}
func BenchmarkUnmarshalByThrift(b *testing.B) {
	t := thrift.NewTDeserializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = pf.GetProtocol(t.Transport)

	t0 := thrift.NewTSerializer()
	t0.Protocol = pf.GetProtocol(t0.Transport)
	s, _ := t0.Write(context.Background(), &thriftColorGroup)

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
		buf.Reset()
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

func BenchmarkMarshalByUgorjiCodecAndCbor(b *testing.B) {
	var buf bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf, &ch)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = enc.Encode(group)
	}
}
func BenchmarkUnmarshalByUgorjiCodecAndCbor(b *testing.B) {
	var buf bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf, &ch)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()

	var g ColorGroup
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &ch)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(objectBytes)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByUgorjiCodecAndMsgp(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = enc.Encode(group)
	}
}
func BenchmarkUnmarshalByUgorjiCodecAndMsgp(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
	var g ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(objectBytes)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByUgorjiCodecAndBinc(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.BincHandle
	enc := codec.NewEncoder(&buf, &mh)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = enc.Encode(group)
	}
}
func BenchmarkUnmarshalByUgorjiCodecAndBinc(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.BincHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
	var g ColorGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(objectBytes)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByUgorjiCodecAndJson(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.JsonHandle
	enc := codec.NewEncoder(&buf, &mh)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = enc.Encode(group)
	}
}
func BenchmarkUnmarshalByUgorjiCodecAndJson(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.JsonHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
	var g ColorGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(objectBytes)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByEasyjson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := egroup.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkUnmarshalByEasyjson(b *testing.B) {
	data, err := egroup.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}

	var g EColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.UnmarshalJSON(data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByFfjson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := fgroup.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkUnmarshalByFfjson(b *testing.B) {
	data, err := fgroup.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}

	var g FColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.UnmarshalJSON(data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByJsoniter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := jsoniter.Marshal(&group); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkUnmarshalByJsoniter(b *testing.B) {
	data, err := jsoniter.Marshal(&group)
	if err != nil {
		b.Fatal(err)
	}

	var g ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := jsoniter.Unmarshal(data, &g); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalByGJSON(b *testing.B) {
	data, err := json.Marshal(group)
	if err != nil {
		b.Fatal(err)
	}

	var g ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := gjson.Unmarshal(data, &g); err != nil {
			b.Fatal(err)
		}
	}
}

// func BenchmarkMarshalByIntelFastjson(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		if _, err := fastjson.Marshal(&group); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }
// func BenchmarkUnmarshalByIntelFastjson(b *testing.B) {
// 	data, err := fastjson.Marshal(&group)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	var g ColorGroup
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		if err := fastjson.Unmarshal(data, &g); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

func BenchmarkMarshalByGoMemdump(b *testing.B) {
	var buf bytes.Buffer
	//w := bufio.NewWriter(&buf)

	for i := 0; i < b.N; i++ {
		buf.Reset()
		memdump.Encode(&buf, &group)
	}
}
func BenchmarkUnmarshalByGoMemdump(b *testing.B) {
	result := &ColorGroup{}

	var buf bytes.Buffer
	memdump.Encode(&buf, &group)

	objectBytes := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memdump.Decode(bytes.NewReader(objectBytes), &result)
	}
}

func BenchmarkMarshalByColfer(b *testing.B) {
	l, _ := colferGroup.MarshalLen()
	buf := make([]byte, l)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		colferGroup.MarshalTo(buf)
	}
}

func BenchmarkUnmarshalByColfer(b *testing.B) {
	result := &ColferColorGroup{}
	buf, _ := colferGroup.MarshalBinary()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalBinary(buf)
	}
}

func BenchmarkMarshalByZebrapack(b *testing.B) {
	var bytes []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ = zgroup.MarshalMsg(bytes)
	}
}

func BenchmarkUnmarshalByZebrapack(b *testing.B) {
	bts, _ := zgroup.MarshalMsg(nil)
	v := &ZColorGroup{}
	//b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByGotiny(b *testing.B) {
	var bytes []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes = gotiny.Encodes(&group)
	}

	_ = bytes
}

func BenchmarkUnmarshalByGotiny(b *testing.B) {
	bytes := gotiny.Encodes(&group)
	v := &ColorGroup{}
	//b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gotiny.Decodes(bytes, v)
	}
}

type HproseSerializer struct {
	writer *hprose.Writer
	reader *hprose.Reader
}

func (s *HproseSerializer) Marshal(o *ColorGroup) []byte {
	writer := s.writer
	writer.WriteInt(int64(o.Id))
	writer.WriteString(o.Name)
	writer.WriteValue(reflect.ValueOf(o.Colors))
	return writer.Bytes()
}

func (s *HproseSerializer) Unmarshal(o *ColorGroup) error {
	reader := s.reader
	id := reader.ReadInt()
	o.Id = int(id)
	o.Name = reader.ReadString()
	var colors []string
	reader.ReadValue(reflect.ValueOf(&colors))
	o.Colors = colors
	return nil
}

func BenchmarkMarshalByHprose(b *testing.B) {
	writer := hprose.NewWriter(true)
	s := &HproseSerializer{writer: writer}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Marshal(&group)
	}
}

func BenchmarkUnmarshalByHprose(b *testing.B) {
	v := &ColorGroup{}

	writer := hprose.NewWriter(true)

	s := &HproseSerializer{writer: writer}
	bs := s.Marshal(&group)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := hprose.NewReader(bs, true)
		s.reader = reader
		s.Unmarshal(v)
	}
}

func BenchmarkMarshalBySereal(b *testing.B) {
	encoder := sereal.NewEncoderV3()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Marshal(&group)
	}
}

func BenchmarkUnmarshalBySereal(b *testing.B) {
	encoder := sereal.NewEncoderV3()
	bytes, _ := encoder.Marshal(&group)

	decoder := sereal.NewDecoder()
	v := &ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Unmarshal(bytes, v)
	}
}

func BenchmarkMarshalByMsgpackV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpackv2.Marshal(&group)
	}
}
func BenchmarkUnmarshalByMsgpackv2(b *testing.B) {
	bytes, _ := msgpackv2.Marshal(&group)
	v := &ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpackv2.Unmarshal(bytes, v)
	}
}
