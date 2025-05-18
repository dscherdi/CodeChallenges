package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s | ", name, elapsed)
}

func FactorialUInt64(n uint64) uint64 {
	var p uint64 = 1
	for i := uint64(1); i <= n; i++ {
		p *= i
	}
	return p
}

func FactorialBigInt(n int) *big.Int {
	p := big.NewInt(1)

	for i := int64(1); i <= int64(n); i++ {
		p = p.Mul(p, big.NewInt(i))
	}
	return p
}
func MaxInt(x ...int) int {
	max := math.MinInt
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}
func MinInt(x ...int) int {
	min := math.MaxInt
	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}
func MaxIntIndex(x ...int) int {
	max := math.MinInt
	index := -1
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
			index = i
		}
	}
	return index
}
func MinIntIndex(x ...int) int {
	min := math.MaxInt
	index := -1
	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
			index = i
		}
	}
	return index
}
