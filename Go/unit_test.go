package main

import (
	"fmt"
	"reflect"
	"testing"
)

func RotateArrayTest(t *testing.T, nums []int, exp []int, k int) {

	if fmt.Sprint(nums) != fmt.Sprint(exp) {
		t.Errorf("is %v, should be %v", nums, exp)
	}
}

func TestRotateArray(t *testing.T) {
	RotateArrayTest(t, []int{-1, 100, 3, 99}, []int{3, 99, -1, 100}, 2)
	RotateArrayTest(t, []int{1, 2, 3, 4, 5, 6, 7}, []int{5, 6, 7, 1, 2, 3, 4}, 3)
}

func TestContainsDuplicate(t *testing.T) {
	r := ContainsDuplicate([]int{1, 2, 3, 1})

	if r != true {
		t.Errorf("is %v, should be %v", r, true)
	}
}

func TestSingleNumber(t *testing.T) {
	r := SingleNumber([]int{1, 2, 2, 3, 1})

	if r != 3 {
		t.Errorf("is %v, should be %v", r, true)
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 2, 3, 1}
	exp := []int{1, 2, 3, 1, 0, 0}
	MoveZeroes(nums)

	if fmt.Sprint(nums) != fmt.Sprint(exp) {
		t.Errorf("is %v, should be %v", nums, exp)
	}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	exp := []int{0, 1}
	r := TwoSum(nums, 9)

	if fmt.Sprint(r) != fmt.Sprint(exp) {
		t.Errorf("is %v, should be %v", r, exp)
	}
}
func TestTwoSum2(t *testing.T) {
	nums := []int{1, 3, 4, 2}
	exp := []int{2, 3}
	r := TwoSum(nums, 6)

	if fmt.Sprint(r) != fmt.Sprint(exp) {
		t.Errorf("is %v, should be %v", r, exp)
	}
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)

	expected := []int{1, 2, 2, 3, 5, 6}
	if !reflect.DeepEqual(nums1, expected) {
		t.Errorf("merge failed: got %v, expected %v", nums1, expected)
	}
}

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"06", 0},
		{"10", 1},
		{"27", 1},
		{"111111111111111111111111111111111111111111111", 1836311903},
	}

	for _, c := range cases {
		got := numDecodings(c.input)
		if got != c.want {
			t.Errorf("numDecodings(%q) == %d, want %d", c.input, got, c.want)
		}
	}
}
