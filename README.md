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
- [go-memdump](https://github.com/alexflint/go-memdump)

### Excluded Serializers

Given existed [benchmark](https://github.com/alecthomas/go_serialization_benchmarks) by alecthomas，or complexity， or activity, the below serializers are excluded from this test because of their poor performance.

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [github.com/alecthomas/binary](http://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](http://github.com/davecgh/go-xdr/xdr)
- [labix.org/v2/mgo/bson](http://labix.org/v2/mgo/bson)
- [github.com/DeDiS/protobuf](http://github.com/DeDiS/protobuf)
- [bson](http://github.com/micro/go-bson)
- [github.com/google/flatbuffers](http://github.com/google/flatbuffers)

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

![](benchmark.png)

**Marshal**

_include marshalled bytes_
The test machine is the MacBook Pro 16 with i7 2.6GHz and 32G 2667MHz DDR4 memory.
The test command we use is `go test -benchtime=5s -bench=. -benchmem`

```
BenchmarkMarshalByJson-128                        	14734310	       410.2 ns/op	        65.00 marshaledBytes	     128 B/op	       2 allocs/op
BenchmarkMarshalByXml-128                         	 1578195	      3735 ns/op	       137.0 marshaledBytes	    4736 B/op	      11 allocs/op
BenchmarkMarshalByMsgp-128                        	60867415	        90.96 ns/op	        47.00 marshaledBytes	      80 B/op	       1 allocs/op
BenchmarkMarshalByProtoBuf-128                    	18273046	       327.0 ns/op	        36.00 marshaledBytes	      64 B/op	       2 allocs/op
BenchmarkMarshalByGogoProtoBuf-128                	65710664	        90.39 ns/op	        36.00 marshaledBytes	      48 B/op	       1 allocs/op
BenchmarkMarshalByThrift-128                      	21794796	       273.7 ns/op	        63.00 marshaledBytes	      64 B/op	       1 allocs/op
BenchmarkMarshalByThriftIterator-128              	14869611	       402.6 ns/op	        63.00 marshaledBytes	     248 B/op	       6 allocs/op
BenchmarkMarshalByThriftIteratorDynamic-128       	14203106	       423.2 ns/op	        63.00 marshaledBytes	     200 B/op	       5 allocs/op
BenchmarkMarshalByThriftIteratorEncoder-128       	22899847	       295.3 ns/op	        63.00 marshaledBytes	     187 B/op	       0 allocs/op
BenchmarkMarshalByAvro-128                        	20259181	       297.5 ns/op	        32.00 marshaledBytes	     112 B/op	       2 allocs/op
BenchmarkMarshalByGencode-128                     	165538641	        35.88 ns/op	        34.00 marshaledBytes	       0 B/op	       0 allocs/op
BenchmarkMarshalByUgorjiCodecAndCbor-128          	 5353467	      1103 ns/op	        47.00 marshaledBytes	    1504 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndMsgp-128          	 5484471	      1093 ns/op	        47.00 marshaledBytes	    1504 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndBinc-128          	 5455846	      1104 ns/op	        47.00 marshaledBytes	    1504 B/op	       6 allocs/op
BenchmarkMarshalByUgorjiCodecAndJson-128          	 4621263	      1291 ns/op	        65.00 marshaledBytes	    1584 B/op	       6 allocs/op
BenchmarkMarshalByEasyjson-128                    	27812725	       215.5 ns/op	        65.00 marshaledBytes	     128 B/op	       1 allocs/op
BenchmarkMarshalByFfjson-128                      	 5023614	      1190 ns/op	        65.00 marshaledBytes	     412 B/op	       9 allocs/op
BenchmarkMarshalByJsoniter-128                    	17086744	       350.6 ns/op	        65.00 marshaledBytes	      88 B/op	       2 allocs/op
BenchmarkMarshalBySonic-128                       	20919127	       288.6 ns/op	        65.00 marshaledBytes	     198 B/op	       4 allocs/op
BenchmarkMarshalByGojay-128                       	13831465	       430.9 ns/op	        65.00 marshaledBytes	     538 B/op	       2 allocs/op
BenchmarkMarshalByGoMemdump-128                   	 1550764	      3864 ns/op	       200.0 marshaledBytes	    1512 B/op	      27 allocs/op
BenchmarkMarshalByColfer-128                      	216226507	        27.93 ns/op	        35.00 marshaledBytes	       0 B/op	       0 allocs/op
BenchmarkMarshalByZebrapack-128                   	46097662	       133.0 ns/op	       109.0 marshaledBytes	     186 B/op	       0 allocs/op
BenchmarkMarshalByHprose-128                      	25639276	       233.2 ns/op	        49.00 marshaledBytes	      24 B/op	       1 allocs/op
BenchmarkMarshalBySereal-128                      	 3448252	      1724 ns/op	        76.00 marshaledBytes	     728 B/op	      22 allocs/op
BenchmarkMarshalByVmihMsgpackv4-128               	11608585	       516.2 ns/op	        55.00 marshaledBytes	     232 B/op	       5 allocs/op
BenchmarkMarshalByRlp-128                         	22588225	       266.8 ns/op	        32.00 marshaledBytes	      64 B/op	       3 allocs/op
BenchmarkMarshalBySegmentioJSON-128               	 9633194	       623.1 ns/op	        65.00 marshaledBytes	    1072 B/op	       2 allocs/op

```

**Unmarshal**

```
BenchmarkUnmarshalByJson-128                      	 3697374	      1631 ns/op	     264 B/op	      10 allocs/op
BenchmarkUnmarshalByXml-128                       	  638480	      9316 ns/op	    2946 B/op	      70 allocs/op
BenchmarkUnmarshalByMsgp-128                      	44057836	       135.5 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByProtoBuf-128                  	 9628360	       620.3 ns/op	     176 B/op	      11 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-128              	16364524	       370.0 ns/op	     160 B/op	      10 allocs/op
BenchmarkUnmarshalByThrift-128                    	10060344	       591.0 ns/op	      96 B/op	       6 allocs/op
BenchmarkUnmarshalByThriftIterator-128            	14740503	       407.5 ns/op	     168 B/op	       7 allocs/op
BenchmarkUnmarshalByThriftIteratorDynamic-128     	 5471894	      1098 ns/op	     592 B/op	      18 allocs/op
BenchmarkUnmarshalByThriftIteratorDecoder-128     	 5110819	      1162 ns/op	     616 B/op	      19 allocs/op
BenchmarkUnmarshalByAvro-128                      	  281683	     21232 ns/op	   12305 B/op	     232 allocs/op
BenchmarkUnmarshalByGencode-128                   	59231443	        99.60 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndCbor-128        	 5306976	      1128 ns/op	     656 B/op	       8 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndMsgp-128        	 4681980	      1284 ns/op	     768 B/op	      10 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndBinc-128        	 5368308	      1118 ns/op	     656 B/op	       8 allocs/op
BenchmarkUnmarshalByUgorjiCodecAndJson-128        	 3233894	      1865 ns/op	    1168 B/op	      10 allocs/op
BenchmarkUnmarshalByEasyjson-128                  	16066514	       371.5 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByFfjson-128                    	 3694323	      1620 ns/op	     474 B/op	      13 allocs/op
BenchmarkUnmarshalByJsoniter-128                  	16905810	       354.9 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalBySonic-128                     	24124342	       247.9 ns/op	      99 B/op	       1 allocs/op
BenchmarkUnmarshalByGojay-128                     	 9106419	       659.9 ns/op	     281 B/op	       9 allocs/op
BenchmarkUnmarshalByGoMemdump-128                 	 9233306	       655.1 ns/op	     736 B/op	       9 allocs/op
BenchmarkUnmarshalByColfer-128                    	34182229	       175.7 ns/op	      96 B/op	       6 allocs/op
BenchmarkUnmarshalByZebrapack-128                 	29730520	       195.6 ns/op	      32 B/op	       5 allocs/op
BenchmarkUnmarshalByHprose-128                    	11646049	       517.5 ns/op	     272 B/op	       9 allocs/op
BenchmarkUnmarshalBySereal-128                    	12377004	       480.3 ns/op	      80 B/op	       6 allocs/op
BenchmarkUnmarshalByVmihMsgpackv4-128             	 7308202	       818.5 ns/op	     264 B/op	      11 allocs/op
BenchmarkUnmarshalByRlp-128                       	 9474120	       634.5 ns/op	     104 B/op	      11 allocs/op
BenchmarkUnmarshalBySegmentioJSON-128             	14143380	       421.1 ns/op	      32 B/op	       5 allocs/op
```

**Marshaled Size**

![](size.png)
