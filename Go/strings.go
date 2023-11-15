package main

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
