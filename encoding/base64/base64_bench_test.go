package base64_benchmark

import (
	"encoding/base64"
	"flag"
	"fmt"
	"testing"

	base64_simd "github.com/zchee/go-benchmarks/encoding/base64/go-base64"
)

var (
	std  bool
	simd bool
)

func TestMain(m *testing.M) {
	flag.BoolVar(&std, "std", false, "run stdlib base64 benchmark")
	flag.BoolVar(&simd, "simd", false, "run tmthrgd/go-base64 benchmark")
	flag.Parse()

	m.Run()
}

func BenchmarkEncodeToString(b *testing.B) {
	var fn func(src []byte) string
	switch {
	case std:
		fn = base64.StdEncoding.EncodeToString
	case simd:
		fn = base64_simd.StdEncoding.EncodeToString
	}

	data := make([]byte, 8192)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(data)
	}
}

func BenchmarkDecodeString(b *testing.B) {
	sizes := []int{2, 4, 8, 64, 8192}
	benchFunc := func(b *testing.B, benchSize int) {
		var fn func(s string) ([]byte, error)
		switch {
		case std:
			fn = base64.StdEncoding.DecodeString
		case simd:
			fn = base64_simd.StdEncoding.DecodeString
		}

		data := base64.StdEncoding.EncodeToString(make([]byte, benchSize))
		b.SetBytes(int64(len(data)))
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fn(data)
		}
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			benchFunc(b, size)
		})
	}
}
