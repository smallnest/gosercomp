#/bin/sh

#
go install github.com/tinylib/msgp
$GOPATH/bin/msgp -o msgp_gen.go -io=false -tests=false -file data.go

# https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.1
# brew install protoc
protoc --go_out=. protobuf.proto
protoc --gogofaster_out=.  -I. -I$GOPATH/src  mygogo.proto

# https://thrift-tutorial.readthedocs.io/en/latest/installation.html
# brew install thrift
thrift -r -out ./.. --gen go thrift_colorgroup.thrift

# https://github.com/andyleap/gencode
go install github.com/andyleap/gencode
$GOPATH/bin/gencode go -schema=gencode.schema -package model

# https://github.com/ugorji/go/tree/master/codec/codecgen
go install github.com/ugorji/go/codec/codecgen
$GOPATH/bin/codecgen -o data_codec.go -r ColorGroup

# https://github.com/glycerine/zebrapack
go install github.com/glycerine/zebrapack
$GOPATH/bin/zebrapack -fast-strings -no-load -no-embedded-schema -file zebrapack_data.go

# https://github.com/pquerna/ffjson
go install github.com/pquerna/ffjson
mkdir temp_ffjson
cp ffjson_data.go temp_ffjson
cd temp_ffjson
$GOPATH/bin/ffjson ffjson_data.go
cp -f * ..
cd ..
rm -fr temp_ffjson

# https://github.com/mailru/easyjson
go install github.com/mailru/easyjson/easyjson
mkdir temp_easyjson
cp easyjson_data.go temp_easyjson
cd temp_easyjson
$GOPATH/bin/easyjson -all easyjson_data.go
cp -f * ..
cd ..
rm -fr temp_easyjson


# https://github.com/actgardner/gogen-avro
go install github.com/actgardner/gogen-avro/gogen-avro
$GOPATH/bin/gogen-avro -package model . data.avsc 

# https://github.com/pascaldekloe/colfer
go install github.com/pascaldekloe/colfer/cmd/colf
$GOPATH/bin/colf Go colorgroup.colf