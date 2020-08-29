// Copyright 2020 The go-benchmarks Authors.
// SPDX-License-Identifier: BSD-3-Clause

package pseudo

import (
	"testing"

	valyala "github.com/valyala/fastrand"
)

func BenchmarkFastrand(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += fastrandn(1e6)
		}
	})
}

func BenchmarkValyala(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		var r valyala.RNG
		s := uint32(0)
		for pb.Next() {
			s += r.Uint32n(1e6)
		}
	})
}
