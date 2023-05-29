package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
func MinimumSum(num int) int {
	nums := make([]int, 4)

	for i := 0; i < 4; i++ {
		nums[i] = num % 10
		num /= 10
		fmt.Println(nums[i])

	}
	// find min set of pairs

	pairs := make([]int, 2)
	pairs[0] = 10000
	pairs[1] = 10000

	pairs[0] = getMinNonZero(nums)
	pairs[0] = getMinNonZero([]int{})

	return pairs[0] + pairs[1]
}

func getMinNonZero(arr []int) int {
	min := math.MaxInt32

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 && arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
