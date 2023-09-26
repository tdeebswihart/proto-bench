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
Benchmark_GoV1_Proto_Marshal-12          |    2426719 |    554 ns/op |    60 |  64 |   1.35 |   14584 |    8.67
Benchmark_GoV1_Proto_Unmarshal-12        |    3167671 |    433 ns/op |    60 | 176 |   1.37 |   19037 |    2.46
Benchmark_GoV1_JSON_Marshal-12           |     643039 |   1887 ns/op |   153 | 972 |   1.21 |    9896 |    1.94
Benchmark_GoV1_JSON_Unmarshal-12         |     199536 |   6225 ns/op |   153 | 3524 |   1.24 |    3070 |    1.77
Benchmark_GoV2_Proto_Marshal-12          |    2605926 |    524 ns/op |    60 |  64 |   1.37 |   15661 |    8.20
Benchmark_GoV2_Proto_Unmarshal-12        |    2826908 |    443 ns/op |    60 | 176 |   1.25 |   16989 |    2.52
Benchmark_GoV2_JSON_Marshal-12           |     755736 |   1505 ns/op |   154 | 913 |   1.14 |   11638 |    1.65
Benchmark_GoV2_JSON_Unmarshal-12         |     603441 |   2119 ns/op |   153 | 617 |   1.28 |    9286 |    3.43
Benchmark_GogoV1_Proto_Marshal-12        |    3661051 |    403 ns/op |    60 |  64 |   1.48 |   22002 |    6.31
Benchmark_GogoV1_Proto_Unmarshal-12      |    5246887 |    313 ns/op |    60 |  48 |   1.64 |   31533 |    6.53
Benchmark_GogoV1_JSON_Marshal-12         |     154326 |   7135 ns/op |   155 | 4457 |   1.10 |    2404 |    1.60
Benchmark_GogoV1_JSON_Unmarshal-12       |     146514 |   8054 ns/op |   155 | 4024 |   1.18 |    2284 |    2.00
Benchmark_VTProto_Proto_Marshal-12       |    3456073 |    414 ns/op |    60 |  64 |   1.43 |   20770 |    6.47
Benchmark_VTProto_Proto_Unmarshal-12     |    4258104 |    310 ns/op |    60 |  48 |   1.32 |   25591 |    6.46
Benchmark_VTProto_JSON_Marshal-12        |     760351 |   1566 ns/op |   157 | 913 |   1.19 |   11967 |    1.72
Benchmark_VTProto_JSON_Unmarshal-12      |     573082 |   2248 ns/op |   157 | 617 |   1.29 |    9020 |    3.64


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GogoV1_Proto_-12               |    8907938 |    717 ns/op |   120 | 112 |   6.39 |  107073 |    6.40
Benchmark_VTProto_Proto_-12              |    7714177 |    724 ns/op |   120 | 112 |   5.59 |   92724 |    6.47
Benchmark_GoV2_Proto_-12                 |    5432834 |    968 ns/op |   120 | 240 |   5.26 |   65302 |    4.04
Benchmark_GoV1_Proto_-12                 |    5594390 |    988 ns/op |   120 | 240 |   5.53 |   67244 |    4.12
Benchmark_GoV2_JSON_-12                  |    1359177 |   3624 ns/op |   307 | 1530 |   4.93 |   41849 |    2.37
Benchmark_VTProto_JSON_-12               |    1333433 |   3814 ns/op |   314 | 1530 |   5.09 |   41976 |    2.49
Benchmark_GoV1_JSON_-12                  |     842575 |   8112 ns/op |   307 | 4496 |   6.83 |   25934 |    1.80
Benchmark_GogoV1_JSON_-12                |     300840 |  15189 ns/op |   311 | 8481 |   4.57 |    9377 |    1.79
