# db

Benchmark for NoSQL database engines.

-	[dgraph-io/badger](https://github.com/dgraph-io/badger)
	-	Fastest key-value store in Go.
-	[syndtr/goleveldb](https://github.com/syndtr/goleveldb)
	-	LevelDB key/value database in Go.
-	[boltdb/bolt](https://github.com/boltdb/bolt)
	-	An embedded key/value database for Go.

## Result

```sh
> go test -run='^$' -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/db
BenchmarkGetBadger-8             5000000               287 ns/op          31.26 MB/s          40 B/op          2 allocs/op
BenchmarkGetLevelDB-8            3000000               558 ns/op          16.12 MB/s         112 B/op          4 allocs/op
BenchmarkGetBolt-8               2000000               766 ns/op          11.74 MB/s         440 B/op          7 allocs/op
PASS
ok      github.com/zchee/go-benchmarks/db       9.114s
```
