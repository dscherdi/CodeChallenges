package main

import (
	"math"
	"sort"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/?envType=study-plan-v2&envId=top-interview-150
func isPalindrome(s string) bool {
	processedString := ""

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			processedString += strings.ToLower(string(r))
		}
	}

	for i, j := 0, len(processedString)-1; i < j; i, j = i+1, j-1 {
		if processedString[i] != processedString[j] {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan-v2&envId=top-interview-150
func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] > target {
			r--
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
	// we can use binary search in order to find the two elements that add to target
	// starting from substracting the highest element we can find the other element with binary search
	// since we substract the highest value the element we are looking for is smaller than this element
	// specifically if we substract the highest element than the remainder should be the minimum
	// moreover if if we iterate towards n/2 where n is the index of the element els[n] < target than target - n = c and n >= c
	// so how do we iterate ?
	// first we find n which el[n] < target and el[n+1] > target
	// then we define c = target - el[n] and while c + el[n] != target n--
}

// https://leetcode.com/problems/container-with-most-water/?envType=study-plan-v2&envId=top-interview-150
func maxArea(height []int) int {
	// the max area is calculated (right-left) * height[right] where left >= right
	// we can start iterating with 2 pointers left and right and keep track of max value
	// The problem is will this algorithm be correct how to prove it?
	// If we find max to be left = 1 and right = 8 then as we iterate towards the middle the area will be less and less unless there is a huge value height[i] > (right-left) * height[right] but our conditions are that height[i] < 10^4 while n < 10^5, so width will almost always be greater than height
	// do we need to iterate double for each element ?
	// no because the smaller we width the smaller the area, but when left > right or the other way around we only increment the  index of the smaller height
	max := 0
	var left, right int = 0, len(height) - 1
	for left < right {

		leftHeight := height[left]
		rightHeight := height[right]
		min := int(math.Min(float64(leftHeight), float64(rightHeight)))
		area := (right - left) * min

		if (leftHeight >= rightHeight || rightHeight >= leftHeight) && area > max {
			max = area
		}

		if leftHeight < rightHeight {
			left++
		} else if rightHeight < leftHeight {
			right--
		} else {
			left++
		}
	}

	return max
}

// https://leetcode.com/problems/3sum/description/
func threeSum(nums []int) [][]int {
	result := [][]int{}

	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]--
		for j := i + 1; j < len(nums); j++ {
			m[nums[j]]--
			if m[-nums[i]-nums[j]] > 0 {
				result = append(result, []int{nums[i], nums[j], -nums[i] - nums[j]})
			}
			m[nums[j]]++
		}
		m[nums[i]]++
	}

	for i := 0; i < len(result); i++ {
		sort.Ints(result[i])
	}

	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i][0] == result[j][0] && result[i][1] == result[j][1] && result[i][2] == result[j][2] {
				result = append(result[:j], result[j+1:]...)
				j--
			}
		}
	}

	return result

}
