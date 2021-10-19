# 🚀 Faster-random

Goroutine-safe, Faster pseudo-random number and string generator.

- Non-goroutine-safe: `var Random = rand.New(rand.NewSource(time.Now().UnixNano()))`
- Goroutine-safe: `var Random = NewRand()`, and has super high performance. See: [Benchmarks](#-benchmarks).

## ⚙️ Installation

```go
go get -u github.com/fufuok/random
```

## ⚡️ Quickstart

```go
package main

import (
	"fmt"

	"github.com/fufuok/random"
)

func main() {
	// random.Rand is a goroutine-safe rand.New() instance
	// FastIntn similar to rand.Intn or random.Rand.Intn, but faster.
	for i := 1; i <= 1000000; i *= 10 {
		fmt.Print(random.Rand.Intn(i), random.FastIntn(i), " ")
	}

	fmt.Println()

	// Fast random string generator
	fmt.Println(random.RandString(20), random.RandString(20))

	// Instantiate *rand.Rand with the current time by default
	rd := random.NewRand()
	fmt.Println(rd.Int63())

	// Returns a new pseudo-random generator seeded with the given value.
	rd = random.NewRand(1)
	fmt.Println(rd.Int63())
}
```

## 📚 Examples

please see: [examples](examples)

```go
package random // import "github.com/fufuok/random"

var Rand = NewRand()
func FastIntn(n int) int
func FastRand() uint32
func FastRandBytes(n int) []byte
func FastRandn(n uint32) uint32
func NewRand(seed ...int64) *rand.Rand
func RandBytes(n int) []byte
func RandHex(nHalf int) string
func RandInt(min, max int) int
func RandString(n int) string
func RandUint32(min, max uint32) uint32
```

## 🤖 Benchmarks

```shell
go test -run=^$ -benchmem -benchtime=1s -count=2 -bench=.
goos: linux
goarch: amd64
pkg: github.com/fufuok/random
cpu: Intel(R) Xeon(R) Gold 6151 CPU @ 3.00GHz
BenchmarkRandInt/RandInt-4                     300554511                3.982 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/RandInt-4                     299879216                4.007 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/RandUint32-4                  266118160                4.466 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/RandUint32-4                  266337582                4.507 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/FastIntn-4                    313761758                3.834 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/FastIntn-4                    312968804                3.811 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/Rand.Intn-4                    35001715                34.30 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/Rand.Intn-4                    34904052                34.54 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/std.rand.Intn-4                56418733                21.57 ns/op            0 B/op          0 allocs/op
BenchmarkRandInt/std.rand.Intn-4                56331698                21.41 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/RandInt-4            1000000000                1.060 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/RandInt-4            1000000000                1.045 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/RandUint32-4          990860647                1.197 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/RandUint32-4         1000000000                1.182 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/FastIntn-4           1000000000                1.060 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/FastIntn-4           1000000000                1.055 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/Rand.Intn-4           130758892                9.132 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/Rand.Intn-4           130173494                9.065 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/std.rand.Intn-4        13878208                88.73 ns/op            0 B/op          0 allocs/op
BenchmarkRandIntParallel/std.rand.Intn-4        13828624                89.97 ns/op            0 B/op          0 allocs/op

BenchmarkRandBytes/RandString-4                 18142927                64.27 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytes/RandString-4                 18944168                64.43 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytes/RandBytes-4                   1730853                694.3 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytes/RandBytes-4                   1719566                687.4 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytes/FastRandBytes-4              18185881                64.99 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytes/FastRandBytes-4              18052567                65.73 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytesParallel/RandString-4         63093751                19.00 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytesParallel/RandString-4         67928510                19.29 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytesParallel/RandBytes-4           1309642                916.2 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytesParallel/RandBytes-4           1315711                916.6 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytesParallel/FastRandBytes-4      64837544                20.05 ns/op           24 B/op          1 allocs/op
BenchmarkRandBytesParallel/FastRandBytes-4      65973478                19.52 ns/op           24 B/op          1 allocs/op
```





*ff*