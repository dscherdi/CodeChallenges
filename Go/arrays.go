package main

import (
	"time"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
func RemoveDuplicates(nums []int) int {
	defer timeTrack(time.Now(), "removeDuplicates")
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	nums = nums[:i+1]

	return i + 1
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/
func MaxProfit(prices []int) {
	defer timeTrack(time.Now(), "maxProfit")
}
