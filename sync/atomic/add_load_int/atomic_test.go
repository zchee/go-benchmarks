// Copyright 2019 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package benchmark is a benchmark for writing and reading to int using atomic operations.
package benchmark

import (
	"sync/atomic"
	"testing"
	"unsafe"
)

const (
	benchMarkSize     = 1 << 6 // 64
	benchMarkSizeLong = benchMarkSize << 3
)

func BenchmarkAtomicInt32(b *testing.B) {
	var a int32

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		a = 0
		for i := int32(0); i < benchMarkSizeLong; i++ {
			atomic.AddInt32(&a, 1)

			if atomic.LoadInt32(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicInt64(b *testing.B) {
	var a int64

	for n := 0; n < b.N; n++ {
		a = 0
		for i := int64(0); i < benchMarkSizeLong; i++ {
			atomic.AddInt64(&a, 1)

			if atomic.LoadInt64(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicUint32(b *testing.B) {
	var a uint32

	for n := 0; n < b.N; n++ {
		a = 0
		for i := uint32(0); i < benchMarkSizeLong; i++ {
			atomic.AddUint32(&a, 1)

			if atomic.LoadUint32(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicUint64(b *testing.B) {
	var a uint64

	for n := 0; n < b.N; n++ {
		a = 0
		for i := uint64(0); i < benchMarkSizeLong; i++ {
			atomic.AddUint64(&a, 1)

			if atomic.LoadUint64(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicUintptr(b *testing.B) {
	var a uintptr

	for n := 0; n < b.N; n++ {
		a = 0
		for i := uintptr(0); i < benchMarkSizeLong; i++ {
			atomic.AddUintptr(&a, 1)

			if atomic.LoadUintptr(&a) != (i + 1) {
				b.Fail()
			}
		}
	}
}

func BenchmarkAtomicPointer(b *testing.B) {
	var a unsafe.Pointer

	for n := 0; n < b.N; n++ {
		a = unsafe.Pointer(new(uintptr))
		for i := uintptr(0); i < benchMarkSizeLong; i++ {
			atomic.AddUintptr((*uintptr)(a), 1)

			if atomic.LoadUintptr((*uintptr)(a)) != (i + 1) {
				b.Fail()
			}
		}
	}
}
