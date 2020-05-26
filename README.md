# Benchmarks of Go Protobuf libraries

This is a test suite for benchmarking various Go Protobuf libraries.

## Libraries

* [Go Protobuf](https://blog.golang.org/protobuf-apiv2):
  - [github.com/golang/protobuf](https://github.com/golang/protobuf/releases/tag/v1.3.5) ~v1.3.5 is the most recent pre-APIv2 version of APIv1 (`old` sub-dir).
  - [github.com/golang/protobuf](https://github.com/golang/protobuf) ^v1.4.0 is a version of APIv1 implemented in terms of APIv2.
  - [google.golang.org/protobuf](https://github.com/protocolbuffers/protobuf-go) ^v1.20.0 is APIv2.

* [Gogo Protobuf](https://github.com/gogo/protobuf):
  - [github.com/gogo/protobuf](https://github.com/gogo/protobuf) Protocol Buffers for Go with Gadgets.

## Running the benchmarks

```bash
make all
make test
```

To update the table in the README:

```bash
./stats.sh
```

## Data

The data being serialized is the following structure with randomly generated values:

```go
type Go struct {
    Name     string
    BirthDay time.Time
    Phone    string
    Siblings int
    Spouse   bool
    Money    float64
}
```


## Results

benchmark                                | iter      | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-----------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGoV1Marshal-4                   |    3399463 |    348 ns/op |    51 |  64 |   1.18 |   17541 |    5.44
BenchmarkGoV1Unmarshal-4                 |    3307972 |    362 ns/op |    51 | 128 |   1.20 |   17069 |    2.83
BenchmarkGoV2Marshal-4                   |    3486094 |    318 ns/op |    51 |  64 |   1.11 |   17988 |    4.97
BenchmarkGoV2Unmarshal-4                 |    2589584 |    428 ns/op |    51 | 128 |   1.11 |   13362 |    3.34
BenchmarkGogoV1Marshal-4                 |    7178234 |    162 ns/op |    51 |  64 |   1.16 |   37039 |    2.53
BenchmarkGogoV1Unmarshal-4               |    5536908 |    211 ns/op |    51 |  96 |   1.17 |   28625 |    2.20
BenchmarkGoV1oldMarshal-4                |    4081132 |    283 ns/op |    51 |  64 |   1.15 |   21058 |    4.42
BenchmarkGoV1oldUnmarshal-4              |    3897282 |    290 ns/op |    51 | 128 |   1.13 |   20109 |    2.27


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGogoV1-4                        |   12715142 |    373 ns/op |   103 | 160 |   4.74 |  131347 |    2.33
BenchmarkGoV1old-4                       |    7978414 |    573 ns/op |   103 | 192 |   4.57 |   82337 |    2.98
BenchmarkGoV1-4                          |    6707435 |    710 ns/op |   103 | 192 |   4.76 |   69220 |    3.70
BenchmarkGoV2-4                          |    6075678 |    746 ns/op |   103 | 192 |   4.53 |   62700 |    3.89

