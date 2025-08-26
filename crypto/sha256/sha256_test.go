package sha256_test

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"testing"

	cpuid "github.com/klauspost/cpuid/v2"
	sha256simd "github.com/minio/sha256-simd"
)

var (
	hasIntelSha = runtime.GOARCH == "amd64" && cpuid.CPU.Supports(cpuid.SSSE3, cpuid.SSE4)
	hasAvx512   = cpuid.CPU.Supports(cpuid.AVX512F, cpuid.AVX512DQ, cpuid.AVX512BW, cpuid.AVX512VL)
)

func hasArmSha2() bool {
	if cpuid.CPU.Has(cpuid.SHA2) {
		return true
	}
	if runtime.GOARCH != "arm64" || runtime.GOOS != "linux" {
		return false
	}

	// Fall back to hacky cpuinfo parsing...
	const procCPUInfo = "/proc/cpuinfo"

	// Feature to check for.
	const sha256Feature = "sha2"

	cpuInfo, err := os.ReadFile(procCPUInfo)
	if err != nil {
		return false
	}
	return bytes.Contains(cpuInfo, []byte(sha256Feature))
}

type blockfuncType int

const (
	blockfuncStdlib blockfuncType = iota
	blockfuncIntelSha
	blockfuncArmSha2
	blockfuncForceGeneric = -1
)

var blockfunc blockfuncType

func init() {
	switch {
	case hasIntelSha:
		blockfunc = blockfuncIntelSha
	case hasArmSha2():
		blockfunc = blockfuncArmSha2
	}
}

func benchmarkSize(b *testing.B, size int) {
	bench := sha256simd.New()
	var buf = make([]byte, size)
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())

	for b.Loop() {
		bench.Reset()
		bench.Write(buf)
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash(b *testing.B) {
	type alg struct {
		n string
		t blockfuncType
	}
	algos := make([]alg, 0, 3)

	fmt.Printf("runtime.GOARCH == amd64: %#v\n", runtime.GOARCH == "amd64")
	fmt.Printf("hasIntelSha: %#v\n", hasIntelSha)
	algos = append(algos, alg{"Generic", blockfuncForceGeneric})
	if hasIntelSha {
		algos = append(algos, alg{"IntelSHA", blockfuncIntelSha})
	}
	if hasArmSha2() {
		algos = append(algos, alg{"ArmSha2", blockfuncArmSha2})
	}
	algos = append(algos, alg{"GoStdlib", blockfuncStdlib})

	sizes := []struct {
		n string
		f func(*testing.B, int)
		s int
	}{
		{"8Bytes", benchmarkSize, 1 << 3},
		{"64Bytes", benchmarkSize, 1 << 6},
		{"1K", benchmarkSize, 1 << 10},
		{"8K", benchmarkSize, 1 << 13},
		{"1M", benchmarkSize, 1 << 20},
		{"5M", benchmarkSize, 5 << 20},
		{"10M", benchmarkSize, 5 << 21},
	}

	for _, a := range algos {
		func() {
			orig := blockfunc
			defer func() { blockfunc = orig }()

			blockfunc = a.t
			for _, y := range sizes {
				s := a.n + "/" + y.n
				b.Run(s, func(b *testing.B) { y.f(b, y.s) })
			}
		}()
	}
}
