## Golang 序列化反序列化库的性能比较


### 测试的 Serializers
以golang自带的*encoding/json*和*encoding/xml*为基准，测试以下性能比较好的几种序列化库。
* [encoding/json](http://golang.org/pkg/encoding/json/)
* [encoding/xml](http://golang.org/pkg/encoding/xml/)
* [github.com/youtube/vitess/go/bson](http://github.com/youtube/vitess/go/bson)
* [github.com/philhofer/msgp](http://github.com/philhofer/msgp)
* [github.com/golang/protobuf](http://github.com/golang/protobuf)
* [github.com/gogo/protobuf](http://github.com/gogo/protobuf)
* [github.com/google/flatbuffers](http://github.com/google/flatbuffers)


### 排除的 Serializers
基于 alecthomas 已有的[测试](https://github.com/alecthomas/go_serialization_benchmarks)，下面的库由于性能的原因没有进行测试。

* [encoding/gob](http://golang.org/pkg/encoding/gob/)
* [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
* [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
* [github.com/ugorji/go/codec](http://github.com/ugorji/go/codec)
* [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
* [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
* [gopkg.in/vmihailenco/msgpack.v2](http://gopkg.in/vmihailenco/msgpack.v2)

### 测试环境
对于`github.com/youtube/vitess/go/bson`，你可能需要安装 `goimports`和`codegen`:
``` go
go get github.com/youtube/vitess/go/bson
go get golang.org/x/tools/cmd/goimports
go get github.com/youtube/vitess/tree/master/go/cmd/bsongen
bsongen -file data.go -o bson_data.go -type ColorGroup
```
对于 `MessagePack`，你需要安装库以及利用`go generate`生成相关的类:
``` go
go get github.com/tinylib/msgp
go generate
```

对于`ProtoBuf`,你需要安装[protoc编译器](https://github.com/google/protobuf/releases)，以及protoc库以及生成相关的类：
``` go
go get github.com/golang/protobuf
go generate
```

对于`gogo/protobuf`,你需要安装库以及生成相关的类：
``` go
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/protobuf/protoc-gen-gofast
go generate
```

对于`flatbuffers`,你需要安装[flatbuffers编译器](https://github.com/google/flatbuffers/releases), 以及flatbuffers库：
``` go
github.com/google/flatbuffers/go
go generate
```

事实上，这里通过`go generate`生成相关的类，你也可以通过命令行生成，请参考`data.go`中的注释。


运行下面的命令测试:
```
go test -bench=.
```
### 测试数据
所有的测试基于以下的struct,自动生成的struct， 比如protobuf也和此结构基本一致。
``` go
type ColorGroup struct {
	ID     int `json:"id" xml:"id,attr""`
	Name   string `json:"name" xml:"name"`
	Colors []string `json:"colors" xml:"colors"`
}
```

    
### 性能测试结果
<pre>
BenchmarkMarshalByJson-4                      1000000              1877 ns/op
BenchmarkUnmarshalByJson-4                  300000                4099 ns/op

BenchmarkMarshalByXml-4                       200000                8315 ns/op
BenchmarkUnmarshalByXml-4                   100000                26627 ns/op

BenchmarkMarshalByBson-4                      500000                3518 ns/op
BenchmarkUnmarshalByBson-4                  1000000              1778 ns/op

BenchmarkMarshalByMsgp-4                     5000000              292 ns/op
BenchmarkUnmarshalByMsgp-4                 3000000              543 ns/op

BenchmarkMarshalByProtoBuf-4                1000000              1011 ns/op
BenchmarkUnmarshalByProtoBuf-4            1000000              1750 ns/op

BenchmarkMarshalByGogoProtoBuf-4        5000000              220 ns/op
BenchmarkUnmarshalByGogoProtoBuf-4    2000000              901 ns/op

BenchmarkMarshalByFlatBuffers-4              3000000              566 ns/op
BenchmarkUnmarshalByFlatBuffers-4          50000000            9.54 ns/op
BenchmarUmByFlatBuffers_withFields-4      3000000              554 ns/op
</pre>

多次测试结果差不多。
从结果上上来看， **MessagePack**,**gogo/protobuf**,和**flatbuffers**差不多，这三个优秀的库在序列化和反序列化上各有千秋，而且都是跨语言的。
从便利性上来讲，你可以选择**MessagePack**和**gogo/protobuf**都可以，两者都有大厂在用。
flatbuffers有点反人类，因为它的操作很底层，而且从结果上来看，序列化的性能要差一点。但是它有一个好处，那就是如果你只需要特定的字段，
你无须将所有的字段都反序列化。从结果上看，不反序列化字段每个调用只用了9.54纳秒，这是因为字段只有在被访问的时候才从byte数组转化为相应的类型。
因此在特殊的场景下，它可以提高N被的性能。但是序列化的代码的面相太难看了。



