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
  go get github.com/apache/thrift/lib/go/thrift
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

**Marshal**

_include marshalled bytes_

```
BenchmarkMarshalByJson-4                 	 1828470	       637 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByXml-4                  	  268522	      3851 ns/op	       137 marshaledBytes
BenchmarkMarshalByMsgp-4                 	10294527	       113 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByProtoBuf-4             	 3857157	       301 ns/op	        36.0 marshaledBytes
BenchmarkMarshalByGogoProtoBuf-4         	 9484000	       125 ns/op	        36.0 marshaledBytes
BenchmarkMarshalByFlatBuffers-4          	 3504392	       345 ns/op	       108 marshaledBytes
BenchmarkMarshalByThrift-4               	 2661844	       444 ns/op	        63.0 marshaledBytes
BenchmarkMarshalByAvro-4                 	 3197880	       343 ns/op	        32.0 marshaledBytes
BenchmarkMarshalByGencode-4              	27099154	        43.2 ns/op	        34.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndCbor-4   	 2198343	       546 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndMsgp-4   	 2246230	       519 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndBinc-4   	 2197182	       535 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByUgorjiCodecAndJson-4   	 1857007	       638 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByEasyjson-4             	 4182236	       285 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByFfjson-4               	 1308387	       913 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByJsoniter-4             	 2374874	       497 ns/op	        65.0 marshaledBytes
BenchmarkMarshalByGoMemdump-4            	  243367	      4411 ns/op	       200 marshaledBytes
BenchmarkMarshalByColfer-4               	32163198	        35.6 ns/op	        35.0 marshaledBytes
BenchmarkMarshalByZebrapack-4            	 8821429	       162 ns/op	        79.0 marshaledBytes
BenchmarkMarshalByGotiny-4               	 3007924	       419 ns/op	        32.0 marshaledBytes
BenchmarkMarshalByHprose-4               	 2850037	       802 ns/op	        49.0 marshaledBytes
BenchmarkMarshalBySereal-4               	  463419	      2545 ns/op	        76.0 marshaledBytes
BenchmarkMarshalByMsgpackV2-4            	  589584	      1905 ns/op	        47.0 marshaledBytes
BenchmarkMarshalByRlp-4                  	 2334664	       518 ns/op	        32.0 marshaledBytes
BenchmarkMarshalBySegmentioJSON-4        	 1834310	       638 ns/op	        65.0 marshaledBytes
```


**Unmarshal**

```
BenchmarkUnmarshalByJson-4                     	  519876	      2052 ns/op
BenchmarkUnmarshalByXml-4                      	   93301	     12936 ns/op
BenchmarkUnmarshalByMsgp-4                     	 5782575	       197 ns/op
BenchmarkUnmarshalByProtoBuf-4                 	 1668422	       724 ns/op
BenchmarkUnmarshalByGogoProtoBuf-4             	 2905324	       408 ns/op
BenchmarkUnmarshalByFlatBuffers-4              	1000000000	         0.495 ns/op
BenchmarkUnmarshalByFlatBuffers_withFields-4   	 9554569	       125 ns/op
BenchmarkUnmarshalByThrift-4                   	 1000000	      1164 ns/op
BenchmarkUnmarshalByAvro-4                     	   53198	     22053 ns/op
BenchmarkUnmarshalByGencode-4                  	 9343574	       125 ns/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-4       	 2124366	       560 ns/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-4       	 2170894	       528 ns/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-4       	 2295622	       520 ns/op
BenchmarkUnmarshalByUgorjiCodecAndJson-4       	 1737640	       683 ns/op
BenchmarkUnmarshalByEasyjson-4                 	 2410402	       514 ns/op
BenchmarkUnmarshalByFfjson-4                   	  861022	      1354 ns/op
BenchmarkUnmarshalByJsoniter-4                 	 2239983	       536 ns/op
BenchmarkUnmarshalByGJSON-4                    	  712545	      1732 ns/op
BenchmarkUnmarshalByGoMemdump-4                	  910912	      1342 ns/op
BenchmarkUnmarshalByColfer-4                   	 5514124	       217 ns/op
BenchmarkUnmarshalByZebrapack-4                	 4782226	       245 ns/op
BenchmarkUnmarshalByGotiny-4                   	 3587878	       326 ns/op
BenchmarkUnmarshalByHprose-4                   	 1844857	       649 ns/op
BenchmarkUnmarshalBySereal-4                   	 1734404	       685 ns/op
BenchmarkUnmarshalByMsgpackv2-4                	  861864	      1456 ns/op
BenchmarkUnmarshalByRlp-4                      	 1232529	       962 ns/op
BenchmarkUnmarshalBySegmentioJSON-4            	 2054912	       578 ns/op
```