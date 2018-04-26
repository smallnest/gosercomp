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
BenchmarkMarshalByJson-4                       	 2000000	       870 ns/op	     368 B/op	       3 allocs/op
BenchmarkUnmarshalByJson-4                     	  500000	      2496 ns/op	     344 B/op	       9 allocs/op
BenchmarkMarshalByXml-4                        	  300000	      3857 ns/op	    4800 B/op	      11 allocs/op
BenchmarkUnmarshalByXml-4                      	  100000	     12779 ns/op	    3171 B/op	      75 allocs/op
BenchmarkMarshalByMsgp-4                       	20000000	       109 ns/op	      80 B/op	       1 allocs/op
BenchmarkUnmarshalByMsgp-4                     	10000000	       219 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByProtoBuf-4                   	 3000000	       469 ns/op	     328 B/op	       5 allocs/op
BenchmarkUnmarshalByProtoBuf-4                 	 2000000	       794 ns/op	     400 B/op	      11 allocs/op
BenchmarkMarshalByGogoProtoBuf-4               	20000000	       106 ns/op	      48 B/op	       1 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-4             	 3000000	       411 ns/op	     144 B/op	       8 allocs/op
BenchmarkMarshalByFlatBuffers-4                	 5000000	       373 ns/op	      16 B/op	       1 allocs/op
BenchmarkUnmarshalByFlatBuffers-4              	2000000000	         0.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByFlatBuffers_withFields-4   	10000000	       156 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalByThrift-4                     	 3000000	       430 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnmarshalByThrift-4                   	 1000000	      1264 ns/op	     656 B/op	      11 allocs/op
BenchmarkMarshalByAvro-4                       	 3000000	       536 ns/op	      48 B/op	       6 allocs/op
BenchmarkUnmarshalByAvro-4                     	  500000	      3183 ns/op	    1672 B/op	      62 allocs/op
BenchmarkMarshalByGencode-4                    	30000000	        40.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByGencode-4                  	10000000	       120 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByUgorjiCodecAndCbor-4         	 2000000	       696 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-4       	 3000000	       603 ns/op	      48 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndMsgp-4         	 2000000	       663 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-4       	 2000000	       609 ns/op	      48 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndBinc-4         	 2000000	       706 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-4       	 1000000	      1069 ns/op	     824 B/op	      10 allocs/op
BenchmarkMarshalByUgorjiCodecAndJson-4         	 2000000	       865 ns/op	     112 B/op	       3 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndJson-4       	 2000000	       764 ns/op	      48 B/op	       6 allocs/op
BenchmarkMarshalByEasyjson-4                   	 5000000	       316 ns/op	     128 B/op	       1 allocs/op
BenchmarkUnmarshalByEasyjson-4                 	 3000000	       481 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByFfjson-4                     	 2000000	       930 ns/op	     424 B/op	       9 allocs/op
BenchmarkUnmarshalByFfjson-4                   	 1000000	      1365 ns/op	     480 B/op	      13 allocs/op
BenchmarkMarshalByJsoniter-4                   	 2000000	       761 ns/op	     800 B/op	       5 allocs/op
BenchmarkUnmarshalByJsoniter-4                 	 3000000	       452 ns/op	     112 B/op	       6 allocs/op
BenchmarkUnmarshalByGJSON-4                    	 1000000	      1807 ns/op	     624 B/op	       7 allocs/op
BenchmarkMarshalByGoMemdump-4                  	  300000	      4862 ns/op	    1032 B/op	      30 allocs/op
BenchmarkUnmarshalByGoMemdump-4                	 1000000	      1462 ns/op	    2400 B/op	      12 allocs/op
BenchmarkMarshalByColfer-4                     	50000000	        29.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByColfer-4                   	10000000	       200 ns/op	      96 B/op	       6 allocs/op
BenchmarkMarshalByZebrapack-4                  	20000000	       278 ns/op	     132 B/op	       0 allocs/op
BenchmarkUnmarshalByZebrapack-4                	 5000000	       244 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByGotiny-4                     	 5000000	       363 ns/op	     144 B/op	       5 allocs/op
BenchmarkUnmarshalByGotiny-4                   	 5000000	       262 ns/op	      88 B/op	       2 allocs/op
BenchmarkMarshalByHprose-4                     	 3000000	       493 ns/op	     210 B/op	       1 allocs/op
BenchmarkUnmarshalByHprose-4                   	 2000000	       641 ns/op	     288 B/op	       9 allocs/op
BenchmarkMarshalBySereal-4                     	 1000000	      2169 ns/op	     792 B/op	      22 allocs/op
BenchmarkUnmarshalBySereal-4                   	 2000000	       720 ns/op	      80 B/op	       6 allocs/op
BenchmarkMarshalByMsgpackV2-4                  	 1000000	      1872 ns/op	     192 B/op	       4 allocs/op
BenchmarkUnmarshalByMsgpackv2-4                	 1000000	      1603 ns/op	     232 B/op	      11 allocs/op
```


## Size of marshalled results


```
	gosercomp_test.go:91: json:				 65 bytes
	gosercomp_test.go:94: xml:				 137 bytes
	gosercomp_test.go:97: msgp:				 47 bytes
	gosercomp_test.go:100: protobuf:				 36 bytes
	gosercomp_test.go:103: gogoprotobuf:			 36 bytes
	gosercomp_test.go:107: flatbuffers:			 108 bytes
	gosercomp_test.go:113: thrift:				 63 bytes
	gosercomp_test.go:127: avro:				 32 bytes
	gosercomp_test.go:136: gencode:				 34 bytes
	gosercomp_test.go:142: UgorjiCodec_Cbor:		 47 bytes
	gosercomp_test.go:148: UgorjiCodec_Msgp:		 47 bytes
	gosercomp_test.go:154: UgorjiCodec_Bin:			 53 bytes
	gosercomp_test.go:156: UgorjiCodec_Json:		 91 bytes
	gosercomp_test.go:159: easyjson:			 65 bytes
	gosercomp_test.go:162: ffjson:				 65 bytes
	gosercomp_test.go:165: jsoniter:			 65 bytes
	gosercomp_test.go:169: memdump:				 200 bytes
	gosercomp_test.go:172: colfer:				 35 bytes
	gosercomp_test.go:175: zebrapack:			 35 bytes
	gosercomp_test.go:178: gotiny:				 32 bytes
	gosercomp_test.go:183: hprose:				 32 bytes
	gosercomp_test.go:187: sereal:				 76 bytes
	gosercomp_test.go:190: msgpackv2:			 47 bytes
```