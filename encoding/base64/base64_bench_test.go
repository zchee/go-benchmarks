package base64_benchmark

import (
	"encoding/base64"
	"flag"
	"fmt"
	"testing"

	base64_chenzhuoyu "github.com/chenzhuoyu/base64x"
	base64_cristalhq "github.com/cristalhq/base64"
	base64_segmentio "github.com/segmentio/asm/base64"
)

var (
	std        bool
	chenzhuoyu bool
	cristalhq  bool
	segmentio  bool
	tmthrgd    bool
)

func TestMain(m *testing.M) {
	flag.BoolVar(&std, "std", false, "run stdlib base64 benchmark")
	flag.BoolVar(&chenzhuoyu, "chenzhuoyu", false, "run chenzhuoyu/base64x benchmark")
	flag.BoolVar(&cristalhq, "cristalhq", false, "run cristalhq/base64 benchmark")
	flag.BoolVar(&segmentio, "segmentio", false, "run segmentio/asm/base64 benchmark")
	flag.Parse()

	m.Run()
}

func BenchmarkEncodeToString(b *testing.B) {
	var fn func(src []byte) string
	switch {
	case std:
		fn = base64.StdEncoding.EncodeToString
	case chenzhuoyu:
		fn = base64_chenzhuoyu.StdEncoding.EncodeToString
	case cristalhq:
		fn = base64_cristalhq.StdEncoding.EncodeToString
	case segmentio:
		fn = base64_segmentio.StdEncoding.EncodeToString
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
		case chenzhuoyu:
			fn = base64_chenzhuoyu.StdEncoding.DecodeString
		case cristalhq:
			fn = base64_cristalhq.StdEncoding.DecodeString
		case segmentio:
			fn = base64_segmentio.StdEncoding.DecodeString
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
