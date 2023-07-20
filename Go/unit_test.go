package main

import (
	"fmt"
	"testing"
)

func RotateArrayTest(t *testing.T, nums []int, exp []int, k int) {
	RotateArray(nums, k)

	if fmt.Sprint(nums) != fmt.Sprint(exp) {
		t.Errorf("is %v, should be %v", nums, exp)
	}
}

func TestRotateArray(t *testing.T) {
	RotateArrayTest(t, []int{-1, 100, 3, 99}, []int{3, 99, -1, 100}, 2)
	RotateArrayTest(t, []int{1, 2, 3, 4, 5, 6, 7}, []int{5, 6, 7, 1, 2, 3, 4}, 3)
}
