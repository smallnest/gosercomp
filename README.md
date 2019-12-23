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
BenchmarkMarshalByJson-4                 	 1950553	       780 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByXml-4                  	  317310	      3780 ns/op	       137 marshaledBytes
BenchmarkMarshalByMsgp-4                 	10264906	       109 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByProtoBuf-4             	 6027548	       254 ns/op	        36.0 marshaledBytes
BenchmarkMarshalByGogoProtoBuf-4         	 7823614	       136 ns/op	        36.0 marshaledBytes
BenchmarkMarshalByThrift-4               	 2713792	       443 ns/op	        63.0 marshaledBytes
BenchmarkMarshalByAvro-4                 	 2710900	       577 ns/op	        32.0 marshaledBytes
BenchmarkMarshalByGencode-4              	25360005	        42.6 ns/op	        34.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndCbor-4   	  736860	      1488 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndMsgp-4   	  927762	      1474 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndBinc-4   	  976227	      1198 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndJson-4   	  841494	      1412 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByEasyjson-4             	 4410404	       379 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByFfjson-4               	 1284018	       968 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByJsoniter-4             	 2420109	       678 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByGoMemdump-4            	  146306	      8485 ns/op	       200 marshaledBytes
BenchmarkMarshalByColfer-4               	31515162	        42.1 ns/op	        35.0 marshaledBytes
BenchmarkMarshalByZebrapack-4            	 6037338	       391 ns/op	       109 marshaledBytes
BenchmarkMarshalByGotiny-4               	 3053563	       526 ns/op	        32.0 marshaledBytes
BenchmarkMarshalByHprose-4               	 2981386	       351 ns/op	        49.0 marshaledBytes
BenchmarkMarshalBySereal-4               	  454028	      3649 ns/op	        76.0 marshaledBytes
BenchmarkMarshalByVmihMsgpackv4-4        	 1531735	       760 ns/op	        55.0 marshaledBytes
BenchmarkMarshalByRlp-4                  	 2019198	       792 ns/op	        32.0 marshaledBytes
BenchmarkMarshalBySegmentioJSON-4        	 1995334	      1568 ns/op	        65.0 marshaledBytes
```


**Unmarshal**

```
BenchmarkUnmarshalByJson-4                 	  430582	      3001 ns/op
BenchmarkUnmarshalByXml-4                  	   86936	     12710 ns/op
BenchmarkUnmarshalByMsgp-4                 	 5982232	       196 ns/op
BenchmarkUnmarshalByProtoBuf-4             	 2427032	       573 ns/op
BenchmarkUnmarshalByGogoProtoBuf-4         	 2964386	       393 ns/op
BenchmarkUnmarshalByThrift-4               	  872664	      1147 ns/op
BenchmarkUnmarshalByAvro-4                 	   53508	     29774 ns/op
BenchmarkUnmarshalByGencode-4              	 9607225	       123 ns/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-4   	  682351	      1526 ns/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-4   	  607396	      1650 ns/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-4   	  530145	      2313 ns/op
BenchmarkUnmarshalByUgorjiCodecAndJson-4   	  445137	      2759 ns/op
BenchmarkUnmarshalByEasyjson-4             	 2448640	       588 ns/op
BenchmarkUnmarshalByFfjson-4               	  904053	      1337 ns/op
BenchmarkUnmarshalByJsoniter-4             	 2127112	       551 ns/op
BenchmarkUnmarshalByGJSON-4                	  578313	      2776 ns/op
BenchmarkUnmarshalByGoMemdump-4            	  910186	      1298 ns/op
BenchmarkUnmarshalByColfer-4               	 5520969	       254 ns/op
BenchmarkUnmarshalByZebrapack-4            	 4125756	       290 ns/op
BenchmarkUnmarshalByGotiny-4               	 3579372	       333 ns/op
BenchmarkUnmarshalByHprose-4               	 1883326	       816 ns/op
BenchmarkUnmarshalBySereal-4               	 1772971	       673 ns/op
BenchmarkUnmarshalByVmihMsgpackv4-4        	  950239	      1070 ns/op
BenchmarkUnmarshalByRlp-4                  	 1000000	      1221 ns/op
BenchmarkUnmarshalBySegmentioJSON-4        	 2093215	       570 ns/op
```
