# uuid

## Go
- [runtime/uuid.go at master · kata-containers/runtime](https://github.com/kata-containers/runtime/blob/master/virtcontainers/pkg/uuid/uuid.go)
- [rogpeppe/fastuuid: Fast generation of 192-bit UUIDs](https://github.com/rogpeppe/fastuuid)
- [google/uuid: Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.](https://github.com/google/uuid)
- [m4rw3r/uuid: A fast implementation of the UUID datatype for Go, with support for SQL and JSON.](https://github.com/m4rw3r/uuid)
- [satori/go.uuid: UUID package for Go](https://github.com/satori/go.uuid)
- ~~[edwingeng/wuid: An extremely fast UUID alternative written in golang](https://github.com/edwingeng/wuid)~~
- [lithammer/shortuuid: A generator library for concise, unambiguous and URL-safe UUIDs](https://github.com/lithammer/shortuuid)

## C/C++
- [qemu/uuid.c at master · qemu/qemu](https://github.com/qemu/qemu/blob/5acad5bf480321f178866dc28e38eeda5a3f19bb/util/uuid.c#L20-L38)
- [uuid/include/boost/uuid at develop · boostorg/uuid](https://github.com/boostorg/uuid/tree/develop/include/boost/uuid)
- [SIMD-UUID/UUID.h at master · yesmey/SIMD-UUID](https://github.com/yesmey/SIMD-UUID/blob/master/UUID.h)

```console
$ go test -v -run='^$' -bench=. -benchmem -count 10 .
goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/uuid
BenchmarkFastUUID
BenchmarkFastUUID-16              	18713439	        61.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	20132677	        66.5 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	17861589	        63.4 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	19992573	        63.7 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	20263380	        66.5 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	19062441	        64.5 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	19760810	        63.2 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	20377656	        66.2 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	17268458	        64.3 ns/op	      48 B/op	       1 allocs/op
BenchmarkFastUUID-16              	19850604	        64.4 ns/op	      48 B/op	       1 allocs/op
BenchmarkVirtContainersUUID
BenchmarkVirtContainersUUID-16    	 1896793	       608 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 2043817	       639 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 2116083	       629 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 1960450	       602 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 2009354	       609 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 1841523	       620 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 1989430	       612 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 1974718	       603 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 2012625	       608 ns/op	      80 B/op	       5 allocs/op
BenchmarkVirtContainersUUID-16    	 1850124	       585 ns/op	      80 B/op	       5 allocs/op
Benchmark_m4rw3rUUID
Benchmark_m4rw3rUUID-16           	 8341630	       149 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 8625968	       146 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 8383804	       138 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 6984210	       150 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 8582816	       150 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 8794380	       135 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 7097628	       147 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 8760465	       146 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 8911485	       134 ns/op	      64 B/op	       2 allocs/op
Benchmark_m4rw3rUUID-16           	 7197931	       146 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID
BenchmarkGoogleUUID-16            	 9005892	       148 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8381691	       132 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8160595	       153 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8318919	       145 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 9356400	       145 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8384248	       146 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8404401	       133 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8843724	       150 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 9577484	       199 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleUUID-16            	 8450498	       151 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID
BenchmarkSatoriUUID-16            	 8637020	       143 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 8515070	       143 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 8030643	       137 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 9239304	       145 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 8607825	       146 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 8649537	       136 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 8565550	       148 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 9124804	       146 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 8440560	       139 ns/op	      64 B/op	       2 allocs/op
BenchmarkSatoriUUID-16            	 7325452	       143 ns/op	      64 B/op	       2 allocs/op
BenchmarkShortUUID
BenchmarkShortUUID-16             	  205273	      6090 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  201103	      5753 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  219541	      6096 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  206103	      6164 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  198032	      5731 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  218810	      6074 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  206460	      6175 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  179667	      5807 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  214382	      6067 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID-16             	  204294	      5885 ns/op	    2953 B/op	     136 allocs/op
BenchmarkShortUUID2
BenchmarkShortUUID2-16            	 2804804	       372 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3271734	       394 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3334507	       391 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3241248	       392 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3171487	       390 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 2942880	       368 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3068930	       392 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3292690	       397 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3280203	       390 ns/op	      96 B/op	       3 allocs/op
BenchmarkShortUUID2-16            	 3116347	       384 ns/op	      96 B/op	       3 allocs/op
PASS
ok  	github.com/zchee/go-benchmarks/uuid	103.797s
```
