package hashbench

import (
	"flag"
	"hash/maphash"
	"strconv"
	"testing"
	"unsafe"

	bdxxhash3 "github.com/bytedance/gopkg/util/xxhash3"
	xxhash "github.com/cespare/xxhash/v2"
	simdxxhash "github.com/go-simd/xxhash"
	"github.com/zeebo/xxh3"
)

var (
	cespare   = flag.Bool("cespare", false, "benchmark the github.com/cespare/xxhash/v2")
	zeebo     = flag.Bool("zeebo", false, "benchmark the github.com/zeebo/xxh3")
	byteshash = flag.Bool("maphash", false, "benchmark the hash/maphash")
	bytedance = flag.Bool("bytedance", false, "benchmark the github.com/bytedance/gopkg/util/xxhash3")
	gosimd    = flag.Bool("gosimd", false, "benchmark the github.com/go-simd/xxhash")
)

func BenchmarkHash(b *testing.B) {
	sizes := []int{
		0, 1, 3, 4, 8, 9, 16, 17, 32,
		33, 64, 65, 96, 97, 128, 129, 240, 241,
		512, 1024, 100 * 1024,
	}

	if *cespare {
		for _, size := range sizes {
			b.Run(strconv.Itoa(size), func(b *testing.B) {
				b.SetBytes(int64(size))
				b.ReportAllocs()
				d := string(make([]byte, size))

				for b.Loop() {
					xxhash.Sum64String(d)
				}
			})
		}
	}

	if *zeebo {
		for _, size := range sizes {
			b.Run(strconv.Itoa(size), func(b *testing.B) {
				b.SetBytes(int64(size))
				b.ReportAllocs()
				d := string(make([]byte, size))

				for b.Loop() {
					xxh3.HashString(d)
				}
			})
		}
	}

	if *byteshash {
		for _, size := range sizes {
			b.Run(strconv.Itoa(size), func(b *testing.B) {
				b.SetBytes(int64(size))
				b.ReportAllocs()
				d := string(make([]byte, size))
				seed := maphash.MakeSeed()

				for b.Loop() {
					maphash.String(seed, d)
				}
			})
		}
	}

	if *bytedance {
		for _, size := range sizes {
			b.Run(strconv.Itoa(size), func(b *testing.B) {
				b.SetBytes(int64(size))
				b.ReportAllocs()
				d := string(make([]byte, size))

				for b.Loop() {
					bdxxhash3.HashString(d)
				}
			})
		}
	}

	if *gosimd {
		for _, size := range sizes {
			b.Run(strconv.Itoa(size), func(b *testing.B) {
				b.SetBytes(int64(size))
				b.ReportAllocs()
				d := string(make([]byte, size))

				for b.Loop() {
					simdxxhash.Sum64(unsafe.Slice(unsafe.StringData(d), len(d)))
				}
			})
		}
	}
}
