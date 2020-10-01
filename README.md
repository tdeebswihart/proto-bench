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
Benchmark_JSON_GogoV1_Marshal-4          |      62697 |  18257 ns/op |   156 | 4562 |   1.14 |     978 |    4.00
Benchmark_JSON_GogoV1_Unmarshal-4        |      58150 |  20696 ns/op |   156 | 4603 |   1.20 |     907 |    4.50
Benchmark_JSON_GoV1_Marshal-4            |     220243 |   5507 ns/op |   154 | 1002 |   1.21 |    3391 |    5.50
Benchmark_JSON_GoV1old_Marshal-4         |      67186 |  16358 ns/op |   154 | 4016 |   1.10 |    1034 |    4.07
Benchmark_JSON_GoV1old_Unmarshal-4       |      58538 |  19958 ns/op |   154 | 4681 |   1.17 |     901 |    4.26
Benchmark_JSON_GoV1_Unmarshal-4          |      70801 |  15931 ns/op |   154 | 4104 |   1.13 |    1090 |    3.88
Benchmark_JSON_GoV2_Marshal-4            |     255117 |   4990 ns/op |   154 | 761 |   1.27 |    3928 |    6.56
Benchmark_JSON_GoV2_Unmarshal-4          |     210552 |   5490 ns/op |   154 | 741 |   1.16 |    3242 |    7.41
Benchmark_Proto_GogoV1_Marshal-4         |    1550792 |    816 ns/op |    60 |  64 |   1.27 |    9320 |   12.75
Benchmark_Proto_GogoV1_Unmarshal-4       |    2682651 |    478 ns/op |    60 |  48 |   1.28 |   16122 |    9.96
Benchmark_Proto_GoV1_Marshal-4           |    1000000 |   1287 ns/op |    60 |  64 |   1.29 |    6010 |   20.11
Benchmark_Proto_GoV1old_Marshal-4        |    1000000 |   1227 ns/op |    60 |  72 |   1.23 |    6010 |   17.04
Benchmark_Proto_GoV1old_Unmarshal-4      |    1445637 |    868 ns/op |    60 | 160 |   1.25 |    8688 |    5.42
Benchmark_Proto_GoV1_Unmarshal-4         |    1226025 |    997 ns/op |    60 | 176 |   1.22 |    7368 |    5.66
Benchmark_Proto_GoV2_Marshal-4           |    1000000 |   1225 ns/op |    60 |  64 |   1.23 |    6010 |   19.14
Benchmark_Proto_GoV2_Unmarshal-4         |    1272902 |    964 ns/op |    60 | 176 |   1.23 |    7650 |    5.48


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_Proto_GogoV1_-4                |    4233443 |   1294 ns/op |   120 | 112 |   5.48 |   50885 |   11.55
Benchmark_Proto_GoV1old_-4               |    2445637 |   2095 ns/op |   120 | 232 |   5.12 |   29396 |    9.03
Benchmark_Proto_GoV2_-4                  |    2272902 |   2189 ns/op |   120 | 240 |   4.98 |   27320 |    9.12
Benchmark_Proto_GoV1_-4                  |    2226025 |   2284 ns/op |   120 | 240 |   5.08 |   26756 |    9.52
Benchmark_JSON_GoV2_-4                   |     465669 |  10480 ns/op |   308 | 1502 |   4.88 |   14342 |    6.98
Benchmark_JSON_GoV1_-4                   |     291044 |  21438 ns/op |   308 | 5106 |   6.24 |    8964 |    4.20
Benchmark_JSON_GoV1old_-4                |     125724 |  36316 ns/op |   308 | 8697 |   4.57 |    3872 |    4.18
Benchmark_JSON_GogoV1_-4                 |     120847 |  38953 ns/op |   312 | 9165 |   4.71 |    3770 |    4.25
