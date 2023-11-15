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

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/769/
func isValidSudoku(board [][]byte) bool {
	defer timeTrack(time.Now(), "isValidSudoku")
	var row [9][9]bool
	var col [9][9]bool
	var box [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row[i][j] = false
			col[i][j] = false
			box[i][j] = false
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] != '.' {
				num := board[i][j] - '1'
				k := (i/3)*3 + j/3
				if row[i][num] || col[j][num] || box[k][num] {
					return false
				}
				row[i][num] = true
				col[j][num] = true
				box[k][num] = true
			}
		}
	}
	return true
}

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/770/
func rotate(matrix [][]int) {
	defer timeTrack(time.Now(), "rotate image")

	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]         // left -> top
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1] // bottom -> left
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1] // right -> bottom
			matrix[j][n-i-1] = temp                 // top -> right
		}

	}
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

// https://leetcode.com/problems/text-justification/?envType=study-plan-v2&envId=top-interview-150
func fullJustify(words []string, maxWidth int) []string {
	// if a word is 3 chars long and maxWidth is 10 then we need 7 spaces
	// the next word is 4 chars long and we need 6 spaces
	// how do we distribute the spaces?
	// if we use greedy approach, which means we take the solution that is locally optimal with the intent to reach a global optimum
	// how do we choose which solution ?
	// how do we select the local optimum ?
	// greedy implies there are subproblems that are solved by selecting the local optimum choice with the hope that the greater problem reaches the global optimum

	// maxw - len(w)
	// once we do this for each word we know how many spaces each word needs
	// the global optimum is to pack as many words in a line while being fully justified
	// best case we need once space between words
	// when does need to increase the spaces
	// we need to increase the number of spaces until we are fully justified
	// at minimum we need 1 space between words
	// ex: This is an // is the max of words we can fit in line 1
	// we are still missing 6 letters which need tobe distributed left and right
	// how to make it work in an algorithm?
	// how to split the remaining spaces

	result := make([]string, 0)

	count := 0
	lineWords := make([]string, 0)
	for i := 0; i < len(words); {
		if count+len(words[i])+len(lineWords) <= maxWidth {
			count += len(words[i])
			lineWords = append(lineWords, words[i])
			i++
		} else {

			// distribute the number of spaces left
			diff := maxWidth - count
			spots := len(lineWords) - 1
			if spots == 0 {
				lineWords[0] += strings.Repeat(" ", diff)
			} else {
				spacePerSpot := diff / spots
				remainder := diff % spots
				for j := 1; j < len(lineWords); j++ {
					lineWords[j] = strings.Repeat(" ", spacePerSpot) + lineWords[j]

				}
				for j := 1; j < len(lineWords) && remainder > 0; j++ {
					lineWords[j] = " " + lineWords[j]
					remainder--
				}
			}
			result = append(result, strings.Join(lineWords, ""))
			lineWords = make([]string, 0)
			count = 0

		}

	}

	// last line
	if len(lineWords) > 0 {
		lastLine := strings.Join(lineWords, " ")
		lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
		result = append(result, lastLine)
	}
	return result
}
