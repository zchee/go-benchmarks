# sync/atomic/add_load_int

Package benchmark is a benchmark for writing and reading to int using atomic operations.

```sh
goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/sync/atomic/add_load_int
BenchmarkAtomicInt32-16      	  299862	      3581 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicInt64-16      	  346306	      3571 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicUint32-16     	  339705	      3611 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicUint64-16     	  315796	      3599 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicUintptr-16    	  348063	      3563 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-16    	  337605	      3578 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/zchee/go-benchmarks/sync/atomic/add_load_int	7.355s
```
