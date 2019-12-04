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
- [vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
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
- [gopkg.in/vmihailenco/msgpack.v2](http://gopkg.in/vmihailenco/msgpack.v2)
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
BenchmarkMarshalByJson-4                 	 1869745	       634 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByXml-4                  	  275817	      3859 ns/op	       137 marshaledBytes
BenchmarkMarshalByMsgp-4                 	10177845	       114 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByProtoBuf-4             	 5958943	       200 ns/op	        36.0 marshaledBytes
BenchmarkMarshalByGogoProtoBuf-4         	 8599736	       138 ns/op	        36.0 marshaledBytes
BenchmarkMarshalByThrift-4               	 2524033	       449 ns/op	        63.0 marshaledBytes
BenchmarkMarshalByAvro-4                 	 3502730	       338 ns/op	        32.0 marshaledBytes
BenchmarkMarshalByGencode-4              	26821284	        43.4 ns/op	        34.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndCbor-4   	 2152038	       555 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndMsgp-4   	 2217151	       538 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndBinc-4   	 2119299	       562 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndJson-4   	 1622152	       676 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByEasyjson-4             	 4240388	       274 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByFfjson-4               	 1321222	       902 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByJsoniter-4             	 2413506	       493 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByGoMemdump-4            	  261883	      4333 ns/op	       200 marshaledBytes
BenchmarkMarshalByColfer-4               	34064367	        33.7 ns/op	        35.0 marshaledBytes
BenchmarkMarshalByZebrapack-4            	 6609897	       252 ns/op	       109 marshaledBytes
BenchmarkMarshalByGotiny-4               	 3376788	       348 ns/op	        32.0 marshaledBytes
BenchmarkMarshalByHprose-4               	 3227366	       367 ns/op	        49.0 marshaledBytes
BenchmarkMarshalBySereal-4               	  481123	      2132 ns/op	        76.0 marshaledBytes
BenchmarkMarshalByMsgpackV2-4            	  894158	      1373 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByRlp-4                  	 2521772	       477 ns/op	        32.0 marshaledBytes
BenchmarkMarshalBySegmentioJSON-4        	 2069793	       584 ns/op	        65.0 marshaledBytes
```


**Unmarshal**

```

```