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
Benchmark_GoV1_Proto_Marshal-4           |    1401818 |    827 ns/op |    51 |  64 |   1.16 |    7233 |   12.92
Benchmark_GoV1_Proto_Unmarshal-4         |    1829254 |    687 ns/op |    51 | 128 |   1.26 |    9438 |    5.37
Benchmark_GoV2_Proto_Marshal-4           |    1510862 |    822 ns/op |    51 |  64 |   1.24 |    7796 |   12.84
Benchmark_GoV2_Proto_Unmarshal-4         |    1735849 |    711 ns/op |    51 | 128 |   1.23 |    8956 |    5.55
Benchmark_GogoV1_Proto_Marshal-4         |    2045078 |    604 ns/op |    51 |  64 |   1.24 |   10552 |    9.44
Benchmark_GogoV1_Proto_Unmarshal-4       |    2538510 |    531 ns/op |    51 |  32 |   1.35 |   13098 |   16.59
Benchmark_GoV1old_Proto_Marshal-4        |    1586356 |    755 ns/op |    51 |  64 |   1.20 |    8185 |   11.80
Benchmark_GoV1old_Proto_Unmarshal-4      |    2228892 |    582 ns/op |    51 | 128 |   1.30 |   11501 |    4.55


benchmark                                | iter      | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-----------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GoV1_JSON_Marshal-4            |     290047 |   4018 ns/op |   126 | 771 |   1.17 |    3654 |    5.21
Benchmark_GoV1_JSON_Unmarshal-4          |      94789 |  12648 ns/op |   126 | 3557 |   1.20 |    1194 |    3.56
Benchmark_GoV2_JSON_Marshal-4            |     331656 |   3564 ns/op |   126 | 582 |   1.18 |    4178 |    6.12
Benchmark_GoV2_JSON_Unmarshal-4          |     282900 |   4630 ns/op |   126 | 621 |   1.31 |    3564 |    7.46
Benchmark_GogoV1_JSON_Marshal-4          |      88027 |  13594 ns/op |   126 | 3405 |   1.20 |    1109 |    3.99
Benchmark_GogoV1_JSON_Unmarshal-4        |      73362 |  16428 ns/op |   126 | 3974 |   1.21 |     924 |    4.13
Benchmark_GoV1old_JSON_Marshal-4         |     100249 |  12146 ns/op |   126 | 2965 |   1.22 |    1263 |    4.10
Benchmark_GoV1old_JSON_Unmarshal-4       |      75784 |  16610 ns/op |   126 | 3973 |   1.26 |     954 |    4.18


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GogoV1_Proto_-4                |    4583588 |   1135 ns/op |   103 |  96 |   5.20 |   47302 |   11.82
Benchmark_GoV1old_Proto_-4               |    3815248 |   1337 ns/op |   103 | 192 |   5.10 |   39373 |    6.96
Benchmark_GoV1_Proto_-4                  |    3231072 |   1514 ns/op |   103 | 192 |   4.89 |   33344 |    7.89
Benchmark_GoV2_Proto_-4                  |    3246711 |   1533 ns/op |   103 | 192 |   4.98 |   33506 |    7.98


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GoV2_JSON_-4                   |     614556 |   8194 ns/op |   252 | 1203 |   5.04 |   15486 |    6.81
Benchmark_GoV1_JSON_-4                   |     384836 |  16666 ns/op |   252 | 4328 |   6.41 |    9697 |    3.85
Benchmark_GoV1old_JSON_-4                |     176033 |  28756 ns/op |   252 | 6938 |   5.06 |    4436 |    4.14
Benchmark_GogoV1_JSON_-4                 |     161389 |  30022 ns/op |   252 | 7379 |   4.85 |    4067 |    4.07
