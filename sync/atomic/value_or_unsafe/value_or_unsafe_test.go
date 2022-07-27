package benchmark

import (
	"sync/atomic"
	"testing"
	"unsafe"
)

type testStruct struct{}

var valStruct *testStruct

func BenchmarkAtomicValue(b *testing.B) {
	var val atomic.Value
	val.Store(&testStruct{})

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		valStruct = val.Load().(*testStruct)
	}
}

func BenchmarkAtomicPointer(b *testing.B) {
	var val unsafe.Pointer
	atomic.StorePointer(&val, unsafe.Pointer(&testStruct{}))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		valStruct = (*testStruct)(atomic.LoadPointer(&val))
	}
}
