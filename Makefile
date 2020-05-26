ifndef GOPATH
GOPATH := $(shell go env GOPATH)
endif

all: Colfer.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go gencode.schema.gen.go gencode-unsafe.schema.gen.go structdefxdr_generated.go

Colfer.go:
	go run github.com/pascaldekloe/colfer/cmd/colf go
	mv goserbench/Colfer.go .
	rmdir goserbench

msgp_gen.go: structdef.go
	go run github.com/tinylib/msgp -o msgp_gen.go -file structdef.go -io=false -tests=false

structdef_easyjson.go: structdef.go
	go run github.com/mailru/easyjson/easyjson -all structdef.go

structdef-gogo.pb.go: structdef-gogo.proto
	protoc --gogofaster_out=paths=source_relative:. -I. -I${GOPATH}/src  -I${GOPATH}/src/github.com/gogo/protobuf/protobuf structdef-gogo.proto

structdef.pb.go: structdef.proto
	protoc --go_out=paths=source_relative:. structdef.proto

gencode.schema.gen.go: gencode.schema
	go run github.com/andyleap/gencode go -schema=gencode.schema -package=goserbench

gencode-unsafe.schema.gen.go: gencode-unsafe.schema
	go run github.com/andyleap/gencode go -schema=gencode-unsafe.schema -package=goserbench -unsafe

structdefxdr_generated.go: structdefxdr.go
	go run github.com/calmh/xdr/cmd/genxdr -o structdefxdr_generated.go structdefxdr.go

.PHONY: clean
clean:
	rm -f Colfer.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go gencode.schema.gen.go gencode-unsafe.schema.gen.go structdefxdr_generated.go

.PHONY: install
install:
	go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/tinylib/msgp
	go get -u github.com/andyleap/gencode
	go get -u github.com/mailru/easyjson/...
	go get -u go.dedis.ch/protobuf
	go get -u github.com/Sereal/Sereal/Go/sereal
	go get -u github.com/alecthomas/binary
	go get -u github.com/davecgh/go-xdr/xdr
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/google/flatbuffers/go
	go get -u github.com/tinylib/msgp/msgp
	go get -u github.com/ugorji/go/codec
	go get -u gopkg.in/mgo.v2/bson
	go get -u github.com/vmihailenco/msgpack/v4
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/hprose/hprose-go/io
	go get -u github.com/pascaldekloe/colfer/cmd/colf
	go get -u github.com/calmh/xdr
	go get -u github.com/niubaoshu/gotiny
