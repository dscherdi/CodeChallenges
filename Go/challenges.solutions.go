package main

import "math"

// https://leetcode.com/problems/longest-valid-parentheses/
func longestValidParentheses(str string) int {
	var ps = make([]int, 0)
	var dp = make([]bool, len(str))

	for index, char := range str {
		if char == '(' {
			ps = append(ps, index)
		} else if char == ')' {
			if len(ps) != 0 {
				matchIndex := ps[len(ps)-1]
				ps = ps[:len(ps)-1]
				dp[matchIndex] = true
				dp[index] = true
			}
		}
	}

	max, current := 0, 0
	for i := 0; i < len(dp); i++ {
		if dp[i] {
			current++
			if current > max {
				max = current
			}
		} else {
			current = 0
		}
	}

	return max
}

func longestValidParenthesesDynamic(str string) int {
	var dp = make([]int, len(str))
	maxans := 0
	for i := 1; i < len(str); i++ {
		if str[i] == ')' {
			if str[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2]
				}
				dp[i] += 2
			} else if i-dp[i-1] > 0 && str[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] += dp[i-dp[i-1]-2]
				}
				dp[i] += dp[i-1] + 2
			}
			maxans = MaxInt(maxans, dp[i])
		}
	}
	return maxans
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	var dp = make([]int, len(height))
	var s []int
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(height); i++ {
		if height[i-1] < height[i] && height[i-1] < height[i-2] {
			h := height[i]
			if h > height[i-2] {
				h = height[i-2]
			}
			dp[i] = MaxInt(height[i], height[i-2]) - height[i-1] + dp[i-1]
		} else {
			if height[i] > height[i-1] {

			}
		}

	}

	return dp[len(height)-1]
}

// Utils
func MaxInt(x ...int) int {
	max := math.MinInt
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
}
func MinInt(x ...int) int {
	max := math.MaxInt
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
}
