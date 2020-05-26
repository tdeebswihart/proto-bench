test:
	go test -bench=. | tee -a results.txt

all: structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go

structdef-gogo-v1.pb.go: structdef-gogo-v1.proto
	go install github.com/gogo/protobuf/protoc-gen-gogofaster
	protoc --gogofaster_out=paths=source_relative:. structdef-gogo-v1.proto

structdef-go-v1.pb.go: structdef-go-v1.proto
	go install github.com/golang/protobuf/protoc-gen-go
	protoc --go_out=paths=source_relative:. structdef-go-v1.proto

structdef-go-v2.pb.go: structdef-go-v2.proto
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	protoc --go_out=paths=source_relative:. structdef-go-v2.proto

clean:
	rm -f *.pb.go
