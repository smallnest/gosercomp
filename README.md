## Golang Serialization Benchmark

### Serializers

This project test the below go serializers, which compares with go standard _json_ and _xml_.

- [encoding/json](http://golang.org/pkg/encoding/json/)
- [encoding/xml](http://golang.org/pkg/encoding/xml/)
- [github.com/tinylib/msgp](http://github.com/tinylib/msgp)
- [github.com/golang/protobuf](http://github.com/golang/protobuf)
- [github.com/gogo/protobuf](http://github.com/gogo/protobuf)
- [github.com/google/flatbuffers](http://github.com/google/flatbuffers)
- [Apache/Thrift](https://github.com/apache/thrift/tree/master/lib/go)
- [Apache/Avro](https://github.com/linkedin/goavro)
- [andyleap/gencode](https://github.com/andyleap/gencode)
- [ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [go-memdump](https://github.com/alexflint/go-memdump)
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

Given existed [benchmark](https://github.com/alecthomas/go_serialization_benchmarks) by alecthomas，the below serializers are excluded from this test because of their poor performance.

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [gopkg.in/vmihailenco/msgpack.v2](http://gopkg.in/vmihailenco/msgpack.v2)
- [bson](http://github.com/micro/go-bson)

### Test Environment
go version: **1.10**


- For `MessagePack`，you need install the tool and use `go generate` to generate code:

  ```go
  go get github.com/tinylib/msgp
  go generate
  ```

- For `ProtoBuf`, you need to install [protoc](https://github.com/google/protobuf/releases)，protobuf lib and generate code：

  ```go
  go get github.com/golang/protobuf
  go generate
  ```

- For `gogo/protobuf`, use the below commands：

  ```go
  go get github.com/gogo/protobuf/gogoproto
  go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
  go generate
  ```

- For `flatbuffers`, you need to install [flatbuffers compiler](https://github.com/google/flatbuffers/releases,  and flatbuffers lib：

  ```go
  go get github.com/google/flatbuffers/go
  go generate
  ```

- For `thrift`, you need to install [thrift compiler](https://thrift.apache.org/download), and thrift lib：

  ```go
  go get git.apache.org/thrift.git/lib/go/thrift
  go generate
  ```

- For `Avro`, you need to install goavro：

    ```go
    go get github.com/linkedin/goavro
    go generate
    ```

- For `gencode`, you need to install gencode, and geneate code by gencode：

  ```go
  go get github.com/andyleap/gencode
  bin\gencode.exe go -schema=gencode.schema -package gosercomp
  ```


- For `easyjson`, you need to install easyjson:

  ```go
  go get github.com/mailru/easyjson
  go generate
  ```

- For `zebraPack `, you need to install zebraPack, and generate code：

  ```go
  go get github.com/glycerine/zebrapack
  go generate zebrapack_data.go 
  ```

- For `ugorji/go/codec` you need to install codecgen and `codec` lib:

```go
  go get -tags=unsafe  -u github.com/ugorji/go/codec/codecgen
  go get -tags=unsafe -u github.com/ugorji/go/codec

  codecgen.exe -o data_codec.go data.go
```


`ugorji/go/codec` supports msgpack、cbor、binc、json, and this project test its  cbor and msgpack.

> Actually，you can use `go generate` to generate code. 

**Test:**

```
go test -bench=. -benchmem
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

```
BenchmarkMarshalByJson-4                       	 2000000	       713 ns/op
BenchmarkUnmarshalByJson-4                     	  500000	      2291 ns/op

BenchmarkMarshalByXml-4                        	  300000	      4140 ns/op
BenchmarkUnmarshalByXml-4                      	  100000	     13508 ns/op

BenchmarkMarshalByMsgp-4                       	10000000	       119 ns/op
BenchmarkUnmarshalByMsgp-4                     	10000000	       222 ns/op

BenchmarkMarshalByProtoBuf-4                   	 5000000	       307 ns/op
BenchmarkUnmarshalByProtoBuf-4                 	 2000000	       766 ns/op

BenchmarkMarshalByGogoProtoBuf-4               	10000000	       109 ns/op
BenchmarkUnmarshalByGogoProtoBuf-4             	 5000000	       398 ns/op

BenchmarkMarshalByFlatBuffers-4                	 5000000	       346 ns/op
BenchmarkUnmarshalByFlatBuffers-4              	2000000000	         0.87 ns/op
BenchmarkUnmarshalByFlatBuffers_withFields-4   	10000000	       147 ns/op

BenchmarkMarshalByThrift-4                     	 3000000	       445 ns/op
BenchmarkUnmarshalByThrift-4                   	 1000000	      1336 ns/op

BenchmarkMarshalByAvro-4                       	 3000000	       509 ns/op
BenchmarkUnmarshalByAvro-4                     	  500000	      3259 ns/op

BenchmarkMarshalByGencode-4                    	30000000	        42.3 ns/op
BenchmarkUnmarshalByGencode-4                  	10000000	       120 ns/op

BenchmarkMarshalByUgorjiCodecAndCbor-4         	 2000000	       731 ns/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-4       	 2000000	       614 ns/op

BenchmarkMarshalByUgorjiCodecAndMsgp-4         	 2000000	       665 ns/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-4       	 2000000	       637 ns/op

BenchmarkMarshalByUgorjiCodecAndBinc-4         	 2000000	       690 ns/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-4       	 1000000	      1075 ns/op

BenchmarkMarshalByUgorjiCodecAndJson-4         	 2000000	       964 ns/op
BenchmarkUnmarshalByUgorjiCodecAndJson-4       	 2000000	       768 ns/op

BenchmarkMarshalByEasyjson-4                   	 5000000	       313 ns/op
BenchmarkUnmarshalByEasyjson-4                 	 3000000	       474 ns/op

BenchmarkMarshalByFfjson-4                     	 2000000	       943 ns/op
BenchmarkUnmarshalByFfjson-4                   	 1000000	      1386 ns/op

BenchmarkMarshalByJsoniter-4                   	 2000000	       751 ns/op
BenchmarkUnmarshalByJsoniter-4                 	 3000000	       457 ns/op

BenchmarkUnmarshalByGJSON-4                    	 1000000	      1733 ns/op

BenchmarkMarshalByGoMemdump-4                  	  300000	      4840 ns/op
BenchmarkUnmarshalByGoMemdump-4                	 1000000	      1445 ns/op

BenchmarkMarshalByColfer-4                     	50000000	        32.9 ns/op
BenchmarkUnmarshalByColfer-4                   	10000000	       206 ns/op

BenchmarkMarshalByZebrapack-4                  	10000000	       287 ns/op
BenchmarkUnmarshalByZebrapack-4                	 5000000	       242 ns/op

BenchmarkMarshalByGotiny-4                     	 5000000	       363 ns/op
BenchmarkUnmarshalByGotiny-4                   	 5000000	       264 ns/op

BenchmarkMarshalByHprose-4                     	 3000000	       517 ns/op
BenchmarkUnmarshalByHprose-4                   	 2000000	       673 ns/op

BenchmarkMarshalBySereal-4                     	 1000000	      2197 ns/op
BenchmarkUnmarshalBySereal-4                   	 2000000	       727 ns/op

BenchmarkMarshalByMsgpackV2-4                  	 1000000	      1840 ns/op
BenchmarkUnmarshalByMsgpackv2-4                	 1000000	      1585 ns/op

BenchmarkMarshalByRlp-4                        	 3000000	       518 ns/op
BenchmarkUnmarshalByRlp-4                      	 1000000	      1104 ns/op
```


## Size of marshalled results


```
    gosercomp_test.go:98: json:				     65 bytes
    gosercomp_test.go:101: xml:				     137 bytes
    gosercomp_test.go:104: msgp:				 47 bytes
    gosercomp_test.go:107: protobuf:		     36 bytes
    gosercomp_test.go:110: gogoprotobuf:		 36 bytes
    gosercomp_test.go:114: flatbuffers:			 108 bytes
    gosercomp_test.go:120: thrift:				 63 bytes
    gosercomp_test.go:134: avro:				 32 bytes
    gosercomp_test.go:143: gencode:				 34 bytes
    gosercomp_test.go:149: UgorjiCodec_Cbor:	 47 bytes
    gosercomp_test.go:155: UgorjiCodec_Msgp:	 47 bytes
    gosercomp_test.go:161: UgorjiCodec_Bin:		 53 bytes
    gosercomp_test.go:163: UgorjiCodec_Json:	 91 bytes
    gosercomp_test.go:166: easyjson:			 65 bytes
    gosercomp_test.go:169: ffjson:				 65 bytes
    gosercomp_test.go:172: jsoniter:			 65 bytes
    gosercomp_test.go:176: memdump:				 200 bytes
    gosercomp_test.go:179: colfer:				 35 bytes
    gosercomp_test.go:182: zebrapack:			 35 bytes
    gosercomp_test.go:185: gotiny:				 32 bytes
    gosercomp_test.go:190: hprose:				 32 bytes
    gosercomp_test.go:194: sereal:				 76 bytes
    gosercomp_test.go:197: msgpackv2:			 47 bytes
    gosercomp_test.go:203: rlp:			         32 bytes
```