package main

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150
func canConstruct(ransomNote string, magazine string) bool {
	tableA := make(map[byte]int)

	for i := range ransomNote {
		tableA[ransomNote[i]]++
	}

	for i := range magazine {
		if _, ok := tableA[magazine[i]]; ok {
			tableA[magazine[i]]--
		}
	}

	for i := range tableA {
		if tableA[i] > 0 {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/isomorphic-strings/?envType=study-plan-v2&envId=top-interview-150
func isIsomorphic(s string, t string) bool {

	mappingST := make(map[byte]byte)
	mappingTS := make(map[byte]byte)

	for i := range s {
		if val, ok := mappingST[s[i]]; ok && val != t[i] {
			return false
		}
		if val, ok := mappingTS[t[i]]; ok && val != s[i] {
			return false
		}
		mappingST[s[i]] = t[i]
		mappingTS[t[i]] = s[i]
	}

	return true
}

// https://leetcode.com/problems/word-pattern/?envType=study-plan-v2&envId=top-interview-150
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	mappingPS := make(map[byte]string)
	mappingSP := make(map[string]byte)

	for i := range pattern {
		if val, ok := mappingPS[pattern[i]]; ok && val != words[i] {
			return false
		}
		if val, ok := mappingSP[words[i]]; ok && val != pattern[i] {
			return false
		}

		mappingPS[pattern[i]] = words[i]
		mappingSP[words[i]] = pattern[i]
	}
	return true
}

// https://leetcode.com/problems/valid-anagram/?envType=study-plan-v2&envId=top-interview-150
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26) // Assuming only lowercase letters

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/group-anagrams/?envType=study-plan-v2&envId=top-interview-150
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		// Convert each string to a rune slice, sort it, and then convert it back to a string
		r := []rune(str)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		sortedStr := string(r)

		// Add the original string to the group of its sorted string
		m[sortedStr] = append(m[sortedStr], str)
	}

	// Convert the map to a slice of slices
	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

// https://leetcode.com/problems/happy-number/?envType=study-plan-v2&envId=top-interview-150
func isHappy(n int) bool {
	m := make(map[int]bool)

	getSumOfSquared := func(n int) int {
		sum := 0
		for n != 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		return sum
	}

	for n != 1 {
		sum := getSumOfSquared(n)

		if _, ok := m[sum]; ok {
			return false
		}

		m[sum] = true

		n = sum

	}

	return true
}

// https://leetcode.com/problems/contains-duplicate-ii/?envType=study-plan-v2&envId=top-interview-150
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	m := make(map[int]bool)

	start := 0
	end := 0
	for ; end <= k && end < len(nums); end++ {
		if _, ok := m[nums[end]]; ok {
			return true
		}

		m[nums[end]] = true
	}

	for ; end < len(nums); end++ {
		m[nums[start]] = false
		if v, ok := m[nums[end]]; ok && v {
			return true
		}
		m[nums[end]] = true
		start++
	}

	return false

}
