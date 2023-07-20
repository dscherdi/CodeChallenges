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

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
func ContainsDuplicate(nums []int) bool {
	defer timeTrack(time.Now(), "ContainsDuplicate")
	m := make(map[int]bool)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}

	return false
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
func SingleNumber(nums []int) int {
	defer timeTrack(time.Now(), "SingleNumber")
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/
func Intersect(nums1 []int, nums2 []int) []int {
	defer timeTrack(time.Now(), "Intersect")
	m := make(map[int]int)
	var res []int

	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}

	return res

}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/
func PlusOne(digits []int) []int {

	defer timeTrack(time.Now(), "PlusOne")
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++

			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)

}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/567/
func MoveZeroes(nums []int) {
	defer timeTrack(time.Now(), "MoveZeroes")

	m := make(map[int]int)
	c := 0
	for i, v := range nums {
		if v == 0 {
			c++
		} else {
			m[i] = c
		}

	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && m[i] != 0 {
			nums[i-m[i]] = nums[i]
			nums[i] = 0
		}
	}

}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/
func TwoSum(nums []int, target int) []int {
	defer timeTrack(time.Now(), "TwoSum")

	d := make(map[int]int)
	v := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		d[i] = target - nums[i]
		v[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if v[d[i]] != 0 && i != v[d[i]] {
			return []int{i, v[d[i]]}
		}
	}
	return []int{}
}
