package main

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





// Utils
func MaxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
