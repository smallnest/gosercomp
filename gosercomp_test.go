package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/Sereal/Sereal/Go/sereal"
	memdump "github.com/alexflint/go-memdump"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/bytedance/sonic"

	// "github.com/bytedance/sonic"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/francoispqt/gojay"
	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	hprose "github.com/hprose/hprose-golang/io"
	jsoniter "github.com/json-iterator/go"
	sjson "github.com/segmentio/encoding/json"
	model "github.com/smallnest/gosercomp/model"
	thrift_iter "github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go/general"
	"github.com/ugorji/go/codec"
	msgpackv4 "github.com/vmihailenco/msgpack/v4"
)

func BenchmarkMarshalByJson(b *testing.B) {
	bb := make([]byte, 0, 1024)
	b.ResetTimer()
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
	bb := make([]byte, 0, 1024)
	b.ResetTimer()
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

func BenchmarkMarshalByMsgp(b *testing.B) {
	bb := make([]byte, 0, 1024)
	b.ResetTimer()

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
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
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

// func BenchmarkMarshalByProtoBuf_csproto(b *testing.B) {
// 	bb := make([]byte, 0, 1024)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		bb, _ = csproto.Marshal(&protobufGroup)
// 	}
// 	b.ReportMetric(float64(len(bb)), "marshaledBytes")
// }

// func BenchmarkUnmarshalByProtoBuf_csproto(b *testing.B) {
// 	bytes, _ := proto.Marshal(&protobufGroup)
// 	result := model.ProtoColorGroup{}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		csproto.Unmarshal(bytes, &result)
// 	}
// }

func BenchmarkMarshalByGogoProtoBuf(b *testing.B) {
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
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

// func BenchmarkMarshalByGogoProtoBuf_csproto(b *testing.B) {
// 	bb := make([]byte, 0, 1024)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		bb, _ = csproto.Marshal(&gogoProtobufGroup)
// 	}

// 	b.ReportMetric(float64(len(bb)), "marshaledBytes")
// }

// func BenchmarkUnmarshalByGogoProtoBuf_csproto(b *testing.B) {
// 	bytes, _ := proto.Marshal(&gogoProtobufGroup)
// 	result := model.GogoProtoColorGroup{}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		csproto.Unmarshal(bytes, &result)
// 	}
// }

func BenchmarkMarshalByThrift(b *testing.B) {
	t := thrift.NewTSerializer()
	pf := thrift.NewTBinaryProtocolFactoryDefault() // NewTCompactProtocolFactory() or NewTJSONProtocolFactory()
	t.Protocol = pf.GetProtocol(t.Transport)

	bb := make([]byte, 0, 1024)

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
		t.Read(context.Background(), &result, s)
	}
}

func BenchmarkMarshalByThriftIterator(b *testing.B) {
	bb := make([]byte, 0, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = thrift_iter.Marshal(thrfitIterGroup)
	}
	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByThriftIterator(b *testing.B) {
	bb, _ := thrift_iter.Marshal(thrfitIterGroup)
	var val model.ThriftIterColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		thrift_iter.Unmarshal(bb, &val)
	}
}

func BenchmarkMarshalByThriftIteratorDynamic(b *testing.B) {
	bb := make([]byte, 0, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = thrift_iter.Marshal(thrfitIterGroupDynamic)
	}
	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByThriftIteratorDynamic(b *testing.B) {
	bb, _ := thrift_iter.Marshal(thrfitIterGroup)
	var val general.Struct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = thrift_iter.Unmarshal(bb, &val)
	}
}

func BenchmarkMarshalByThriftIteratorEncoder(b *testing.B) {
	t := thrift.NewTSerializer()
	encoder := thrift_iter.NewEncoder(t.Transport)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = encoder.Encode(thrfitIterGroupDynamic)
	}
	b.ReportMetric(float64(t.Transport.Len()/b.N), "marshaledBytes")
}

func BenchmarkUnmarshalByThriftIteratorDecoder(b *testing.B) {
	bb, _ := thrift_iter.Marshal(thrfitIterGroup)
	var val general.Struct

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder := thrift_iter.NewDecoder(nil, bb)
		_ = decoder.Decode(&val)
	}
}

func BenchmarkMarshalByAvro(b *testing.B) {
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		err := avroGroup.Serialize(buf)
		if err != nil {
			b.Fatal(err)
		}
		bb = buf.Bytes()
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
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
	group := model.GencodeColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	bb := make([]byte, 0, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = group.Marshal(bb)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByGencode(b *testing.B) {
	group := model.GencodeColorGroup{
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
	var ch codec.CborHandle
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := codec.NewEncoder(&buf, &ch)
		enc.Encode(group)
		bb = buf.Bytes()
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByUgorjiCodecAndCbor(b *testing.B) {
	var buf bytes.Buffer
	var ch codec.CborHandle
	enc := codec.NewEncoder(&buf, &ch)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()

	var g model.ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec := codec.NewDecoder(bytes.NewReader(objectBytes), &ch)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByUgorjiCodecAndMsgp(b *testing.B) {
	var mh codec.MsgpackHandle
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := codec.NewEncoder(&buf, &mh)
		_ = enc.Encode(group)
		bb = buf.Bytes()
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByUgorjiCodecAndMsgp(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.MsgpackHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var g model.ColorGroup
		dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByUgorjiCodecAndBinc(b *testing.B) {
	var mh codec.BincHandle
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := codec.NewEncoder(&buf, &mh)
		_ = enc.Encode(group)
		bb = buf.Bytes()
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByUgorjiCodecAndBinc(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.BincHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	var g model.ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByUgorjiCodecAndJson(b *testing.B) {
	var mh codec.JsonHandle
	bb := make([]byte, 0, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := codec.NewEncoder(&buf, &mh)
		_ = enc.Encode(group)
		bb = buf.Bytes()
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByUgorjiCodecAndJson(b *testing.B) {
	var buf bytes.Buffer
	var mh codec.JsonHandle
	enc := codec.NewEncoder(&buf, &mh)
	_ = enc.Encode(group)

	objectBytes := buf.Bytes()
	var g model.ColorGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec := codec.NewDecoder(bytes.NewReader(objectBytes), &mh)
		_ = dec.Decode(&g)
	}
}

func BenchmarkMarshalByEasyjson(b *testing.B) {
	bb := make([]byte, 0, 1024)
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
	bb := make([]byte, 0, 1024)
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
	bb := make([]byte, 0, 1024)
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

func BenchmarkMarshalBySonic(b *testing.B) {
	bb := make([]byte, 0, 1024)
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if bb, err = sonic.Marshal(&group); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalBySonic(b *testing.B) {
	data, err := sonic.Marshal(&group)
	if err != nil {
		b.Fatal(err)
	}

	var g model.ColorGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := sonic.Unmarshal(data, &g); err != nil {
			b.Fatal(err)
		}
	}
}

// func BenchmarkUnmarshalByGJSON(b *testing.B) {
// 	data, err := json.Marshal(group)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	var g model.ColorGroup
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		if err := gjson.Unmarshal(data, &g); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

func BenchmarkMarshalByGojay(b *testing.B) {
	bb := make([]byte, 0, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bb, _ = gojay.Marshal(gojayGroup)
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
}

func BenchmarkUnmarshalByGojay(b *testing.B) {
	bytes, _ := json.Marshal(gojayGroup)
	result := model.GojayColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gojay.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByGoMemdump(b *testing.B) {
	var bb []byte

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.Reset()
		memdump.Encode(&buf, &group)
		bb = buf.Bytes()
	}

	b.ReportMetric(float64(len(bb)), "marshaledBytes")
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
	buf := make([]byte, 1024)

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
	bytes := make([]byte, 0, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ = zgroup.MarshalMsg(bytes)
	}

	b.ReportMetric(float64(zgroup.Msgsize()), "marshaledBytes")
}

func BenchmarkUnmarshalByZebrapack(b *testing.B) {
	bts, _ := zgroup.MarshalMsg(nil)
	v := &model.ZColorGroup{}
	// b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
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

	bb := make([]byte, 0, 1024)
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

	bb := make([]byte, 0, 1024)

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
	bb := make([]byte, 0, 1024)

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
	bb := make([]byte, 0, 1024)

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
	bb := make([]byte, 0, 1024)
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
