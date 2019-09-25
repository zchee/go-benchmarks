package hashbench

import (
	"bytes/hash"
	"flag"
	"fmt"
	"os"
	"runtime"
	"testing"

	xxhash "github.com/cespare/xxhash/v2"
	"github.com/zeebo/xxh3"
)

var (
	cespare   bool
	zeebo     bool
	byteshash bool
)

func TestMain(m *testing.M) {
	var status int
	defer func() { os.Exit(status) }()

	flag.BoolVar(&cespare, "cespare", false, "benchmark the github.com/cespare/xxhash/v2")
	flag.BoolVar(&zeebo, "zeebo", false, "benchmark the github.com/zeebo/xxh3")
	flag.BoolVar(&byteshash, "byteshash", false, "benchmark the bytes/hash")

	status = m.Run()
}

func BenchmarkHash(b *testing.B) {
	sizes := []int{
		0, 1, 3, 4, 8, 9, 16, 17, 32,
		33, 64, 65, 96, 97, 128, 129, 240, 241,
		512, 1024, 100 * 1024,
	}

	if cespare {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
				b.SetBytes(int64(size))
				var acc uint64
				d := string(make([]byte, size))
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					acc = xxhash.Sum64String(d)
				}
				runtime.KeepAlive(acc)
			})
		}
	}

	if zeebo {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
				b.SetBytes(int64(size))
				var acc uint64
				d := string(make([]byte, size))
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					acc = xxh3.HashString(d)
				}
				runtime.KeepAlive(acc)
			})
		}
	}

	if byteshash {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
				b.SetBytes(int64(size))
				var acc uint64
				d := string(make([]byte, size))
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					h := hash.New()
					h.AddString(d)
					acc = h.Hash()
				}
				runtime.KeepAlive(acc)
			})
		}
	}
}
