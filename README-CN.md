## Golang 序列化反序列化库的性能比较

### 测试的 Serializers

以golang自带的_encoding/json_和_encoding/xml_为基准，测试以下性能比较好的几种序列化库。

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

### 排除的 Serializers

基于 alecthomas 已有的[测试](https://github.com/alecthomas/go_serialization_benchmarks)，下面的库由于性能的原因没有进行测试。

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [gopkg.in/vmihailenco/msgpack.v2](http://gopkg.in/vmihailenco/msgpack.v2)
- [bson](http://github.com/micro/go-bson)

### 测试环境
go version: **1.10**


- 对于 `MessagePack`，你需要安装库以及利用`go generate`生成相关的类:

  ```go
  go get github.com/tinylib/msgp
  go generate
  ```

- 对于`ProtoBuf`,你需要安装[protoc编译器](https://github.com/google/protobuf/releases)，以及protoc库以及生成相关的类：

  ```go
  go get github.com/golang/protobuf
  go generate
  ```

- 对于`gogo/protobuf`,你需要安装库以及生成相关的类：

  ```go
  go get github.com/gogo/protobuf/gogoproto
  go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
  go generate
  ```

- 对于`flatbuffers`,你需要安装[flatbuffers编译器](https://github.com/google/flatbuffers/releases, 以及flatbuffers库：

  ```go
  go get github.com/google/flatbuffers/go
  go generate
  ```

- 对于`thrift`,), 你需要安装[thrift编译器](https://thrift.apache.org/download)以及thrift库：

  ```go
  go get github.com/apache/thrift/lib/go/thrift
  go generate
  ```

- 对于`Avro`,你需要安装goavro库：

    ```go
    go get github.com/linkedin/goavro
    go generate
    ```

- 对于`gencode`,你需要安装gencode库,并使用gencode库的工具产生数据对象：

  ```go
  go get github.com/andyleap/gencode
  bin\gencode.exe go -schema=gencode.schema -package gosercomp
  ```

  `gencode`也是一个高性能的编解码库，提供了代码生成工具，而且产生的数据非常的小。

- 对于`easyjson`,你需要安装easyjson库:

  ```go
  go get github.com/mailru/easyjson
  go generate
  ```

- 对于`zebraPack `,你需要安装zebraPack库,并使用zebraPack工具产生数据对象：

  ```go
  go get github.com/glycerine/zebrapack
  go generate zebrapack_data.go 
  ```

- 对于`ugorji/go/codec`,你需要安装代码生成工具和`codec`库:

```go
  go get -tags=unsafe  -u github.com/ugorji/go/codec/codecgen
  go get -tags=unsafe -u github.com/ugorji/go/codec

  codecgen.exe -o data_codec.go data.go
```

`ugorji/go/codec`是一个高性能的编解码框架，支持 msgpack、cbor、binc、json等格式。本测试中测试了 cbor  和 msgpack的编解码，可以和上面的 `tinylib/msgp`框架进行比较。

> 事实上，这里通过`go generate`生成相关的类，你也可以通过命令行生成，请参考`data.go`中的注释。 但是你需要安装相关的工具，如Thrift,并把它们加入到环境变量Path中

**运行下面的命令测试:**

```
go test -bench=. -benchmem
```

### 测试数据

所有的测试基于以下的struct,自动生成的struct， 比如protobuf也和此结构基本一致。
所以本测试的数据以小数据为主， 不同的测试数据(数据大小，数据类型)可能会导致各框架的表现不一样，注意区别。

```go
type ColorGroup struct {
    ID     int `json:"id" xml:"id,attr""`
    Name   string `json:"name" xml:"name"`
    Colors []string `json:"colors" xml:"colors"`
}
`
```

### 性能测试结果

**Marshal**

_包含序列化后的数据大小_

```
BenchmarkMarshalByJson-4                 	 1920922	       623 ns/op	        65.0 marshaledBytes	     128 B/op	       2 allocs/op
BenchmarkMarshalByXml-4                  	  292286	      3715 ns/op	       137 marshaledBytes	    4736 B/op	      11 allocs/op
BenchmarkMarshalByMsgp-4                 	10683840	       109 ns/op	        47.0 marshaledBytes	      80 B/op	       1 allocs/op
BenchmarkMarshalByProtoBuf-4             	 3971731	       304 ns/op	        36.0 marshaledBytes	      80 B/op	       2 allocs/op
BenchmarkMarshalByGogoProtoBuf-4         	 9633577	       121 ns/op	        36.0 marshaledBytes	      48 B/op	       1 allocs/op
BenchmarkMarshalByFlatBuffers-4          	 3628875	       328 ns/op	       108 marshaledBytes	      16 B/op	       1 allocs/op
BenchmarkMarshalByThrift-4               	 2702652	       440 ns/op	        63.0 marshaledBytes	      64 B/op	       1 allocs/op
BenchmarkMarshalByAvro-4                 	 3656216	       327 ns/op	        32.0 marshaledBytes	       0 B/op	       0 allocs/op
BenchmarkMarshalByGencode-4              	27901552	        42.4 ns/op	        34.0 marshaledBytes	       0 B/op	       0 allocs/op
BenchmarkMarshalByUgorjiCodecAndCbor-4   	 2023094	       589 ns/op	        47.0 marshaledBytes	      96 B/op	       2 allocs/op
BenchmarkMarshalByUgorjiCodecAndMsgp-4   	 1987399	       593 ns/op	        47.0 marshaledBytes	      96 B/op	       2 allocs/op
BenchmarkMarshalByUgorjiCodecAndBinc-4   	 1988001	       600 ns/op	        47.0 marshaledBytes	      96 B/op	       2 allocs/op
BenchmarkMarshalByUgorjiCodecAndJson-4   	 1667370	       714 ns/op	        65.0 marshaledBytes	      96 B/op	       2 allocs/op
BenchmarkMarshalByEasyjson-4             	 4324338	       272 ns/op	        65.0 marshaledBytes	     128 B/op	       1 allocs/op
BenchmarkMarshalByFfjson-4               	 1335627	       896 ns/op	        65.0 marshaledBytes	     424 B/op	       9 allocs/op
BenchmarkMarshalByJsoniter-4             	 2443514	       506 ns/op	        65.0 marshaledBytes	      88 B/op	       2 allocs/op
BenchmarkMarshalByGoMemdump-4            	  260101	      4306 ns/op	       200 marshaledBytes	    1024 B/op	      30 allocs/op
BenchmarkMarshalByColfer-4               	33450418	        34.9 ns/op	        35.0 marshaledBytes	       0 B/op	       0 allocs/op
BenchmarkMarshalByZebrapack-4            	 9203962	       134 ns/op	 322138670 marshaledBytes	      72 B/op	       0 allocs/op
BenchmarkMarshalByGotiny-4               	 3498696	       334 ns/op	        32.0 marshaledBytes	     144 B/op	       5 allocs/op
BenchmarkMarshalByHprose-4               	 3028422	       394 ns/op	 148392678 marshaledBytes	     209 B/op	       1 allocs/op
BenchmarkMarshalBySereal-4               	  512172	      2084 ns/op	        76.0 marshaledBytes	     792 B/op	      22 allocs/op
BenchmarkMarshalByMsgpackV2-4            	  715090	      1710 ns/op	        47.0 marshaledBytes	     192 B/op	       5 allocs/op
BenchmarkMarshalByRlp-4                  	 2544832	       463 ns/op	        32.0 marshaledBytes	      64 B/op	       3 allocs/op
```


**Unmarshal**

```
BenchmarkUnmarshalByJson-4                     	  569730	      1998 ns/op	     248 B/op	       9 allocs/op
BenchmarkUnmarshalByXml-4                      	   98018	     12382 ns/op	    3033 B/op	      74 allocs/op
BenchmarkUnmarshalByMsgp-4                     	 6168337	       194 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByProtoBuf-4                 	 1591766	       705 ns/op	     192 B/op	      11 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-4             	 2921624	       405 ns/op	     144 B/op	       8 allocs/op
BenchmarkUnmarshalByFlatBuffers-4              	1000000000	         0.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByFlatBuffers_withFields-4   	 9671664	       123 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByThrift-4                   	 1000000	      1037 ns/op	     416 B/op	      11 allocs/op
BenchmarkUnmarshalByAvro-4                     	   54542	     21568 ns/op	   12186 B/op	     224 allocs/op
BenchmarkUnmarshalByGencode-4                  	 9364777	       123 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-4       	 2225858	       538 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-4       	 2172009	       550 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-4       	 2227891	       537 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndJson-4       	 1713320	       702 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByEasyjson-4                 	 2470573	       482 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByFfjson-4                   	  910078	      1315 ns/op	     480 B/op	      13 allocs/op
BenchmarkUnmarshalByJsoniter-4                 	 2254818	       515 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByGJSON-4                    	  742209	      1664 ns/op	     624 B/op	       7 allocs/op
BenchmarkUnmarshalByGoMemdump-4                	  904111	      1276 ns/op	    2280 B/op	      10 allocs/op
BenchmarkUnmarshalByColfer-4                   	 5646224	       209 ns/op	      96 B/op	       6 allocs/op
BenchmarkUnmarshalByZebrapack-4                	 5051925	       232 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByGotiny-4                   	 3737158	       336 ns/op	     120 B/op	       7 allocs/op
BenchmarkUnmarshalByHprose-4                   	 1915485	       624 ns/op	     288 B/op	       9 allocs/op
BenchmarkUnmarshalBySereal-4                   	 1759366	       681 ns/op	      80 B/op	       6 allocs/op
BenchmarkUnmarshalByMsgpackv2-4                	  872976	      1425 ns/op	     232 B/op	      11 allocs/op
BenchmarkUnmarshalByRlp-4                      	 1326530	       901 ns/op	     104 B/op	      11 allocs/op
```
