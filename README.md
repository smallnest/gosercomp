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

    
### 性能测试结果


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

