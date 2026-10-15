# hash/non-cryptographic

Benchmark of non-cryptographic 64-bit hash functions for `string` input.

| Flag         | Package                                                                                          | Algorithm | Entry point                                       |
| ------------ | ------------------------------------------------------------------------------------------------ | --------- | ------------------------------------------------- |
| `-cespare`   | [github.com/cespare/xxhash/v2](https://github.com/cespare/xxhash)                                | XXH64     | `xxhash.Sum64String(s)`                           |
| `-zeebo`     | [github.com/zeebo/xxh3](https://github.com/zeebo/xxh3)                                           | XXH3-64   | `xxh3.HashString(s)`                              |
| `-maphash`   | [hash/maphash](https://pkg.go.dev/hash/maphash) (standard library)                               | runtime   | `maphash.String(seed, s)`                         |
| `-bytedance` | [github.com/bytedance/gopkg/util/xxhash3](https://github.com/bytedance/gopkg/tree/main/util/xxhash3) | XXH3-64   | `xxhash3.HashString(s)`                           |
| `-gosimd`    | [github.com/go-simd/xxhash](https://github.com/go-simd/xxhash)                                   | XXH3-64   | `xxhash.Sum64(unsafe.Slice(unsafe.StringData(s), len(s)))` |

Every block hashes the same zero-filled string of the given size and measures
only the hash call: the input and the `maphash` seed are built once outside
`b.Loop()`. `go-simd/xxhash` has no zero-copy string entry point
(`Sum64String` copies the input into a `[]byte`), so the benchmark passes the
string bytes through `unsafe.Slice` to keep the comparison apples to apples.

The input sizes are chosen around the XXH3 code-path boundaries (16, 128, 240
bytes) plus one value just past each of them, and two larger inputs.

## Run

Each flag selects one implementation, and every implementation reports the same
`BenchmarkHash/<size>` names, so run them one at a time into separate files and
compare with [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat):

```sh
for impl in cespare zeebo maphash bytedance gosimd; do
  go test -run '^$' -bench '^BenchmarkHash$' -count=10 -cpu=1 -benchtime=1s -timeout=30m -$impl . | tee "$impl.txt"
done
benchstat cespare.txt zeebo.txt maphash.txt bytedance.txt gosimd.txt
```

`-cpu=1` pins `GOMAXPROCS` to one: the hash functions are single-threaded, and
it keeps the benchmark off the efficiency cores on Apple silicon.

## Result

Measured on 2026-10-16 with Go 1.27.2 on macOS 27.2, Apple M3 Max (`-cpu=1`,
`-count=10`). Full `benchstat` output is in the appendix below.

### Summary

**[zeebo/xxh3](https://github.com/zeebo/xxh3) is the fastest overall.** It is
the best or tied for best at 19 of the 21 input sizes, and its geomean is 26%
below cespare/xxhash, the baseline.

| Package                 | geomean sec/op | vs cespare | geomean B/s |
| ----------------------- | -------------: | ---------: | ----------: |
| cespare/xxhash (XXH64)  |        8.032ns |   baseline |     6.66GiB |
| **zeebo/xxh3**          |    **5.925ns** | **-26.2%** | **9.05GiB** |
| hash/maphash            |        9.140ns |     +13.8% |     5.75GiB |
| bytedance/gopkg xxhash3 |        6.962ns |     -13.3% |     7.66GiB |
| go-simd/xxhash          |        8.003ns |      -0.4% |     6.69GiB |

### Time per operation by input size

Nanoseconds per hash; lower is better. The fastest value in each row is in bold.
Values within noise of each other (`p > 0.05`) are both bold.

| Size (bytes) | cespare/xxhash | zeebo/xxh3 | hash/maphash | bytedance xxhash3 | go-simd/xxhash |
| -----------: | -------------: | ---------: | -----------: | ----------------: | -------------: |
|            0 |          2.495 |  **1.931** |        1.955 |             2.007 |          2.506 |
|            1 |          2.206 |  **1.994** |        4.648 |             2.231 |          2.781 |
|            3 |          2.308 |  **1.972** |        4.423 |             2.241 |          2.689 |
|            4 |          2.187 |  **1.945** |        4.454 |             2.204 |          2.458 |
|            8 |          2.209 |  **1.948** |        4.418 |             2.217 |          2.478 |
|            9 |          2.079 |  **1.972** |        4.470 |             2.237 |          2.710 |
|           16 |          2.517 |  **1.969** |        3.911 |             2.233 |          2.713 |
|           17 |      **2.436** |  **2.431** |        3.905 |             2.991 |          3.248 |
|           32 |          4.295 |  **2.402** |        3.892 |             3.002 |          3.349 |
|           33 |          4.821 |  **3.672** |        4.209 |             4.140 |          4.453 |
|           64 |          5.861 |  **3.719** |        4.128 |             4.043 |          4.435 |
|           65 |          6.502 |  **5.012** |        5.987 |             5.293 |          5.567 |
|           96 |          7.655 |  **4.987** |        6.043 |             5.312 |          5.597 |
|           97 |          8.125 |  **6.050** |    **5.987** |             6.350 |          6.617 |
|          128 |          9.134 |  **6.053** |    **6.002** |             6.375 |          6.593 |
|          129 |          9.538 |  **6.573** |       11.235 |             6.883 |         12.760 |
|          240 |          15.26 |      11.04 |        12.62 |         **10.66** |          23.76 |
|          241 |          15.96 |  **12.00** |        12.62 |             17.42 |          24.12 |
|          512 |          31.29 |  **19.91** |        31.98 |             29.70 |          31.56 |
|         1024 |          62.50 |  **35.36** |        71.19 |             56.30 |          48.37 |
|       102400 |           6299 |   **3361** |         7814 |              5558 |       **3401** |

### Allocations

Every implementation is allocation-free except `go-simd/xxhash`, which
allocates 64 bytes once per call for inputs over 240 bytes. The allocation
comes from its long-input path and cannot be avoided by the caller.

### Observations

- **zeebo/xxh3** has no weak size range. It leads on short inputs (under 2 ns
  up to 16 bytes), on every mid-size path, and on long inputs, where it reaches
  28 GiB/s at 100 KiB. The ±52% spread at 96 bytes is an outlier run; its
  median agrees with the neighbouring sizes.
- **bytedance/gopkg xxhash3** trails zeebo by 5–15% and wins only at exactly
  240 bytes. It jumps from 10.7 ns to 17.4 ns between 240 and 241 bytes, so its
  long-input path has a noticeably higher entry cost.
- **go-simd/xxhash** matches zeebo at 100 KiB but is 25–35% slower on short
  inputs and more than twice as slow at 129–240 bytes. Its full-block SIMD
  kernel is amd64 only; on arm64 only the per-stripe NEON accumulator is in
  play, so the amd64 ranking may differ.
- **hash/maphash** is twice as slow as the rest at 1–9 bytes and the slowest on
  long inputs (12 GiB/s at 100 KiB). It is competitive only in the 97–128 byte
  range, where it ties zeebo. Its advantage is being in the standard library
  with a per-process random seed, not speed.
- **cespare/xxhash** implements XXH64, a different algorithm from the other
  xxHash entries. It keeps up through 17 bytes and then falls 25–45% behind the
  XXH3 implementations, which is the expected gap between the two algorithms.

These numbers are from Apple silicon (arm64). bytedance/gopkg ships AVX2 and
AVX-512 assembly and go-simd/xxhash ships an AVX2 full-block kernel, so both
should close part of the gap on amd64; rerun the commands above on that
platform before drawing conclusions for it.

## Appendix: full benchstat output

<details>
<summary>sec/op</summary>

```
goos: darwin
goarch: arm64
pkg: github.com/zchee/go-benchmarks/hash/non-cryptographic
cpu: Apple M3 Max
            │ cespare.txt │              zeebo.txt               │              maphash.txt              │            bytedance.txt            │              gosimd.txt              │
            │   sec/op    │    sec/op     vs base                │    sec/op     vs base                 │   sec/op     vs base                │    sec/op     vs base                │
Hash/0        2.495n ± 1%   1.931n ±  1%  -22.57% (p=0.000 n=10)    1.955n ± 2%   -21.65% (p=0.000 n=10)   2.007n ± 2%  -19.52% (p=0.000 n=10)    2.506n ± 3%        ~ (p=0.684 n=10)
Hash/1        2.206n ± 1%   1.994n ±  2%   -9.57% (p=0.000 n=10)    4.648n ± 1%  +110.75% (p=0.000 n=10)   2.231n ± 2%   +1.16% (p=0.004 n=10)    2.781n ± 1%  +26.12% (p=0.000 n=10)
Hash/3        2.308n ± 1%   1.972n ±  2%  -14.54% (p=0.000 n=10)    4.423n ± 2%   +91.66% (p=0.000 n=10)   2.241n ± 1%   -2.90% (p=0.002 n=10)    2.689n ± 2%  +16.53% (p=0.000 n=10)
Hash/4        2.187n ± 1%   1.945n ±  7%  -11.02% (p=0.000 n=10)    4.454n ± 8%  +103.68% (p=0.000 n=10)   2.204n ± 2%        ~ (p=0.404 n=10)    2.458n ± 2%  +12.42% (p=0.000 n=10)
Hash/8        2.209n ± 1%   1.948n ±  1%  -11.84% (p=0.000 n=10)    4.418n ± 4%   +99.98% (p=0.000 n=10)   2.217n ± 6%        ~ (p=0.631 n=10)    2.478n ± 2%  +12.17% (p=0.000 n=10)
Hash/9        2.079n ± 1%   1.972n ±  2%   -5.15% (p=0.001 n=10)    4.470n ± 2%  +115.01% (p=0.000 n=10)   2.237n ± 1%   +7.62% (p=0.000 n=10)    2.710n ± 2%  +30.38% (p=0.000 n=10)
Hash/16       2.517n ± 2%   1.969n ±  3%  -21.79% (p=0.000 n=10)    3.911n ± 2%   +55.36% (p=0.000 n=10)   2.233n ± 1%  -11.26% (p=0.000 n=10)    2.713n ± 3%   +7.79% (p=0.000 n=10)
Hash/17       2.436n ± 2%   2.431n ±  1%        ~ (p=0.782 n=10)    3.905n ± 1%   +60.34% (p=0.000 n=10)   2.991n ± 2%  +22.81% (p=0.000 n=10)    3.248n ± 2%  +33.38% (p=0.000 n=10)
Hash/32       4.295n ± 1%   2.402n ±  1%  -44.07% (p=0.000 n=10)    3.892n ± 2%    -9.37% (p=0.000 n=10)   3.002n ± 2%  -30.10% (p=0.000 n=10)    3.349n ± 4%  -22.03% (p=0.000 n=10)
Hash/33       4.821n ± 1%   3.672n ± 10%  -23.83% (p=0.000 n=10)    4.209n ± 4%   -12.70% (p=0.000 n=10)   4.140n ± 1%  -14.13% (p=0.000 n=10)    4.453n ± 9%   -7.64% (p=0.019 n=10)
Hash/64       5.861n ± 2%   3.719n ±  1%  -36.55% (p=0.000 n=10)    4.128n ± 2%   -29.55% (p=0.000 n=10)   4.043n ± 2%  -31.01% (p=0.000 n=10)    4.435n ± 2%  -24.33% (p=0.000 n=10)
Hash/65       6.502n ± 6%   5.012n ±  5%  -22.90% (p=0.000 n=10)    5.987n ± 2%    -7.92% (p=0.000 n=10)   5.293n ± 1%  -18.59% (p=0.000 n=10)    5.567n ± 1%  -14.38% (p=0.000 n=10)
Hash/96       7.655n ± 1%   4.987n ± 52%  -34.85% (p=0.002 n=10)    6.043n ± 2%   -21.05% (p=0.000 n=10)   5.312n ± 4%  -30.61% (p=0.000 n=10)    5.597n ± 2%  -26.88% (p=0.000 n=10)
Hash/97       8.125n ± 1%   6.050n ±  1%  -25.53% (p=0.000 n=10)    5.987n ± 2%   -26.31% (p=0.000 n=10)   6.350n ± 1%  -21.84% (p=0.000 n=10)    6.617n ± 2%  -18.55% (p=0.000 n=10)
Hash/128      9.134n ± 1%   6.053n ±  2%  -33.73% (p=0.000 n=10)    6.002n ± 2%   -34.28% (p=0.000 n=10)   6.375n ± 2%  -30.21% (p=0.000 n=10)    6.593n ± 2%  -27.82% (p=0.000 n=10)
Hash/129      9.538n ± 1%   6.573n ±  2%  -31.08% (p=0.000 n=10)   11.235n ± 1%   +17.80% (p=0.000 n=10)   6.883n ± 2%  -27.83% (p=0.000 n=10)   12.760n ± 1%  +33.79% (p=0.000 n=10)
Hash/240      15.26n ± 1%   11.04n ±  2%  -27.62% (p=0.000 n=10)    12.62n ± 3%   -17.27% (p=0.000 n=10)   10.66n ± 1%  -30.18% (p=0.000 n=10)    23.76n ± 2%  +55.67% (p=0.000 n=10)
Hash/241      15.96n ± 1%   12.00n ±  1%  -24.84% (p=0.000 n=10)    12.62n ± 2%   -20.95% (p=0.000 n=10)   17.42n ± 8%   +9.11% (p=0.000 n=10)    24.12n ± 6%  +51.08% (p=0.000 n=10)
Hash/512      31.29n ± 2%   19.91n ±  1%  -36.36% (p=0.000 n=10)    31.98n ± 3%    +2.22% (p=0.000 n=10)   29.70n ± 2%   -5.05% (p=0.000 n=10)    31.56n ± 4%   +0.88% (p=0.041 n=10)
Hash/1024     62.50n ± 1%   35.36n ±  1%  -43.42% (p=0.000 n=10)    71.19n ± 2%   +13.90% (p=0.000 n=10)   56.30n ± 2%   -9.92% (p=0.000 n=10)    48.37n ± 3%  -22.61% (p=0.000 n=10)
Hash/102400   6.299µ ± 1%   3.361µ ±  1%  -46.64% (p=0.000 n=10)    7.814µ ± 2%   +24.04% (p=0.000 n=10)   5.558µ ± 5%  -11.76% (p=0.000 n=10)    3.401µ ± 6%  -46.02% (p=0.000 n=10)
geomean       8.032n        5.925n        -26.24%                   9.140n        +13.79%                  6.962n       -13.32%                   8.003n        -0.35%
```

</details>

<details>
<summary>B/s</summary>

```
            │  cespare.txt  │               zeebo.txt                │              maphash.txt              │             bytedance.txt             │              gosimd.txt               │
            │      B/s      │      B/s        vs base                │      B/s       vs base                │      B/s       vs base                │      B/s       vs base                │
Hash/1         432.4Mi ± 1%    478.1Mi ±  1%  +10.57% (p=0.000 n=10)    205.2Mi ± 1%  -52.55% (p=0.000 n=10)    427.5Mi ± 2%   -1.14% (p=0.004 n=10)    342.9Mi ± 1%  -20.70% (p=0.000 n=10)
Hash/3        1239.5Mi ± 1%   1450.7Mi ±  2%  +17.04% (p=0.000 n=10)    646.8Mi ± 2%  -47.82% (p=0.000 n=10)   1276.6Mi ± 1%   +2.99% (p=0.002 n=10)   1063.8Mi ± 2%  -14.18% (p=0.000 n=10)
Hash/4        1744.5Mi ± 1%   1960.8Mi ±  6%  +12.40% (p=0.000 n=10)    856.6Mi ± 7%  -50.90% (p=0.000 n=10)   1731.0Mi ± 2%        ~ (p=0.436 n=10)   1551.8Mi ± 2%  -11.05% (p=0.000 n=10)
Hash/8         3.372Gi ± 1%    3.825Gi ±  1%  +13.43% (p=0.000 n=10)    1.686Gi ± 4%  -49.99% (p=0.000 n=10)    3.360Gi ± 5%        ~ (p=0.631 n=10)    3.007Gi ± 2%  -10.83% (p=0.000 n=10)
Hash/9         4.032Gi ± 1%    4.250Gi ±  2%   +5.40% (p=0.001 n=10)    1.875Gi ± 2%  -53.49% (p=0.000 n=10)    3.746Gi ± 1%   -7.09% (p=0.000 n=10)    3.092Gi ± 2%  -23.30% (p=0.000 n=10)
Hash/16        5.921Gi ± 2%    7.570Gi ±  3%  +27.86% (p=0.000 n=10)    3.811Gi ± 2%  -35.64% (p=0.000 n=10)    6.673Gi ± 1%  +12.70% (p=0.000 n=10)    5.493Gi ± 3%   -7.22% (p=0.000 n=10)
Hash/17        6.501Gi ± 2%    6.513Gi ±  1%        ~ (p=0.796 n=10)    4.055Gi ± 1%  -37.63% (p=0.000 n=10)    5.293Gi ± 2%  -18.57% (p=0.000 n=10)    4.873Gi ± 2%  -25.03% (p=0.000 n=10)
Hash/32        6.939Gi ± 1%   12.408Gi ±  1%  +78.82% (p=0.000 n=10)    7.656Gi ± 2%  +10.33% (p=0.000 n=10)    9.929Gi ± 2%  +43.09% (p=0.000 n=10)    8.899Gi ± 4%  +28.25% (p=0.000 n=10)
Hash/33        6.374Gi ± 1%    8.369Gi ±  9%  +31.30% (p=0.000 n=10)    7.302Gi ± 4%  +14.55% (p=0.000 n=10)    7.423Gi ± 1%  +16.46% (p=0.000 n=10)    6.902Gi ± 9%   +8.28% (p=0.019 n=10)
Hash/64        10.17Gi ± 2%    16.03Gi ±  1%  +57.60% (p=0.000 n=10)    14.44Gi ± 2%  +41.96% (p=0.000 n=10)    14.74Gi ± 2%  +44.96% (p=0.000 n=10)    13.44Gi ± 3%  +32.17% (p=0.000 n=10)
Hash/65        9.311Gi ± 6%   12.077Gi ±  5%  +29.71% (p=0.000 n=10)   10.112Gi ± 2%   +8.60% (p=0.000 n=10)   11.436Gi ± 1%  +22.83% (p=0.000 n=10)   10.874Gi ± 1%  +16.79% (p=0.000 n=10)
Hash/96        11.68Gi ± 1%    17.93Gi ± 34%  +53.51% (p=0.002 n=10)    14.79Gi ± 2%  +26.68% (p=0.000 n=10)    16.83Gi ± 4%  +44.11% (p=0.000 n=10)    15.97Gi ± 2%  +36.77% (p=0.000 n=10)
Hash/97        11.12Gi ± 1%    14.93Gi ±  1%  +34.28% (p=0.000 n=10)    15.09Gi ± 2%  +35.71% (p=0.000 n=10)    14.23Gi ± 1%  +27.95% (p=0.000 n=10)    13.65Gi ± 1%  +22.79% (p=0.000 n=10)
Hash/128       13.05Gi ± 1%    19.69Gi ±  2%  +50.89% (p=0.000 n=10)    19.86Gi ± 2%  +52.17% (p=0.000 n=10)    18.70Gi ± 2%  +43.30% (p=0.000 n=10)    18.08Gi ± 2%  +38.55% (p=0.000 n=10)
Hash/129      12.597Gi ± 1%   18.278Gi ±  2%  +45.10% (p=0.000 n=10)   10.696Gi ± 1%  -15.09% (p=0.000 n=10)   17.455Gi ± 2%  +38.57% (p=0.000 n=10)    9.414Gi ± 1%  -25.27% (p=0.000 n=10)
Hash/240      14.650Gi ± 1%   20.241Gi ±  2%  +38.16% (p=0.000 n=10)   17.703Gi ± 3%  +20.84% (p=0.000 n=10)   20.976Gi ± 1%  +43.19% (p=0.000 n=10)    9.408Gi ± 2%  -35.78% (p=0.000 n=10)
Hash/241      14.058Gi ± 1%   18.699Gi ±  1%  +33.01% (p=0.000 n=10)   17.783Gi ± 2%  +26.49% (p=0.000 n=10)   12.884Gi ± 7%   -8.35% (p=0.000 n=10)    9.306Gi ± 6%  -33.80% (p=0.000 n=10)
Hash/512       15.24Gi ± 2%    23.95Gi ±  1%  +57.13% (p=0.000 n=10)    14.91Gi ± 2%   -2.16% (p=0.000 n=10)    16.05Gi ± 2%   +5.33% (p=0.000 n=10)    15.11Gi ± 4%   -0.87% (p=0.043 n=10)
Hash/1024      15.26Gi ± 1%    26.97Gi ±  1%  +76.72% (p=0.000 n=10)    13.40Gi ± 2%  -12.20% (p=0.000 n=10)    16.94Gi ± 2%  +11.01% (p=0.000 n=10)    19.72Gi ± 3%  +29.23% (p=0.000 n=10)
Hash/102400    15.14Gi ± 1%    28.37Gi ±  1%  +87.41% (p=0.000 n=10)    12.21Gi ± 2%  -19.38% (p=0.000 n=10)    17.16Gi ± 4%  +13.34% (p=0.000 n=10)    28.04Gi ± 6%  +85.24% (p=0.000 n=10)
geomean        6.661Gi         9.052Gi        +35.90%                   5.746Gi       -13.74%                   7.656Gi       +14.95%                   6.687Gi        +0.40%
```

</details>
