package popcount

import (
	"testing"
)

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// 原版本
/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch2/popcount
BenchmarkPopCount-8             	2000000000	         0.37 ns/op
BenchmarkBitCount-8             	2000000000	         0.68 ns/op
BenchmarkPopCountByClearing-8   	100000000	        16.0 ns/op
BenchmarkPopCountByShifting-8   	30000000	        50.6 ns/op
*/

// Once Do 版本
/*
goos: darwin
goarch: amd64
pkg: github.com/linehk/gopl/ch9/exercise9.2
BenchmarkPopCount-8             	200000000	         7.60 ns/op
BenchmarkBitCount-8             	2000000000	         0.31 ns/op
BenchmarkPopCountByClearing-8   	50000000	        28.3 ns/op
BenchmarkPopCountByShifting-8   	30000000	        49.1 ns/op
*/
