package protobench

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	proto1jsonpb "github.com/golang/protobuf/jsonpb"
	proto1 "github.com/golang/protobuf/proto"
	proto2json "google.golang.org/protobuf/encoding/protojson"
	proto2 "google.golang.org/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

// github.com/golang/protobuf (aka V1)

func generateGoV1(n int) []*GoV1 {
	a := make([]*GoV1, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, &GoV1{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     TypeV1(rand.Intn(4)),
			Values:   &GoV1_ValueS{ValueS: randString(5)},
		})
	}
	return a
}

func Benchmark_GoV1_Proto_Marshal(b *testing.B) {
	data := generateGoV1(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto1.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoV1_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoV1(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto1.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GoV1{}
		err := proto1.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_GoV1_JSON_Marshal(b *testing.B) {
	data := generateGoV1(b.N)
	marshaler := proto1jsonpb.Marshaler{}
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		err := marshaler.Marshal(&buf, data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoV1_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoV1(b.N)
	marshaler := proto1jsonpb.Marshaler{}
	ser := make([]bytes.Buffer, len(data))
	var serialSize int
	for i, d := range data {
		err := marshaler.Marshal(&ser[i], d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += ser[i].Len()
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	unmarshaler := proto1jsonpb.Unmarshaler{}
	randomI := randomI(b.N)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := randomI[i]
		o := &GoV1{}
		err := unmarshaler.Unmarshal(&ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n].String())
		}
	}

}

// google.golang.org/protobuf (aka V2)

func generateGoV2(n int) []*GoV2 {
	a := make([]*GoV2, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, &GoV2{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     TypeV2(rand.Intn(4)),
			Values:   &GoV2_ValueS{ValueS: randString(5)},
		})
	}
	return a
}

func Benchmark_GoV2_Proto_Marshal(b *testing.B) {
	data := generateGoV2(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto2.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoV2_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoV2(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto2.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GoV2{}
		err := proto2.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_GoV2_JSON_Marshal(b *testing.B) {
	data := generateGoV2(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto2json.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoV2_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoV2(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto2json.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	randomI := randomI(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := randomI[i]
		o := &GoV2{}
		err := proto2json.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

// github.com/gogo/protobuf/proto (aka gogo)

func generateGogoV1(n int) []*GogoV1 {
	a := make([]*GogoV1, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, &GogoV1{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     TypeGoGoV1(rand.Intn(4)),
			Values:   &GogoV1_ValueS{ValueS: randString(5)},
		})
	}
	return a
}

func Benchmark_GogoV1_Proto_Marshal(b *testing.B) {
	data := generateGogoV1(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].Marshal()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GogoV1_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoV1(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = d.Marshal()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GogoV1{}
		err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("gogoprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_GogoV1_JSON_Marshal(b *testing.B) {
	data := generateGogoV1(b.N)
	marshaler := gogojsonpb.Marshaler{}
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		err := marshaler.Marshal(&buf, data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GogoV1_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoV1(b.N)
	marshaler := gogojsonpb.Marshaler{}
	ser := make([]bytes.Buffer, len(data))
	var serialSize int
	for i, d := range data {
		err := marshaler.Marshal(&ser[i], d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += ser[i].Len()
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	unmarshaler := gogojsonpb.Unmarshaler{}
	randomI := randomI(b.N)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := randomI[i]
		o := &GogoV1{}
		err := unmarshaler.Unmarshal(&ser[n], o)
		if err != nil {
			b.Fatalf("gogoprotobuf failed to unmarshal: %s (%s)", err, ser[n].String())
		}
	}
}

// github.com/planetscale/vtprotobuf (aka vtproto)

func generateVTProto(n int) []*VTProto {
	a := make([]*VTProto, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, &VTProto{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     TypeVTProtoV2(rand.Intn(4)),
			Values:   &VTProto_ValueS{ValueS: randString(5)},
		})
	}
	return a
}

func Benchmark_VTProto_Proto_Marshal(b *testing.B) {
	data := generateVTProto(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].MarshalVT()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_VTProto_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateVTProto(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = d.MarshalVT()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &VTProto{}
		err := o.UnmarshalVT(ser[n])
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_VTProto_JSON_Marshal(b *testing.B) {
	data := generateVTProto(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto2json.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_VTProto_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateVTProto(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto2json.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	randomI := randomI(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := randomI[i]
		o := &VTProto{}
		err := proto2json.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func randomI(n int) []int {
	randomI := make([]int, n)
	for i := 0; i < len(randomI); i++ {
		randomI[i] = i
	}
	rand.Shuffle(len(randomI), func(i, j int) {
		randomI[i], randomI[j] = randomI[j], randomI[i]
	})
	return randomI
}
