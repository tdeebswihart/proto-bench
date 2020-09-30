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
BenchmarkGoV1Marshal-4                   |    1497213 |    827 ns/op |    51 |  64 |   1.24 |    7725 |   12.92
BenchmarkGoV1Unmarshal-4                 |    1820446 |    662 ns/op |    51 | 128 |   1.21 |    9393 |    5.17
BenchmarkGoV1MarshalJSON-4               |     309465 |   3988 ns/op |   126 | 772 |   1.23 |    3899 |    5.17
BenchmarkGoV1UnmarshalJSON-4             |      88539 |  12743 ns/op |   126 | 3557 |   1.13 |    1115 |    3.58
BenchmarkGoV2Marshal-4                   |    1552492 |    807 ns/op |    51 |  64 |   1.25 |    8010 |   12.61
BenchmarkGoV2Unmarshal-4                 |    1785612 |    676 ns/op |    51 | 128 |   1.21 |    9213 |    5.28
BenchmarkGoV2MarshalJSON-4               |     354578 |   3561 ns/op |   126 | 582 |   1.26 |    4467 |    6.12
BenchmarkGoV2UnmarshalJSON-4             |     270232 |   4377 ns/op |   126 | 621 |   1.18 |    3404 |    7.05
BenchmarkGogoV1Marshal-4                 |    2105110 |    604 ns/op |    51 |  64 |   1.27 |   10862 |    9.44
BenchmarkGogoV1Unmarshal-4               |    2621703 |    503 ns/op |    51 |  32 |   1.32 |   13527 |   15.72
BenchmarkGogoV1MarshalJSON-4             |      90421 |  13376 ns/op |   126 | 3406 |   1.21 |    1139 |    3.93
BenchmarkGogoV1UnmarshalJSON-4           |      71382 |  16188 ns/op |   126 | 3973 |   1.16 |     899 |    4.07
BenchmarkGoV1oldMarshal-4                |    1611070 |    776 ns/op |    51 |  64 |   1.25 |    8313 |   12.12
BenchmarkGoV1oldUnmarshal-4              |    2087022 |    592 ns/op |    51 | 128 |   1.24 |   10769 |    4.62
BenchmarkGoV1MarshalJSON-4               |      99934 |  11786 ns/op |   126 | 2962 |   1.18 |    1259 |    3.98
BenchmarkGoV1UnmarshalJSON-4             |      77542 |  15425 ns/op |   126 | 3973 |   1.20 |     977 |    3.88


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGogoV1-4                        |    4726813 |   1107 ns/op |   103 |  96 |   5.23 |   48780 |   11.53
BenchmarkGoV1old-4                       |    3698092 |   1368 ns/op |   103 | 192 |   5.06 |   38164 |    7.12
BenchmarkGoV2-4                          |    3338104 |   1483 ns/op |   103 | 192 |   4.95 |   34449 |    7.72
BenchmarkGoV1-4                          |    3317659 |   1489 ns/op |   103 | 192 |   4.94 |   34238 |    7.76
BenchmarkGoV2JSON-4                      |     624810 |   7938 ns/op |   252 | 1203 |   4.96 |   15745 |    6.60
BenchmarkGogoV1JSON-4                    |     161803 |  29564 ns/op |   252 | 7379 |   4.78 |    4077 |    4.01
BenchmarkGoV1JSON-4                      |     575480 |  43942 ns/op |   504 | 11264 |  25.29 |   29004 |    3.90
