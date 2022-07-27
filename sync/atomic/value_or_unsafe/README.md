# value_or_unsafe 

```console
$ go test -v -run='^$' -benchmem -count=10 -bench=. .
goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/sync/atomic/value_or_unsafe
cpu: Intel(R) Xeon(R) W-2150B CPU @ 3.00GHz
BenchmarkAtomicValue
BenchmarkAtomicValue-20      	907985367	         1.312 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	924415546	         1.312 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	923279490	         1.312 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	910984804	         1.305 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	917892320	         1.307 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	921821854	         1.319 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	902309553	         1.305 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	918721095	         1.313 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	891985299	         1.314 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicValue-20      	905389354	         1.307 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer
BenchmarkAtomicPointer-20    	1000000000	         0.4721 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4747 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4754 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4753 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4790 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4730 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4732 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4737 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4750 ns/op	       0 B/op	       0 allocs/op
BenchmarkAtomicPointer-20    	1000000000	         0.4750 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/zchee/go-benchmarks/sync/atomic/value_or_unsafe	18.570s
```
