# path/filepath/walk

`filepath.Walk` related packages benchmark.

```console
$ go test -v -run=^$ -count=10 -bench=. -benchmem .

goos: darwin
goarch: amd64
pkg: github.com/zchee/go-benchmarks/path/filepath/walk
BenchmarkFilepathwalk
BenchmarkFilepathwalk-16    	      10	 110762716 ns/op	 7979750 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	      10	 109474622 ns/op	 7979553 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	       9	 111859005 ns/op	 7979752 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	       9	 111225096 ns/op	 7979543 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	      10	 111339014 ns/op	 7979525 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	      10	 111331049 ns/op	 7979523 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	       9	 111717218 ns/op	 7979530 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	      10	 111163409 ns/op	 7979520 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	      10	 109286129 ns/op	 7979518 B/op	   54913 allocs/op
BenchmarkFilepathwalk-16    	      10	 112996944 ns/op	 7979544 B/op	   54913 allocs/op
BenchmarkFastWalk
BenchmarkFastWalk-16        	      40	  29070998 ns/op	  903265 B/op	   26187 allocs/op
BenchmarkFastWalk-16        	      40	  29225386 ns/op	  901800 B/op	   26175 allocs/op
BenchmarkFastWalk-16        	      40	  29475843 ns/op	  901802 B/op	   26174 allocs/op
BenchmarkFastWalk-16        	      40	  29396814 ns/op	  902393 B/op	   26175 allocs/op
BenchmarkFastWalk-16        	      38	  29358819 ns/op	  902080 B/op	   26174 allocs/op
BenchmarkFastWalk-16        	      40	  29382107 ns/op	  900670 B/op	   26174 allocs/op
BenchmarkFastWalk-16        	      38	  29614305 ns/op	  902096 B/op	   26174 allocs/op
BenchmarkFastWalk-16        	      40	  29829689 ns/op	  900877 B/op	   26173 allocs/op
BenchmarkFastWalk-16        	      39	  29825739 ns/op	  902176 B/op	   26173 allocs/op
BenchmarkFastWalk-16        	      40	  29657151 ns/op	  901689 B/op	   26173 allocs/op
BenchmarkGodirwalk
BenchmarkGodirwalk-16       	      10	 111497738 ns/op	 8701694 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 109479668 ns/op	 8701716 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 109425017 ns/op	 8701668 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 107113736 ns/op	 8701752 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 108049259 ns/op	 8701671 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 106667290 ns/op	 8701695 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 107778048 ns/op	 8701655 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 108456500 ns/op	 8701712 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 108179095 ns/op	 8701746 B/op	   53637 allocs/op
BenchmarkGodirwalk-16       	      10	 109400274 ns/op	 8701747 B/op	   53637 allocs/op
BenchmarkPwalk
BenchmarkPwalk-16           	       8	 133905817 ns/op	 8516052 B/op	   77193 allocs/op
BenchmarkPwalk-16           	       8	 134757255 ns/op	 8519744 B/op	   77201 allocs/op
BenchmarkPwalk-16           	       8	 133639575 ns/op	 8515170 B/op	   77190 allocs/op
BenchmarkPwalk-16           	       8	 133132960 ns/op	 8515261 B/op	   77194 allocs/op
BenchmarkPwalk-16           	       8	 133889081 ns/op	 8514912 B/op	   77190 allocs/op
BenchmarkPwalk-16           	       8	 134523434 ns/op	 8514962 B/op	   77190 allocs/op
BenchmarkPwalk-16           	       8	 135024137 ns/op	 8514799 B/op	   77189 allocs/op
BenchmarkPwalk-16           	       8	 132917612 ns/op	 8515592 B/op	   77197 allocs/op
BenchmarkPwalk-16           	       8	 132481027 ns/op	 8515166 B/op	   77192 allocs/op
BenchmarkPwalk-16           	       8	 133226718 ns/op	 8517176 B/op	   77193 allocs/op
PASS
ok  	github.com/zchee/go-benchmarks/path/filepath/walk	54.404s
```

```console
$ benchstat -geomean bench.txt
name             time/op
Filepathwalk-16   111ms ± 2%
FastWalk-16      29.5ms ± 1%
Godirwalk-16      109ms ± 3%
Pwalk-16          134ms ± 1%
[Geo mean]       83.1ms     

name             alloc/op
Filepathwalk-16  7.98MB ± 0%
FastWalk-16       902kB ± 0%
Godirwalk-16     8.70MB ± 0%
Pwalk-16         8.52MB ± 0%
[Geo mean]       4.81MB     

name             allocs/op
Filepathwalk-16   54.9k ± 0%
FastWalk-16       26.2k ± 0%
Godirwalk-16      53.6k ± 0%
Pwalk-16          77.2k ± 0%
[Geo mean]        49.4k     
```
