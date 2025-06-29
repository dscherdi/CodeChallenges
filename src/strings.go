package main

import (
	"math"
	"unicode"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/
func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}

func firstUniqueChar(s string) int {
	m := make(map[rune]int)

	for _, r := range s {
		m[r]++
	}

	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}

	return -1
}

func myAtoi(s string) int {
	fd := false
	num := 0
	sign := 1
	for _, r := range s {

		if r == ' ' && !fd {
			continue
		} else if r == '-' {
			if fd {
				break
			}
			sign = -1
			fd = true
			continue
		} else if r == '+' {
			if fd {
				break
			}
			sign = 1
			fd = true
			continue
		} else if unicode.IsDigit(r) {
			fd = true
			d := int(r - '0')
			if num == 0 && d == 0 {
				continue
			}
			if num > 0 {
				if sign*(num*10+d) > math.MaxInt32 {
					num = math.MaxInt32
					break
				} else if sign*(num*10+d) < math.MinInt32 {
					num = math.MaxInt32 + 1
					break
				} else {
					num = num*10 + d
				}
			} else {
				num = d
			}
		} else if unicode.IsGraphic(r) {
			break
		}
	}
	res := sign * num

	return res
}
