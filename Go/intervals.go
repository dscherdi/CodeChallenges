package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.com/problems/summary-ranges/?envType=study-plan-v2&envId=top-interview-150
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string
	start := nums[0]

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i] != nums[i+1]-1 {
			if nums[i] == start {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
			}
			if i+1 < len(nums) {
				start = nums[i+1]
			}
		}
	}

	return result
}

// https://leetcode.com/problems/merge-intervals/?envType=study-plan-v2&envId=top-interview-150
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return [][]int{intervals[0]}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	var result [][]int

	currSlice := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= currSlice[1] {
			currSlice[1] = int(math.Max(float64(intervals[i][1]), float64(currSlice[1])))
		} else {
			result = append(result, currSlice)
			currSlice = intervals[i]
		}
	}
	result = append(result, currSlice)
	return result
}
