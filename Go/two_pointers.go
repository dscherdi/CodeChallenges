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
	sort.Ints(nums)
	result := [][]int{}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return result
}

// https://leetcode.com/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=top-interview-150
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := n + 1
	sum := 0
	start := 0

	for end := 0; end < n; end++ {
		sum += nums[end]

		for sum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&envId=top-interview-150
func lengthOfLongestSubstring(s string) int {
	table := make(map[byte]int)

	maxLen := 0
	start := 0

	for end := 0; end < len(s); end++ {
		char := s[end]
		if _, ok := table[char]; ok {
			start = int(math.Max(float64(start), float64(table[char]+1)))
		}
		table[char] = end
		maxLen = int(math.Max(float64(maxLen), float64(end-start+1)))
	}

	return maxLen
}

// https://leetcode.com/problems/minimum-window-substring/?envType=study-plan-v2&envId=top-interview-150
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	if len(t) == 1 {
		if strings.Contains(s, t) {
			return t
		}
		return ""
	}

	table := make(map[byte]int)

	for i := range t {
		table[t[i]]++
	}

	start := 0
	end := 0
	minEnd := 0
	minStart := 0
	minLen := len(s) + 1

	isValidTable := func() bool {
		for _, v := range table {
			if v > 0 {
				return false
			}
		}
		return true
	}

	for ; end < len(s); end++ {
		if _, ok := table[s[end]]; ok {
			table[s[end]]--
		}

		for isValidTable() {
			if end-start+1 < minLen {
				minLen = end - start + 1
				minStart = start
				minEnd = end
			}
			if _, ok := table[s[start]]; ok {
				table[s[start]]++
			}
			start++
		}

	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[minStart:minEnd]
}

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/?envType=study-plan-v2&envId=top-interview-150
func findSubstring(s string, words []string) []int {
	table := make(map[string]int)

	for _, word := range words {
		table[word]++
	}

	result := []int{}
	wordLen := len(words[0])
	wordsLen := len(words)
	totalLen := wordLen * wordsLen

	for i := 0; i < len(s)-totalLen+1; i++ {
		seen := make(map[string]int)
		for j := 0; j < totalLen; j += wordLen {
			word := s[i+j : i+j+wordLen]
			if _, ok := table[word]; ok {
				seen[word]++
				if seen[word] > table[word] {
					break
				}
			} else {
				break
			}
		}
		same := true
		for k, v := range table {
			if seen[k] != v {
				same = false
				break
			}
		}
		if same {
			result = append(result, i)
		}
	}

	return result
}

// https://leetcode.com/problems/game-of-life/?envType=study-plan-v2&envId=top-interview-150
func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			alive := 0
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if k >= 0 && k < len(board) && l >= 0 && l < len(board[i]) && board[k][l]%2 == 1 && (k != i || l != j) {
						alive++
					}

				}
			}
			if board[i][j] == 1 {
				if alive > 3 {
					board[i][j] = 1
				} else if alive == 2 || alive == 3 {
					board[i][j] = 3
				} else {
					board[i][j] = 1
				}
			} else {
				if alive == 3 {
					board[i][j] = 2
				}
			}

		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] /= 2
		}
	}
}
