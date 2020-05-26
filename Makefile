test:
	go test -count=1 -bench=. | tee results.txt
	cd old && go test -count=1 -bench=. | tee -a ../results.txt

all: old/structdef-go-v1.pb.go structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go

old/structdef-go-v1.pb.go: structdef-go-v1.proto
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.4
	protoc --go_out=paths=source_relative:./old structdef-go-v1.proto

structdef-gogo-v1.pb.go: structdef-gogo-v1.proto
	go install github.com/gogo/protobuf/protoc-gen-gogofaster
	protoc --gogofaster_out=paths=source_relative:. structdef-gogo-v1.proto

structdef-go-v1.pb.go: structdef-go-v1.proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. structdef-go-v1.proto

structdef-go-v2.pb.go: structdef-go-v2.proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc --go_out=paths=source_relative:. structdef-go-v2.proto

clean:
	rm -f *.pb.go
	rm -f old/*.pb.go
