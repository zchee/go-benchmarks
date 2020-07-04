package walk_benchmark

import (
	"flag"
	"go/build"
	"os"
	"path/filepath"
	"testing"

	"github.com/karrick/godirwalk"
	"github.com/opencontainers/selinux/pkg/pwalk"

	"github.com/zchee/go-benchmarks/path/filepath/walk/fastwalk"
)

var benchDir = flag.String("benchdir", build.Default.GOROOT, "The directory to scan for walk Benchmark")

func BenchmarkFilepathwalk(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := filepath.Walk(*benchDir, func(path string, info os.FileInfo, err error) error { return nil })
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFastWalk(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := fastwalk.Walk(*benchDir, func(path string, typ os.FileMode) error { return nil })
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGodirwalk(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := godirwalk.Walk(*benchDir, &godirwalk.Options{
			Callback:      func(_ string, _ *godirwalk.Dirent) error { return nil },
			ScratchBuffer: make([]byte, os.Getpagesize()),
			Unsorted:      true,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPwalk(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := pwalk.Walk(*benchDir, func(path string, info os.FileInfo, err error) error { return nil })
		if err != nil {
			b.Fatal(err)
		}
	}
}
