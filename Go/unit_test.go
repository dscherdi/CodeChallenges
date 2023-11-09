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

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "Test 1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want:     []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			name:     "Test 2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want:     []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			name:     "Test 3",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want:     []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
		{
			name:     "Test 4",
			words:    []string{"Listen", "to", "many,", "speak", "to", "a", "few."},
			maxWidth: 6,
			want:     []string{"Listen", "to", "many,", "speak", "to", "a", "few."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.words, tt.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
