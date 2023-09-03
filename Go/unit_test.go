package main

import (
	"fmt"
	"testing"
)

func RotateArrayTest(t *testing.T, nums []int, exp []int, k int) {

	if fmt.Sprint(nums) != fmt.Sprint(exp) {
		t.Errorf("is %v, should be %v", nums, exp)
	}
}
