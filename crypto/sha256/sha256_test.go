/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Using this part of Minio codebase under the license
// Apache License Version 2.0 with modifications

// SHA256 hash algorithm.  See FIPS 180-2.

package sha256_test

import (
	"bytes"
	"os"
	"runtime"
	"testing"

	cpuid "github.com/klauspost/cpuid/v2"
	sha256simd "github.com/minio/sha256-simd"
)

var (
	hasIntelSha = runtime.GOARCH == "amd64" && cpuid.CPU.Supports(cpuid.SHA, cpuid.SSSE3, cpuid.SSE4)
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

func BenchmarkHash(b *testing.B) {
	type alg struct {
		n string
		t blockfuncType
	}
	algos := make([]alg, 0, 2)

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
