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
func MaxProfitArray(prices []int) {
	defer timeTrack(time.Now(), "maxProfit")
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
// [1,2,3,4,5,6,7] => [5,6,7,1,2,3,4]
// 1 step => [7,1,2,3,4,5,6] RotatteOnce
// 2 step => [6,7,1,2,3,4,5] RotateOnce
// 3 step => [5,6,7,1,2,3,4] RotateOnce
// O(n^2)
// Opt: Jump
func RotateArray(nums []int, k int) {
	defer timeTrack(time.Now(), "rotateArray")
	k = k % len(nums)

	for i := 0; i < k; i++ {
		c := nums[len(nums)-1]
		prev, curr := nums[0], nums[0]

		// single position rotation
		for j := 1; j < len(nums); j++ {
			curr = nums[j]
			nums[j] = prev
			prev = curr
		}
		nums[0] = c
	}
}
