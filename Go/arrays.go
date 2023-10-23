package main

import (
	"math"
	"sort"
	"strings"
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

// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
// Example 1:

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Copy the first m elements of nums1 to a separate slice
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1)

	// Merge the two arrays by comparing elements from nums1Copy and nums2
	i, j := 0, 0
	for k := 0; k < m+n; k++ {
		if i >= m {
			nums1[k] = nums2[j]
			j++
		} else if j >= n {
			nums1[k] = nums1Copy[i]
			i++
		} else if nums1Copy[i] < nums2[j] {
			nums1[k] = nums1Copy[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}

}

// https://leetcode.com/problems/gray-code/
// 01 ^ (01 >> 1) = 01 ^ 00 = 01
// 11 ^ (11 >> 1) = 11 ^ 01 = 10
// 10 ^ (10 >> 1) = 10 ^ 01 = 11
// 100 ^ (100 >> 1) = 100 ^ 10 = 110
func grayCode(n int) []int {
	defer timeTrack(time.Now(), "grayCode")
	res := make([]int, 0)
	for i := 0; i < 1<<uint(n); i++ {
		res = append(res, i^(i>>1))
	}
	return res
}

// https://leetcode.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	defer timeTrack(time.Now(), "Subsets-ii")
	lastSize := 0
	sort.Ints(nums)
	res := [][]int{{}}
	start := 1
	for i := 0; i < len(nums); i++ {
		size := len(res)
		if i > 0 && nums[i] == nums[i-1] {
			start = size - lastSize
		} else {
			start = 0
		}
		lastSize = 0
		for j := start; j < size; j++ {
			subset := make([]int, len(res[j]))
			copy(subset, res[j])
			subset = append(subset, nums[i])
			res = append(res, subset)
			lastSize++
		}
	}
	return res

}

// https://leetcode.com/problems/decode-ways/
// s = '12' => 2
// s = '226' => 3
// s = '0' => 0
// s = '06' => 0

func numDecodings(s string) int {
	defer timeTrack(time.Now(), "numDecodings")
	cs := []byte(s)

	if len(cs) == 0 || cs[0] == '0' {
		return 0
	}

	dp := make([]int, len(cs)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(cs); i++ {
		if cs[i-1] != '0' { // 1-9
			dp[i] += dp[i-1]
		}

		if cs[i-2] == '1' || (cs[i-2] == '2' && cs[i-1] <= '6') { // 10-26
			dp[i] += dp[i-2]

		}
	}

	return dp[len(cs)]

}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4->5->NULL, left = 2, right = 4
// 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	vals := make([]*ListNode, right+1)

	node := head
	for i := 1; i <= right; i++ {
		vals[i] = node
		node = node.Next

	}

	// reverse
	for left < right {
		vals[left].Val, vals[right].Val = vals[right].Val, vals[left].Val
		right--
		left++
	}

	return head
}

// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func removeElement(nums []int, val int) int {

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i

}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	k := 0
	for j := 0; j < len(nums); j++ {
		if j+1 < len(nums) && nums[j] == nums[j+1] {
			continue
		}
		nums[k] = nums[j]
		k++
	}

	return k
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates2(nums []int) int {
	k := 0
	for j := 0; j < len(nums); j++ {
		if j+2 < len(nums) && nums[j] == nums[j+2] {
			continue
		}
		nums[k] = nums[j]
		k++
	}

	return k
}

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	// Boyer-Moore Voting Algorithm
	c := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == c {
			count++
		} else {
			count--
			if count == 0 {
				c = nums[i]
				count = 1
			}
		}
	}
	return c
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit += int(math.Max(float64(prices[i]-prices[i-1]), 0))
	}
	return profit
}

// https://leetcode.com/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150
func canJump(nums []int) bool {

	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}

// https://leetcode.com/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150
func jump2(nums []int) int {
	jumps := 0
	currEnd := 0
	currFarthest := 0
	for i := 0; i < len(nums)-1; i++ {
		currFarthest = int(math.Max(float64(currFarthest), float64(i+nums[i]))) // i = 0, nums[i] = 2, currFarthest = 2; i = 1, nums[i] = 3, currFarthest = 4
		if i == currEnd {                                                       // i = 0, currEnd = 0; i = 1, currEnd = 2; i = 2, currEnd = 2; i = 3, currEnd = 4; i = 4, currEnd = 4; i = 5, currEnd = 4
			jumps++
			currEnd = currFarthest
		}
	}
	return jumps
}

// https://leetcode.com/problems/h-index/?envType=study-plan-v2&envId=top-interview-150
// citations = [3,0,6,1,5]
// Output: 3
// Explanation: there are 3 papers with at least 3 citations
// sol1: sort and find
// sol2: bucket sort

func hIndex(citations []int) int {
	sort.Ints(citations)

	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}

// https://leetcode.com/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-interview-150
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1

	// Calculate the product of all elements to the left of i
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// Calculate the product of all elements to the right of i
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}

// https://leetcode.com/problems/gas-station/?envType=study-plan-v2&envId=top-interview-150
func canCompleteCircuit(gas []int, cost []int) int {
	i := 0

	for i < len(gas) {
		sumOfGas := 0
		sumOfCost := 0
		count := 0
		for count < len(gas) {
			j := (i + count) % len(gas)
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			count++
		}

		if count == len(gas) {
			return i
		} else {
			i = i + count + 1
		}
	}
	return -1
}

// https://leetcode.com/problems/candy/?envType=study-plan-v2&envId=top-interview-150
func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = int(math.Max(float64(candies[i]), float64(candies[i+1]+1)))
		}
	}

	sum := 0

	for _, v := range candies {
		sum += v
	}

	return sum
}

func romanToInt(s string) int {
	c := []byte(s)

	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	sum := 0

	for i := 0; i < len(c); i++ {
		if i+1 < len(c) && c[i] == 'I' && c[i+1] == 'V' {
			c[i] = ' '
			c[i+1] = ' '
			i++
			sum += 4
		} else if i+1 < len(c) && c[i] == 'I' && c[i+1] == 'X' {
			c[i] = ' '
			c[i+1] = ' '
			i++
			sum += 9
		} else if i+1 < len(c) && c[i] == 'X' && c[i+1] == 'L' {
			c[i] = ' '
			c[i+1] = ' '
			i++
			sum += 40
		} else if i+1 < len(c) && c[i] == 'X' && c[i+1] == 'C' {
			c[i] = ' '
			c[i+1] = ' '
			i++
			sum += 90
		} else if i+1 < len(c) && c[i] == 'C' && c[i+1] == 'D' {
			c[i] = ' '
			c[i+1] = ' '
			i++
			sum += 400
		} else if i+1 < len(c) && c[i] == 'C' && c[i+1] == 'M' {
			c[i] = ' '
			c[i+1] = ' '
			i++
			sum += 900
		} else {
			sum += m[c[i]]
		}
	}
	return sum

}

// https://leetcode.com/problems/integer-to-roman/?envType=study-plan-v2&envId=top-interview-150
func intToRoman(num int) string {
	m := make(map[int]string)
	m[1] = "I"
	m[4] = "IV"
	m[5] = "V"
	m[9] = "IX"
	m[10] = "X"
	m[40] = "XL"
	m[50] = "L"
	m[90] = "XC"
	m[100] = "C"
	m[400] = "CD"
	m[500] = "D"
	m[900] = "CM"
	m[1000] = "M"

	res := ""

	for num > 0 {
		if num >= 1000 {
			res += m[1000]
			num -= 1000
		} else if num >= 900 {
			res += m[900]
			num -= 900
		} else if num >= 500 {
			res += m[500]
			num -= 500
		} else if num >= 400 {
			res += m[400]
			num -= 400
		} else if num >= 100 {
			res += m[100]
			num -= 100
		} else if num >= 90 {
			res += m[90]
			num -= 90
		} else if num >= 50 {
			res += m[50]
			num -= 50
		} else if num >= 40 {
			res += m[40]
			num -= 40
		} else if num >= 10 {
			res += m[10]
			num -= 10
		} else if num >= 9 {
			res += m[9]
			num -= 9
		} else if num >= 5 {
			res += m[5]
			num -= 5
		} else if num >= 4 {
			res += m[4]
			num -= 4
		} else if num >= 1 {
			res += m[1]
			num -= 1
		}
	}

	return res
}

// https://leetcode.com/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	c := []byte(s)
	count := 0
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] != ' ' {
			count++
		} else {
			break
		}
	}
	return count
}

// https://leetcode.com/problems/longest-common-prefix/?envType=study-plan-v2&envId=top-interview-150
func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	res := ""

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}

// https://leetcode.com/problems/reverse-words-in-a-string/?envType=study-plan-v2&envId=top-interview-150
func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

// https://leetcode.com/problems/zigzag-conversion/?envType=study-plan-v2&envId=top-interview-150
func convert(s string, numRows int) string {
	chars := []byte(s)

	if numRows == 1 {
		return s
	}

	res := make([][]byte, numRows)

	for i := 0; i < len(chars); {
		for j := 0; j < numRows && i < len(chars); j++ {
			res[j] = append(res[j], chars[i])
			i++
		}

		for j := numRows - 2; j > 0 && i < len(chars); j-- {
			res[j] = append(res[j], chars[i])
			i++
		}
	}

	for i := 1; i < len(res); i++ {
		res[0] = append(res[0], res[i]...)
	}

	return string(res[0])
}

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/?envType=study-plan-v2&envId=top-interview-150
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	if n < m {
		return -1
	}

	// Preprocessing
	f := make([]int, m+1)
	i, j := 0, -1
	f[0] = -1
	for i < m {
		for j > -1 && needle[i] != needle[j] {
			j = f[j]
		}
		i++
		j++
		if i == m || needle[i] != needle[j] {
			f[i] = j
		} else {
			f[i] = f[j]
		}
	}

	// Searching
	i, j = 0, 0
	for i < n {
		for j > -1 && haystack[i] != needle[j] {
			j = f[j]
		}
		i++
		j++
		if j == m {
			return i - m
		}
	}

	return -1
}
