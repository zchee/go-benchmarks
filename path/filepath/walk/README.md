# path/filepath/walk

`filepath.Walk` related packages benchmark.

```console
$ go test -v -run='^$' -count=10 -bench=. -benchmem .

goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/path/filepath/walk
BenchmarkFilepathwalk
BenchmarkFilepathwalk-16    	       9	 115690309 ns/op	 8349374 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 116072363 ns/op	 8349365 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 114452003 ns/op	 8349352 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 114115967 ns/op	 8349376 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 114293446 ns/op	 8349352 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 113795678 ns/op	 8349340 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 114391736 ns/op	 8349361 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 113494281 ns/op	 8349356 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 114230750 ns/op	 8349358 B/op	   56990 allocs/op
BenchmarkFilepathwalk-16    	       9	 114608933 ns/op	 8349352 B/op	   56990 allocs/op
BenchmarkFastWalk
BenchmarkFastWalk-16        	      37	  31731268 ns/op	  949545 B/op	   27159 allocs/op
BenchmarkFastWalk-16        	      37	  32069274 ns/op	  947866 B/op	   27146 allocs/op
BenchmarkFastWalk-16        	      37	  31827580 ns/op	  947535 B/op	   27142 allocs/op
BenchmarkFastWalk-16        	      37	  31835709 ns/op	  947516 B/op	   27144 allocs/op
BenchmarkFastWalk-16        	      37	  31900710 ns/op	  947401 B/op	   27145 allocs/op
BenchmarkFastWalk-16        	      36	  31846069 ns/op	  946962 B/op	   27143 allocs/op
BenchmarkFastWalk-16        	      38	  31889972 ns/op	  947480 B/op	   27143 allocs/op
BenchmarkFastWalk-16        	      37	  32115082 ns/op	  947094 B/op	   27143 allocs/op
BenchmarkFastWalk-16        	      37	  32087662 ns/op	  947871 B/op	   27145 allocs/op
BenchmarkFastWalk-16        	      36	  32432550 ns/op	  947044 B/op	   27146 allocs/op
BenchmarkGodirwalk
BenchmarkGodirwalk-16       	       9	 122010095 ns/op	 9148089 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 108260351 ns/op	 9148150 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 109604950 ns/op	 9148131 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 108679832 ns/op	 9148032 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 108994968 ns/op	 9148086 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 109827896 ns/op	 9148067 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 109249252 ns/op	 9147990 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 108245061 ns/op	 9148156 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 108980834 ns/op	 9148034 B/op	   55644 allocs/op
BenchmarkGodirwalk-16       	      10	 107923286 ns/op	 9148206 B/op	   55644 allocs/op
BenchmarkPwalk
BenchmarkPwalk-16           	       8	 134722138 ns/op	 8905842 B/op	   80038 allocs/op
BenchmarkPwalk-16           	       8	 134921244 ns/op	 8907290 B/op	   80041 allocs/op
BenchmarkPwalk-16           	       8	 134328248 ns/op	 8903120 B/op	   80030 allocs/op
BenchmarkPwalk-16           	       8	 134767174 ns/op	 8906038 B/op	   80037 allocs/op
BenchmarkPwalk-16           	       8	 134477830 ns/op	 8902942 B/op	   80028 allocs/op
BenchmarkPwalk-16           	       8	 134849876 ns/op	 8903352 B/op	   80032 allocs/op
BenchmarkPwalk-16           	       8	 134403460 ns/op	 8903144 B/op	   80030 allocs/op
BenchmarkPwalk-16           	       8	 134226129 ns/op	 8902902 B/op	   80028 allocs/op
BenchmarkPwalk-16           	       8	 134997118 ns/op	 8903080 B/op	   80029 allocs/op
BenchmarkPwalk-16           	       8	 134982394 ns/op	 8902902 B/op	   80028 allocs/op
BenchmarkPowerwalk
BenchmarkPowerwalk-16       	       8	 136258882 ns/op	 8908592 B/op	   68534 allocs/op
BenchmarkPowerwalk-16       	       8	 137140995 ns/op	 8903602 B/op	   68523 allocs/op
BenchmarkPowerwalk-16       	       8	 141930287 ns/op	 8905968 B/op	   68523 allocs/op
BenchmarkPowerwalk-16       	       8	 136164633 ns/op	 8906002 B/op	   68540 allocs/op
BenchmarkPowerwalk-16       	       8	 136750587 ns/op	 8904494 B/op	   68524 allocs/op
BenchmarkPowerwalk-16       	       8	 135664846 ns/op	 8905534 B/op	   68535 allocs/op
BenchmarkPowerwalk-16       	       8	 135359850 ns/op	 8905222 B/op	   68540 allocs/op
BenchmarkPowerwalk-16       	       8	 135114858 ns/op	 8903716 B/op	   68524 allocs/op
BenchmarkPowerwalk-16       	       8	 138815389 ns/op	 8904146 B/op	   68529 allocs/op
BenchmarkPowerwalk-16       	       8	 136183113 ns/op	 8904678 B/op	   68525 allocs/op
PASS
ok  	github.com/zchee/go-benchmarks/path/filepath/walk	63.049s
```

```console
$ benchstat bench.txt

name             time/op
Filepathwalk-16   114ms ± 1%
FastWalk-16      32.0ms ± 1%
Godirwalk-16      109ms ± 1%
Pwalk-16          135ms ± 0%
Powerwalk-16      136ms ± 2%

name             alloc/op
Filepathwalk-16  8.35MB ± 0%
FastWalk-16       947kB ± 0%
Godirwalk-16     9.15MB ± 0%
Pwalk-16         8.90MB ± 0%
Powerwalk-16     8.91MB ± 0%

name             allocs/op
Filepathwalk-16   57.0k ± 0%
FastWalk-16       27.1k ± 0%
Godirwalk-16      55.6k ± 0%
Pwalk-16          80.0k ± 0%
Powerwalk-16      68.5k ± 0%
```
