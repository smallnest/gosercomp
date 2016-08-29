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

### 排除的 Serializers

基于 alecthomas 已有的[测试](https://github.com/alecthomas/go_serialization_benchmarks)，下面的库由于性能的原因没有进行测试。

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [github.com/ugorji/go/codec](http://github.com/ugorji/go/codec)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [gopkg.in/vmihailenco/msgpack.v2](http://gopkg.in/vmihailenco/msgpack.v2)
- [bson](http://github.com/micro/go-bson)

### 测试环境

- 对于`github.com/youtube/vitess/go/bson`，你可能需要安装 `goimports`和`codegen`:

  ```go
  go get github.com/youtube/vitess/go/bson
  go get golang.org/x/tools/cmd/goimports
  go get github.com/youtube/vitess/tree/master/go/cmd/bsongen
  bsongen -file data.go -o bson_data.go -type ColorGroup
  ```

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

- 对于`flatbuffers`,你需要安装[thrift编译器](https://thrift.apache.org/download), 以及flatbuffers库：

  ```go
  go get github.com/google/flatbuffers/go
  go generate
  ```

- 对于`thrift`,你需要安装[flatbuffers编译器](https://github.com/google/flatbuffers/releases), 以及thrift库：

  ```go
  go get git.apache.org/thrift.git/lib/go/thrift
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

```go
type ColorGroup struct {
    ID     int `json:"id" xml:"id,attr""`
    Name   string `json:"name" xml:"name"`
    Colors []string `json:"colors" xml:"colors"`
}
`
```

### 性能测试结果

```
benchmark _name                                  iter          time/iter        alloc bytes/iter    allocs/iter
-------------------------------------------------------------------------------------------------------------------------
BenchmarkMarshalByJson-8                       	 1000000	      1137 ns/op	     376 B/op	       4 allocs/op
BenchmarkUnmarshalByJson-8                     	 1000000	      2260 ns/op	     296 B/op	       9 allocs/op
BenchmarkMarshalByXml-8                        	  500000	      3351 ns/op	    4801 B/op	      12 allocs/op
BenchmarkUnmarshalByXml-8                      	  100000	     13278 ns/op	    2807 B/op	      67 allocs/op
BenchmarkMarshalByMsgp-8                       	20000000	       102 ns/op	      80 B/op	       1 allocs/op
BenchmarkUnmarshalByMsgp-8                     	 5000000	       266 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByProtoBuf-8                   	 3000000	       450 ns/op	     328 B/op	       5 allocs/op
BenchmarkUnmarshalByProtoBuf-8                 	 2000000	       754 ns/op	     400 B/op	      11 allocs/op
BenchmarkMarshalByGogoProtoBuf-8               	20000000	        99.0 ns/op	      48 B/op	       1 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-8             	 5000000	       389 ns/op	     144 B/op	       8 allocs/op
BenchmarkMarshalByFlatBuffers-8                	 5000000	       352 ns/op	      16 B/op	       1 allocs/op
BenchmarkUnmarshalByFlatBuffers-8              	500000000	         3.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByFlatBuffers_withFields-8   	10000000	       157 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalByThrift-8                     	 3000000	       517 ns/op	      64 B/op	       1 allocs/op
BenchmarkUnmarshalByThrift-8                   	 1000000	      1224 ns/op	     656 B/op	      11 allocs/op
BenchmarkMarshalByAvro-8                       	 2000000	       889 ns/op	     133 B/op	       7 allocs/op
BenchmarkUnmarshalByAvro-8                     	  500000	      3128 ns/op	    1680 B/op	      63 allocs/op
BenchmarkMarshalByGencode-8                    	50000000	        37.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnmarshalByGencode-8                  	10000000	       139 ns/op	      32 B/op	       5 allocs/op
BenchmarkMarshalByCodecAndCbor-8               	 2000000	       697 ns/op	     239 B/op	       2 allocs/op
BenchmarkUnmarshalByCodecAndCbor-8             	10000000	       149 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalByCodecAndMsgp-8               	 2000000	       676 ns/op	     239 B/op	       2 allocs/op
BenchmarkUnmarshalByCodecAndMsgp-8             	10000000	       158 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalByGoMemdump-8                  	  300000	      5217 ns/op	    1487 B/op	      31 allocs/op
BenchmarkUnmarshalByGoMemdump-8                	 3000000	       474 ns/op	     112 B/op	       5 allocs/op
```

多次测试结果差不多。 从结果上上来看， **MessagePack**,**gogo/protobuf**,和**flatbuffers**差不多，这三个优秀的库在序列化和反序列化上各有千秋，而且都是跨语言的。 从便利性上来讲，你可以选择**MessagePack**和**gogo/protobuf**都可以，两者都有大厂在用。 flatbuffers有点反人类，因为它的操作很底层，而且从结果上来看，序列化的性能要差一点。但是它有一个好处，那就是如果你只需要特定的字段， 你无须将所有的字段都反序列化。从结果上看，不反序列化字段每个调用只用了9.54纳秒，这是因为字段只有在被访问的时候才从byte数组转化为相应的类型。 因此在特殊的场景下，它可以提高N被的性能。但是序列化的代码的面相太难看了。

新增加了**gencode**的测试，它表现相当出色，而且生成的字节也非常的小。

**Codec**的Unmarshal性能不错，但是Marshal性能不是太好。
