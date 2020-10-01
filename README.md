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

```proto
message Go {
  string name = 1;
  int64 birthDay = 2;
  string phone = 3;
  int32 siblings = 4;
  bool spouse = 5;
  double money = 6;
  Type type = 7;
  oneof values {
    string value_s = 8;
    int32 value_i = 9;
    double value_d = 10;
  }
}

enum Type {
  TYPE_UNSPECIFIED = 0;
  TYPE_R = 1;
  TYPE_S = 2;
}
```


## Results

benchmark                                | iter      | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-----------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GogoV1_JSON_Marshal-4          |      55704 |  21345 ns/op |   156 | 4561 |   1.19 |     868 |    4.68
Benchmark_GogoV1_JSON_Unmarshal-4        |      52802 |  21064 ns/op |   156 | 4603 |   1.11 |     823 |    4.58
Benchmark_GogoV1_Proto_Marshal-4         |    1481065 |    861 ns/op |    60 |  64 |   1.28 |    8901 |   13.45
Benchmark_GogoV1_Proto_Unmarshal-4       |    2412230 |    568 ns/op |    60 |  48 |   1.37 |   14497 |   11.83
Benchmark_GoV1_JSON_Marshal-4            |     220665 |   6542 ns/op |   154 | 1003 |   1.44 |    3398 |    6.52
Benchmark_GoV1_JSON_Unmarshal-4          |      70056 |  16583 ns/op |   154 | 4104 |   1.16 |    1078 |    4.04
Benchmark_GoV1old_JSON_Marshal-4         |      69189 |  17325 ns/op |   154 | 4016 |   1.20 |    1065 |    4.31
Benchmark_GoV1old_JSON_Unmarshal-4       |      58640 |  21309 ns/op |   154 | 4680 |   1.25 |     903 |    4.55
Benchmark_GoV1old_Proto_Marshal-4        |    1000000 |   1319 ns/op |    60 |  72 |   1.32 |    6010 |   18.32
Benchmark_GoV1old_Proto_Unmarshal-4      |    1408716 |   1117 ns/op |    60 | 160 |   1.57 |    8466 |    6.98
Benchmark_GoV1_Proto_Marshal-4           |    1000000 |   1430 ns/op |    60 |  64 |   1.43 |    6010 |   22.34
Benchmark_GoV1_Proto_Unmarshal-4         |    1305186 |   1023 ns/op |    60 | 176 |   1.34 |    7844 |    5.81
Benchmark_GoV2_JSON_Marshal-4            |     186824 |   5441 ns/op |   154 | 761 |   1.02 |    2877 |    7.15
Benchmark_GoV2_JSON_Unmarshal-4          |     213820 |   5817 ns/op |   154 | 741 |   1.24 |    3292 |    7.85
Benchmark_GoV2_Proto_Marshal-4           |     946030 |   1231 ns/op |    60 |  64 |   1.16 |    5685 |   19.23
Benchmark_GoV2_Proto_Unmarshal-4         |    1000000 |   1016 ns/op |    60 | 176 |   1.02 |    6010 |    5.77


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GogoV1_Proto_-4                |    3893295 |   1429 ns/op |   120 | 112 |   5.56 |   46797 |   12.76
Benchmark_GoV2_Proto_-4                  |    1946030 |   2247 ns/op |   120 | 240 |   4.37 |   23391 |    9.36
Benchmark_GoV1old_Proto_-4               |    2408716 |   2436 ns/op |   120 | 232 |   5.87 |   28952 |   10.50
Benchmark_GoV1_Proto_-4                  |    2305186 |   2453 ns/op |   120 | 240 |   5.65 |   27708 |   10.22
Benchmark_GoV2_JSON_-4                   |     400644 |  11258 ns/op |   308 | 1502 |   4.51 |   12339 |    7.50
Benchmark_GoV1_JSON_-4                   |     290721 |  23125 ns/op |   308 | 5107 |   6.72 |    8954 |    4.53
Benchmark_GoV1old_JSON_-4                |     127829 |  38634 ns/op |   308 | 8696 |   4.94 |    3937 |    4.44
Benchmark_GogoV1_JSON_-4                 |     108506 |  42409 ns/op |   312 | 9164 |   4.60 |    3385 |    4.63
