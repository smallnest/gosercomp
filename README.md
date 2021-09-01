## Golang Serialization Benchmark

### Serializers

This project test the below go serializers, which compares with go standard _json_ and _xml_.

- [encoding/json](http://golang.org/pkg/encoding/json/)
- [encoding/xml](http://golang.org/pkg/encoding/xml/)
- [github.com/tinylib/msgp](http://github.com/tinylib/msgp)
- [github.com/golang/protobuf](http://github.com/golang/protobuf)
- [github.com/gogo/protobuf](http://github.com/gogo/protobuf)
- [Apache/Thrift](https://github.com/apache/thrift/tree/master/lib/go)
- [Apache/Avro](https://github.com/linkedin/goavro)
- [andyleap/gencode](https://github.com/andyleap/gencode)
- [ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [colfer](https://github.com/pascaldekloe/colfer)
- [zebrapack](https://github.com/glycerine/zebrapack)
- [gotiny](https://github.com/niubaoshu/gotiny)
- [github.com/ugorji/go/codec](http://github.com/ugorji/go/codec)
- [hprose-golang](https://github.com/hprose/hprose-golang/tree/master/io)
- [vmihailenco/msgpack/v4](https://github.com/vmihailenco/msgpack)
- [Sereal](https://github.com/Sereal/Sereal)
- [ffjson](https://github.com/pquerna/ffjson)
- [easyjson](https://github.com/mailru/easyjson)
- [jsoniter](https://github.com/json-iterator/go)
- [go-ethereum/rlp](https://github.com/ethereum/go-ethereum)
- [go-memdump](https://github.com/alexflint/go-memdump)

### Excluded Serializers

Given existed [benchmark](https://github.com/alecthomas/go_serialization_benchmarks) by alecthomas，or complexity， or activity, the below serializers are excluded from this test because of their poor performance.

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [bson](http://github.com/micro/go-bson)
- [github.com/google/flatbuffers](http://github.com/google/flatbuffers)

### Test Environment

go version: **1.13.4**

**Test:**

```
go test -bench=.
```

### Test Data Model

All tests are using the same data model as below:

```go
type ColorGroup struct {
    ID     int `json:"id" xml:"id,attr""`
    Name   string `json:"name" xml:"name"`
    Colors []string `json:"colors" xml:"colors"`
}
`
```

### Benchmark

![](benchmark.png)

**Marshal**

_include marshalled bytes_
The test machine is the MacBook Pro 16 with i7 2.6GHz and 32G 2667MHz DDR4 memory.
The test command we use is `go test -benchtime=5s -bench=. -benchmem`

```
BenchmarkMarshalByJson-12                               14336848               406 ns/op                65.0 marshaledBytes          128 B/op          2 allocs/op
BenchmarkMarshalByXml-12                                 2366107              2546 ns/op               137 marshaledBytes           4736 B/op         11 allocs/op
BenchmarkMarshalByMsgp-12                               75395308                77.2 ns/op              47.0 marshaledBytes           80 B/op          1 allocs/op
BenchmarkMarshalByProtoBuf-12                           18706688               320 ns/op                36.0 marshaledBytes           64 B/op          2 allocs/op
BenchmarkMarshalByGogoProtoBuf-12                       62550351                95.2 ns/op              36.0 marshaledBytes           48 B/op          1 allocs/op
BenchmarkMarshalByThrift-12                             20778853               289 ns/op                63.0 marshaledBytes           64 B/op          1 allocs/op
BenchmarkMarshalByThriftIterator-12                     17397620               343 ns/op                63.0 marshaledBytes          248 B/op          6 allocs/op
BenchmarkMarshalByThriftIteratorDynamic-12              16680176               360 ns/op                63.0 marshaledBytes          200 B/op          5 allocs/op
BenchmarkMarshalByThriftIteratorEncoder-12              26840368               253 ns/op                63.0 marshaledBytes          169 B/op          0 allocs/op
BenchmarkMarshalByAvro-12                               18352880               330 ns/op                32.0 marshaledBytes          112 B/op          2 allocs/op
BenchmarkMarshalByGencode-12                            196190748               30.7 ns/op              34.0 marshaledBytes            0 B/op          0 allocs/op
BenchmarkMarshalByUgorjiCodecAndCbor-12                  7795629               764 ns/op                47.0 marshaledBytes         1504 B/op          6 allocs/op
BenchmarkMarshalByUgorjiCodecAndMsgp-12                  7919284               754 ns/op                47.0 marshaledBytes         1504 B/op          6 allocs/op
BenchmarkMarshalByUgorjiCodecAndBinc-12                  7745150               773 ns/op                47.0 marshaledBytes         1504 B/op          6 allocs/op
BenchmarkMarshalByUgorjiCodecAndJson-12                  6579991               912 ns/op                65.0 marshaledBytes         1584 B/op          6 allocs/op
BenchmarkMarshalByEasyjson-12                           36914464               162 ns/op                65.0 marshaledBytes          128 B/op          1 allocs/op
BenchmarkMarshalByFfjson-12                              7739889               739 ns/op                65.0 marshaledBytes          424 B/op          9 allocs/op
BenchmarkMarshalByJsoniter-12                           19453400               306 ns/op                65.0 marshaledBytes           88 B/op          2 allocs/op
BenchmarkMarshalByGojay-12                              19624680               304 ns/op                65.0 marshaledBytes          544 B/op          2 allocs/op
BenchmarkMarshalByGoMemdump-12                           1715050              3428 ns/op               200 marshaledBytes           1368 B/op         26 allocs/op
BenchmarkMarshalByColfer-12                             241465138               24.2 ns/op              35.0 marshaledBytes            0 B/op          0 allocs/op
BenchmarkMarshalByZebrapack-12                          70622082               106 ns/op               109 marshaledBytes            121 B/op          0 allocs/op
BenchmarkMarshalByGotiny-12                             22547247               257 ns/op                32.0 marshaledBytes          144 B/op          5 allocs/op
BenchmarkMarshalByHprose-12                             24208729               246 ns/op                49.0 marshaledBytes           32 B/op          1 allocs/op
BenchmarkMarshalBySereal-12                              3786567              1520 ns/op                76.0 marshaledBytes          792 B/op         22 allocs/op
BenchmarkMarshalByVmihMsgpackv4-12                      12035164               495 ns/op                55.0 marshaledBytes          240 B/op          5 allocs/op
BenchmarkMarshalByRlp-12                                22251540               272 ns/op                32.0 marshaledBytes           64 B/op          3 allocs/op
BenchmarkMarshalBySegmentioJSON-12                      15845877               382 ns/op                65.0 marshaledBytes         1072 B/op          2 allocs/op
```

**Unmarshal**

```
BenchmarkUnmarshalByJson-12                              4311666              1396 ns/op             248 B/op          9 allocs/op
BenchmarkUnmarshalByXml-12                                719829              8465 ns/op            3038 B/op         74 allocs/op
BenchmarkUnmarshalByMsgp-12                             45195660               131 ns/op              32 B/op          5 allocs/op
BenchmarkUnmarshalByProtoBuf-12                         10638994               557 ns/op             176 B/op         11 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-12                     19217240               311 ns/op             160 B/op         10 allocs/op
BenchmarkUnmarshalByThrift-12                            7694175               807 ns/op             416 B/op         11 allocs/op
BenchmarkUnmarshalByThriftIterator-12                   16212927               371 ns/op             168 B/op          7 allocs/op
BenchmarkUnmarshalByThriftIteratorDynamic-12             6664594               904 ns/op             600 B/op         18 allocs/op
BenchmarkUnmarshalByThriftIteratorDecoder-12             6360967               939 ns/op             632 B/op         19 allocs/op
BenchmarkUnmarshalByAvro-12                               379012             15341 ns/op           12439 B/op        234 allocs/op
BenchmarkUnmarshalByGencode-12                          67695685                85.3 ns/op            32 B/op          5 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-12                5988073              1002 ns/op             656 B/op          8 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-12                5408978              1110 ns/op             768 B/op         10 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-12                6010320               995 ns/op             656 B/op          8 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndJson-12                3379785              1782 ns/op            1168 B/op         10 allocs/op
BenchmarkUnmarshalByEasyjson-12                         15563548               387 ns/op              32 B/op          5 allocs/op
BenchmarkUnmarshalByFfjson-12                            5721711              1065 ns/op             480 B/op         13 allocs/op
BenchmarkUnmarshalByJsoniter-12                         17867283               335 ns/op              32 B/op          5 allocs/op
BenchmarkUnmarshalByGJSON-12                             4815091              1210 ns/op             624 B/op          7 allocs/op
BenchmarkUnmarshalByGojay-12                            11270038               538 ns/op             288 B/op          9 allocs/op
BenchmarkUnmarshalByGoMemdump-12                         7590969               810 ns/op            2280 B/op         10 allocs/op
BenchmarkUnmarshalByColfer-12                           40564728               142 ns/op              96 B/op          6 allocs/op
BenchmarkUnmarshalByZebrapack-12                        28911075               205 ns/op              32 B/op          5 allocs/op
BenchmarkUnmarshalByGotiny-12                           26085504               234 ns/op             120 B/op          7 allocs/op
BenchmarkUnmarshalByHprose-12                           12682234               461 ns/op             288 B/op          9 allocs/op
BenchmarkUnmarshalBySereal-12                           11895444               502 ns/op              80 B/op          6 allocs/op
BenchmarkUnmarshalByVmihMsgpackv4-12                     7610895               808 ns/op             264 B/op         11 allocs/op
BenchmarkUnmarshalByRlp-12                               8878078               650 ns/op             104 B/op         11 allocs/op
BenchmarkUnmarshalBySegmentioJSON-12                    14947348               409 ns/op              32 B/op          5 allocs/op
```

**Marshaled Size**

![](size.png)
