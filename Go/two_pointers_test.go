package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Test 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Test 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Test 3",
			s:    "0P",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "Test 1",
			numbers: []int{5, 25, 75},
			target:  100,
			want:    []int{2, 3},
		},
		{
			name:    "Test 1",
			numbers: []int{1, 2, 5, 7, 8, 11, 15},
			target:  9,
			want:    []int{1, 5},
		},
		{
			name:    "Test 2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "Test 3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "Test 4",
			numbers: []int{-3, 3, 4, 90},
			target:  0,
			want:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Test 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Test 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "Test 3",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "Test 4",
			height: []int{1, 2, 1},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
