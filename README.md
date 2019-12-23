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

###  Excluded Serializers

Given existed [benchmark](https://github.com/alecthomas/go_serialization_benchmarks) by alecthomas，or complexity， or activity, the below serializers are excluded from this test because of their poor performance.

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [bson](http://github.com/micro/go-bson)
- [go-memdump](https://github.com/alexflint/go-memdump)
- [github.com/google/flatbuffers](http://github.com/google/flatbuffers)
- [go-memdump](https://github.com/alexflint/go-memdump)

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

**Marshal**

_include marshalled bytes_

```
BenchmarkMarshalByJson-6                         2543781               480 ns/op                65.0 marshaledBytes
BenchmarkMarshalByXml-6                           401060              2842 ns/op               137 marshaledBytes
BenchmarkMarshalByMsgp-6                        15432038                80.7 ns/op              47.0 marshaledBytes
BenchmarkMarshalByProtoBuf-6                     8019920               150 ns/op                36.0 marshaledBytes
BenchmarkMarshalByGogoProtoBuf-6                12277683                94.9 ns/op              36.0 marshaledBytes
BenchmarkMarshalByThrift-6                       3795267               317 ns/op                63.0 marshaledBytes
BenchmarkMarshalByAvro-6                         3668318               331 ns/op                32.0 marshaledBytes
BenchmarkMarshalByGencode-6                     37568914                31.3 ns/op              34.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndCbor-6           1447856               849 ns/op                47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndMsgp-6           1465488               806 ns/op                47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndBinc-6           1469070               823 ns/op                47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndJson-6           1000000              1037 ns/op                65.0 marshaledBytes
BenchmarkMarshalByEasyjson-6                     6106801               196 ns/op                65.0 marshaledBytes
BenchmarkMarshalByFfjson-6                       1666495               728 ns/op                65.0 marshaledBytes
BenchmarkMarshalByJsoniter-6                     3208327               375 ns/op                65.0 marshaledBytes
BenchmarkMarshalByGoMemdump-6                     344076              3496 ns/op               200 marshaledBytes
BenchmarkMarshalByColfer-6                      50134317                23.9 ns/op              35.0 marshaledBytes
BenchmarkMarshalByZebrapack-6                    9549326               116 ns/op               109 marshaledBytes
BenchmarkMarshalByGotiny-6                       4663609               258 ns/op                32.0 marshaledBytes
BenchmarkMarshalByHprose-6                       4645060               259 ns/op                49.0 marshaledBytes
BenchmarkMarshalBySereal-6                        750622              1590 ns/op                76.0 marshaledBytes
BenchmarkMarshalByVmihMsgpackv4-6                2378013               500 ns/op                55.0 marshaledBytes
BenchmarkMarshalByRlp-6                          3351589               359 ns/op                32.0 marshaledBytes
BenchmarkMarshalBySegmentioJSON-6                3249405               369 ns/op                65.0 marshaledBytes
```


**Unmarshal**

```
BenchmarkUnmarshalByJson-6                        752025              1597 ns/op
BenchmarkUnmarshalByXml-6                         128037              9597 ns/op
BenchmarkUnmarshalByMsgp-6                       8075196               148 ns/op
BenchmarkUnmarshalByProtoBuf-6                   3069420               382 ns/op
BenchmarkUnmarshalByGogoProtoBuf-6               4064895               304 ns/op
BenchmarkUnmarshalByThrift-6                     1405612               856 ns/op
BenchmarkUnmarshalByAvro-6                         76148             16803 ns/op
BenchmarkUnmarshalByGencode-6                   12796027                92.0 ns/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-6         1000000              1116 ns/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-6         1000000              1221 ns/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-6         1000000              1104 ns/op
BenchmarkUnmarshalByUgorjiCodecAndJson-6          600756              1934 ns/op
BenchmarkUnmarshalByEasyjson-6                   3208314               384 ns/op
BenchmarkUnmarshalByFfjson-6                      925518              1250 ns/op
BenchmarkUnmarshalByJsoniter-6                   3092865               388 ns/op
BenchmarkUnmarshalByGJSON-6                      1000000              1276 ns/op
BenchmarkUnmarshalByGoMemdump-6                  1339832               934 ns/op
BenchmarkUnmarshalByColfer-6                     7762653               156 ns/op
BenchmarkUnmarshalByZebrapack-6                  5323958               224 ns/op
BenchmarkUnmarshalByGotiny-6                     4870737               241 ns/op
BenchmarkUnmarshalByHprose-6                     2527753               474 ns/op
BenchmarkUnmarshalBySereal-6                     2114608               565 ns/op
BenchmarkUnmarshalByVmihMsgpackv4-6              1432388               823 ns/op
BenchmarkUnmarshalByRlp-6                        1675710               706 ns/op
BenchmarkUnmarshalBySegmentioJSON-6              2772384               430 ns/op
```
