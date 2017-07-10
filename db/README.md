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
goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/db
BenchmarkGetBadger-8             5000000               288 ns/op          31.22 MB/s          40 B/op          2 allocs/op
BenchmarkGetLevelDB-8            3000000               557 ns/op          16.16 MB/s         112 B/op          4 allocs/op
BenchmarkGetBolt-8               2000000               784 ns/op          11.47 MB/s         440 B/op          7 allocs/op
PASS
ok      github.com/zchee/go-benchmarks/db       8.509s
```
