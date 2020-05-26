test:
	rm results.txt
	go test -count=1 -bench=. | grep BenchmarkGo > results.txt
	cd old && go test -count=1 -bench=. | grep BenchmarkGo >> ../results.txt
	cat results.txt

all: old/structdef-go-v1.pb.go structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go

old/structdef-go-v1.pb.go: structdef-go-v1.proto
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.5
	protoc --go_out=paths=source_relative:./old structdef-go-v1.proto

structdef-gogo-v1.pb.go: structdef-gogo-v1.proto
	go get github.com/gogo/protobuf/protoc-gen-gogofaster
	protoc --gogofaster_out=paths=source_relative:. structdef-gogo-v1.proto

structdef-go-v1.pb.go: structdef-go-v1.proto
	go get github.com/golang/protobuf/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. structdef-go-v1.proto

structdef-go-v2.pb.go: structdef-go-v2.proto
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. structdef-go-v2.proto

clean:
	rm -f *.pb.go
	rm -f old/*.pb.go
