package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"

	thrift "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/Sereal/Sereal/Go/sereal"
	memdump "github.com/alexflint/go-memdump"
	"github.com/ethereum/go-ethereum/rlp"
	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	hprose "github.com/hprose/hprose-golang/io"
	jsoniter "github.com/json-iterator/go"
	"github.com/niubaoshu/gotiny"
	sjson "github.com/segmentio/encoding/json"
	model "github.com/smallnest/gosercomp/model"
	"github.com/tidwall/gjson"
	"github.com/ugorji/go/codec"
	msgpackv4 "github.com/vmihailenco/msgpack/v4"
)

func BenchmarkMarshalByJson(b *testing.B) {
	var bb []byte
	for i := 0; i < b.N; i++ {
		bb, _ = json.Marshal(group)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByJson(b *testing.B) {
	bytes, _ := json.Marshal(group)
	result := model.ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByXml(b *testing.B) {
	var bb []byte
	for i := 0; i < b.N; i++ {
		bb, _ = xml.Marshal(group)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByXml(b *testing.B) {
	bytes, _ := xml.Marshal(group)
	result := model.ColorGroup{}
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
	var bb []byte
	for i := 0; i < b.N; i++ {
		bb, _ = group.MarshalMsg(nil)
	}
	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByMsgp(b *testing.B) {
	bytes, _ := group.MarshalMsg(nil)
	result := model.ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalMsg(bytes)
	}
}

func BenchmarkMarshalByProtoBuf(b *testing.B) {
	var bb []byte
	for i := 0; i < b.N; i++ {
		bb, _ = proto.Marshal(&protobufGroup)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByProtoBuf(b *testing.B) {
	bytes, _ := proto.Marshal(&protobufGroup)
	result := model.ProtoColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByGogoProtoBuf(b *testing.B) {
	var bb []byte
	for i := 0; i < b.N; i++ {
		bb, _ = goproto.Marshal(&gogoProtobufGroup)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByGogoProtoBuf(b *testing.B) {
	bytes, _ := proto.Marshal(&gogoProtobufGroup)
	result := model.GogoProtoColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goproto.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByThrift(b *testing.B) {
	t := thrift.NewTSerializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault() //NewTCompactProtocolFactory() or NewTJSONProtocolFactory()
	t.Protocol = pf.GetProtocol(t.Transport)

	var bb []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = t.Write(context.Background(), &thriftColorGroup)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByThrift(b *testing.B) {
	t := thrift.NewTDeserializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault()
	t.Protocol = pf.GetProtocol(t.Transport)

	t0 := thrift.NewTSerializer()
	t0.Protocol = pf.GetProtocol(t0.Transport)
	s, _ := t0.Write(context.Background(), &thriftColorGroup)

	result := model.ThriftColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Read(&result, s)
	}
}

func BenchmarkMarshalByAvro(b *testing.B) {
	buf := new(bytes.Buffer)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		err := avroGroup.Serialize(buf)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ReportMetric(float64(len(buf.Bytes())), "marshaledBytes")
}

func BenchmarkUnmarshalByAvro(b *testing.B) {
	buf := new(bytes.Buffer)
	err := avroGroup.Serialize(buf)
	if err != nil {
		b.Fatal(err)
	}

	bb := buf.Bytes()
	r := bytes.NewReader(bb)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Reset(bb)
		_, err = model.DeserializeAvroColorGroup(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByGencode(b *testing.B) {
	var group = model.GencodeColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	buf := make([]byte, group.Size())
	var bb []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = group.Marshal(buf)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByGencode(b *testing.B) {
	var group = model.GencodeColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	buf, _ := group.Marshal(nil)

	var groupResult model.GencodeColorGroup

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
		enc.Encode(group)
	}

	b.ReportMetric(float64(len(buf.Bytes())), "marshaledBytes")
}
func BenchmarkUnmarshalByUgorjiCodecAndCbor(b *testing.B) {
	var buf bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf, &ch)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()

	var g model.ColorGroup
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

	b.ReportMetric(float64(len(buf.Bytes())), "marshaledBytes")
}
func BenchmarkUnmarshalByUgorjiCodecAndMsgp(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
	var g model.ColorGroup
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

	b.ReportMetric(float64(len(buf.Bytes())), "marshaledBytes")
}
func BenchmarkUnmarshalByUgorjiCodecAndBinc(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.BincHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
	var g model.ColorGroup

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

	b.ReportMetric(float64(len(buf.Bytes())), "marshaledBytes")
}
func BenchmarkUnmarshalByUgorjiCodecAndJson(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.JsonHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
	var g model.ColorGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.ResetBytes(objectBytes)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByEasyjson(b *testing.B) {
	var bb []byte
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if bb, err = egroup.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByEasyjson(b *testing.B) {
	data, err := egroup.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}

	var g model.EColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.UnmarshalJSON(data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByFfjson(b *testing.B) {
	var bb []byte
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if bb, err = fgroup.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByFfjson(b *testing.B) {
	data, err := fgroup.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}

	var g model.FColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.UnmarshalJSON(data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByJsoniter(b *testing.B) {
	var bb []byte
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if bb, err = jsoniter.Marshal(&group); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByJsoniter(b *testing.B) {
	data, err := jsoniter.Marshal(&group)
	if err != nil {
		b.Fatal(err)
	}

	var g model.ColorGroup
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

	var g model.ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := gjson.Unmarshal(data, &g); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalByGoMemdump(b *testing.B) {
	var buf bytes.Buffer
	//w := bufio.NewWriter(&buf)

	for i := 0; i < b.N; i++ {
		buf.Reset()
		memdump.Encode(&buf, &group)
	}

	b.ReportMetric(float64(len(buf.Bytes())), "marshaledBytes")
}
func BenchmarkUnmarshalByGoMemdump(b *testing.B) {
	result := &model.ColorGroup{}

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

	b.ReportMetric(float64(l), "marshaledBytes")
}

func BenchmarkUnmarshalByColfer(b *testing.B) {
	result := &model.ColferColorGroup{}
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

	b.ReportMetric(float64(zgroup.Msgsize()), "marshaledBytes")
}

func BenchmarkUnmarshalByZebrapack(b *testing.B) {
	bts, _ := zgroup.MarshalMsg(nil)
	v := &model.ZColorGroup{}
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
		bytes = gotiny.Marshal(&group)
	}

	b.ReportMetric(float64(len(bytes)), "marshaledBytes")
}

func BenchmarkUnmarshalByGotiny(b *testing.B) {
	bytes := gotiny.Marshal(&group)

	v := &model.ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gotiny.Unmarshal(bytes, v)
	}
}

type HproseSerializer struct {
	writer *hprose.Writer
	reader *hprose.Reader
}

func (s *HproseSerializer) Marshal(o *model.ColorGroup) []byte {
	writer := s.writer
	writer.WriteInt(int64(o.Id))
	writer.WriteString(o.Name)
	writer.WriteValue(reflect.ValueOf(o.Colors))
	return writer.Bytes()
}

func (s *HproseSerializer) Unmarshal(o *model.ColorGroup) error {
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

	var bb []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writer.Clear()
		bb = s.Marshal(&group)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByHprose(b *testing.B) {
	v := &model.ColorGroup{}

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

	var bb []byte

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = encoder.Marshal(&group)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalBySereal(b *testing.B) {
	encoder := sereal.NewEncoderV3()
	bytes, _ := encoder.Marshal(&group)

	decoder := sereal.NewDecoder()
	v := &model.ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Unmarshal(bytes, v)
	}
}

func BenchmarkMarshalByVmihMsgpackv4(b *testing.B) {
	var bb []byte

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = msgpackv4.Marshal(&group)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByVmihMsgpackv4(b *testing.B) {
	bytes, _ := msgpackv4.Marshal(&group)
	v := &model.ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpackv4.Unmarshal(bytes, v)
	}
}

func BenchmarkMarshalByRlp(b *testing.B) {
	var bb []byte

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = rlp.EncodeToBytes(&rlpgroup)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalByRlp(b *testing.B) {
	bytes, _ := rlp.EncodeToBytes(&rlpgroup)
	v := &model.RlpColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rlp.DecodeBytes(bytes, v)
	}
}

func BenchmarkMarshalBySegmentioJSON(b *testing.B) {
	var bb []byte
	for i := 0; i < b.N; i++ {
		bb, _ = sjson.Marshal(group)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}
func BenchmarkUnmarshalBySegmentioJSON(b *testing.B) {
	bytes, _ := json.Marshal(group)
	result := model.ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sjson.Unmarshal(bytes, &result)
	}
}
