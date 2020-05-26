# Benchmarks of Go Protobuf libraries

This is a test suite for benchmarking various Go Protobuf libraries.

## Libraries

* [Go Protobuf](https://blog.golang.org/protobuf-apiv2):
  - [github.com/golang/protobuf](https://github.com/golang/protobuf) v1.3.4 is the most recent pre-APIv2 version of APIv1 (`old` sub-dir).
  - [github.com/golang/protobuf](https://github.com/golang/protobuf) v1.4.0 is a version of APIv1 implemented in terms of APIv2.
  - [google.golang.org/protobuf](https://google.golang.org/protobuf) v1.20.0 is APIv2.

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
BenchmarkGoV1Marshal-4                   |    3318381 |    355 ns/op |    51 |  64 |   1.18 |   17122 |    5.55
BenchmarkGoV1Unmarshal-4                 |    3263350 |    365 ns/op |    51 | 128 |   1.19 |   16838 |    2.85
BenchmarkGoV2Marshal-4                   |    3744885 |    335 ns/op |    51 |  64 |   1.25 |   19361 |    5.23
BenchmarkGoV2Unmarshal-4                 |    2830618 |    407 ns/op |    51 | 128 |   1.15 |   14577 |    3.18
BenchmarkGogoV1Marshal-4                 |    7718380 |    155 ns/op |    51 |  64 |   1.20 |   39826 |    2.42
BenchmarkGogoV1Unmarshal-4               |    5585449 |    210 ns/op |    51 |  96 |   1.17 |   28820 |    2.19
BenchmarkGoV1oldMarshal-4                |    4126480 |    285 ns/op |    51 |  64 |   1.18 |   21292 |    4.45
BenchmarkGoV1oldUnmarshal-4              |    4230336 |    287 ns/op |    51 | 128 |   1.21 |   21828 |    2.24


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGogoV1-4                        |   13303829 |    365 ns/op |   103 | 160 |   4.86 |  137295 |    2.28
BenchmarkGoV1old-4                       |    8356816 |    572 ns/op |   103 | 192 |   4.78 |   86242 |    2.98
BenchmarkGoV1-4                          |    6581731 |    720 ns/op |   103 | 192 |   4.74 |   67923 |    3.75
BenchmarkGoV2-4                          |    6575503 |    742 ns/op |   103 | 192 |   4.88 |   67859 |    3.86

