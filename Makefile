all: old/structdef-go-v1.pb.go structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go structdef-vtproto-v2.pb.go

test: all
	rm results.txt || true
	go test -count=1 -bench=. | grep Benchmark_ > results.txt
	cd old && go test -count=1 -bench=. | grep Benchmark_ >> ../results.txt
	sort -o results.txt results.txt
	cat results.txt

old/structdef-go-v1.pb.go: structdef-go-v1.proto
	go get github.com/golang/protobuf/protoc-gen-go@1.3.5
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

structdef-vtproto-v2.pb.go: structdef-vtproto-v2.proto
	# note: running `go get` here didn't work
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest
	protoc --go-vtproto_out=paths=source_relative:. --go_out=paths=source_relative:. $<

clean:
	rm -f *.pb.go
	rm -f old/*.pb.go
