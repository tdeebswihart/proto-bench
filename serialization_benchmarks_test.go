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

func init(){
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
		})
	}
	return a
}

func BenchmarkGoV1Marshal(b *testing.B) {
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

func BenchmarkGoV1Unmarshal(b *testing.B) {
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

func BenchmarkGoV1MarshalJSON(b *testing.B) {
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

func BenchmarkGoV1UnmarshalJSON(b *testing.B) {
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
		})
	}
	return a
}

func BenchmarkGoV2Marshal(b *testing.B) {
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

func BenchmarkGoV2Unmarshal(b *testing.B) {
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

func BenchmarkGoV2MarshalJSON(b *testing.B) {
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

func BenchmarkGoV2UnmarshalJSON(b *testing.B) {
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
		})
	}
	return a
}

func BenchmarkGogoV1Marshal(b *testing.B) {
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

func BenchmarkGogoV1Unmarshal(b *testing.B) {
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


func BenchmarkGogoV1MarshalJSON(b *testing.B) {
	data := generateGoV1(b.N)
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

func BenchmarkGogoV1UnmarshalJSON(b *testing.B) {
	b.StopTimer()
	data := generateGoV1(b.N)
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
		o := &GoV1{}
		err := unmarshaler.Unmarshal(&ser[n], o)
		if err != nil {
			b.Fatalf("gogoprotobuf failed to unmarshal: %s (%s)", err, ser[n].String())
		}
	}
}

func randomI(n int)[]int{
	randomI := make([]int, n)
	for i := 0; i < len(randomI); i++ {
		randomI[i] = i
	}
	rand.Shuffle(len(randomI), func(i, j int) {
		randomI[i], randomI[j] = randomI[j], randomI[i]
	})
	return randomI
}
