package protobench

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	gogoproto "github.com/gogo/protobuf/proto"
	proto1 "github.com/golang/protobuf/proto"
	proto2 "google.golang.org/protobuf/proto"
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

// github.com/golang/protobuf

func generateGoV1() []*GoV1 {
	a := make([]*GoV1, 0, 1000)
	for i := 0; i < 1000; i++ {
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
	data := generateGoV1()
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
	data := generateGoV1()
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

// google.golang.org/protobuf

func generateGoV2() []*GoV2 {
	a := make([]*GoV2, 0, 1000)
	for i := 0; i < 1000; i++ {
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
	data := generateGoV2()
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
	data := generateGoV2()
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

// github.com/gogo/protobuf/proto

func generateGogoV1() []*GogoV1 {
	a := make([]*GogoV1, 0, 1000)
	for i := 0; i < 1000; i++ {
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
	data := generateGogoV1()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := gogoproto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func BenchmarkGogoV1Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoV1()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = gogoproto.Marshal(d)
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
		err := gogoproto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}
