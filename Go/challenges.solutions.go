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

// https://leetcode.com/problems/trapping-rain-water/solution/
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	ans := 0
	var lMax = make([]int, len(height))
	var rMax = make([]int, len(height))
	lMax[0] = height[0]
	rMax[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		lMax[i] = MaxInt(lMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rMax[i] = MaxInt(rMax[i+1], height[i])
	}
	for i := 0; i < len(height); i++ {
		ans += MinInt(lMax[i], rMax[i]) - height[i]
	}
	return ans
}

// https://leetcode.com/problems/counting-bits/
func CountBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 1; i <= n; i++ {
		if i&1 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}

	}
	return dp
}

// Utils
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
