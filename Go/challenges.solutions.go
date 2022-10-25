package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s | ", name, elapsed)
}

// https://leetcode.com/problems/longest-valid-parentheses/
func LongestValidParentheses(str string) int {
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

func LongestValidParenthesesDynamic(str string) int {
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
func Trap(height []int) int {
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

// https://leetcode.com/problems/fibonacci-number/
func Fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// https://leetcode.com/problems/pascals-triangle/
func PascalsTriangle(numRows int) [][]int {
	if numRows < 2 {
		return [][]int{{1}}
	}
	dp := make([][]int, numRows)
	dp[0] = []int{1}
	dp[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j <= i/2; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			dp[i][i-j] = dp[i][j]
		}
	}

	return dp
}

// https://leetcode.com/problems/n-th-tribonacci-number/
func Tribonacci(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

// https://leetcode.com/problems/min-cost-climbing-stairs/
func MinCostClimbingStairs(cost []int) int {
	defer timeTrack(time.Now(), "MinCostClimbingStair1")
	dp := make([]int, len(cost)+2)
	dp[len(cost)] = 0
	dp[len(cost)+1] = 0
	for i := len(cost) - 1; i > -1; i-- {
		dp[i] = cost[i] + MinInt(dp[i+1], dp[i+2])
	}
	return MinInt(dp[0], dp[1])
}
func MinCostClimbingStairs2(cost []int) int {
	defer timeTrack(time.Now(), "MinCostClimbingStair2")
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		c := cost[i] + MinInt(a, b)
		a = b
		b = c
	}
	return MinInt(a, b)
}

// https://leetcode.com/problems/pascals-triangle-ii/
func PascalsTriangle2(n int) []int {
	defer timeTrack(time.Now(), "PascalsTriangle2")
	if n == 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{1, 1}
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[n] = 1
	fn := FactorialBigInt(n)
	for i := 1; i <= n/2; i++ {
		r := FactorialBigInt(i)
		qr := FactorialBigInt(n - i)
		rec := big.NewInt(1)
		cqr := rec.Div(fn, r.Mul(r, qr))
		dp[i] = int(cqr.Uint64())
		dp[n-i] = dp[i]
	}

	return dp
}

func PascalsTriangle2_2(n int) []int {
	defer timeTrack(time.Now(), "PascalsTriangle2_2")
	if n == 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{1, 1}
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[n] = 1
	up := n
	down := 1
	for i := 1; i <= n/2; i++ {
		dp[i] = dp[i-1] * up / down
		up--
		down++
		dp[n-i] = dp[i]
	}
	return dp
}

// Utils
func FactorialUInt64(n uint64) uint64 {
	var p uint64 = 1
	for i := uint64(1); i <= n; i++ {
		p *= i
	}
	return p
}

func FactorialBigInt(n int) *big.Int {
	p := big.NewInt(1)

	for i := int64(1); i <= int64(n); i++ {
		p = p.Mul(p, big.NewInt(i))
	}
	return p
}
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
