package main

import (
	"math"
	"strings"
)

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

// https://leetcode.com/problems/set-matrix-zeroes/?envType=study-plan-v2&envId=top-interview-150
func setZeroes(matrix [][]int) {
	rowFlag := false
	colFlag := false

	// check if first row or first column has a zero
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && matrix[i][j] == 0 {
				rowFlag = true
			}
			if j == 0 && matrix[i][j] == 0 {
				colFlag = true
			}
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// use the first row and column as markers
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// set first row to zero if needed
	if rowFlag {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	// set first column to zero if needed
	if colFlag {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// https://leetcode.com/problems/spiral-matrix/?envType=study-plan-v2&envId=top-interview-150
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	leftToRight, rightToLeft, topToBottom, bottomToTop := "left-to-right", "right-to-left", "top-to-bottom", "bottom-to-top"
	dirMap := make(map[string][]int)

	dirMap[leftToRight] = []int{0, 1}
	dirMap[rightToLeft] = []int{0, -1}
	dirMap[topToBottom] = []int{1, 0}
	dirMap[bottomToTop] = []int{-1, 0}

	totalElements := len(matrix) * len(matrix[0])
	x, y := 0, -1 // Start outside the matrix
	rowLen, colLen := len(matrix), len(matrix[0])

	currDirection := leftToRight

	result := make([]int, 0)

	leftBound, rightBound, topBound, bottomBound := 0, colLen-1, 0, rowLen-1

	for len(result) < totalElements {
		newX, newY := x+dirMap[currDirection][0], y+dirMap[currDirection][1]

		if currDirection == leftToRight && newY > rightBound {
			currDirection = topToBottom
			topBound++
		} else if currDirection == rightToLeft && newY < leftBound {
			currDirection = bottomToTop
			bottomBound--
		} else if currDirection == topToBottom && newX > bottomBound {
			currDirection = rightToLeft
			rightBound--
		} else if currDirection == bottomToTop && newX < topBound {
			currDirection = leftToRight
			leftBound++
		} else {
			x, y = newX, newY
			result = append(result, matrix[x][y])
		}
	}

	return result
}
