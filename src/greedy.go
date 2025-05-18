package main

import (
	"sort"
	"time"
)

// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
func MinimumSum(num int) int {
	defer timeTrack(time.Now(), "MinimumSum")
	nums := make([]int, 4)

	for i := 0; i < 4; i++ {
		nums[i] = num % 10
		num = num / 10
	}

	sort.Ints(nums)
	pair1 := 0
	pair2 := 0
	i := 0
	for ; i < 4; i++ {
		if nums[i] != 0 {
			break
		}
	}

	pair1 = nums[i]

	if i+1 < 4 {
		pair2 = nums[i+1]
	}

	if i+2 < 4 {

		pair1 = pair1*10 + nums[i+2]
	}
	if i+3 < 4 {
		pair2 = pair2*10 + nums[i+3]
	}

	return pair1 + pair2
}
