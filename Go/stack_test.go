package main

import (
	"testing"
)

func TestProcessPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Test 0",
			path: "/a/b///c///../..",
			want: "/a",
		},
		{
			name: "Test 1",
			path: "/a/b/c/../..",
			want: "/a",
		},
		{
			name: "Test 2",
			path: "/a/./b/../../c/",
			want: "/c",
		},
		{
			name: "Test 3",
			path: "/../",
			want: "/",
		},
		{
			name: "Test 4",
			path: "/home/",
			want: "/home",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("processPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"single digit", "1", 1},
		{"simple addition", "1+1", 2},
		{"simple subtraction", "2-1", 1},
		{"complex expression", "(1+2)-3", 0},
		{"expression with spaces", "1 + 2 - 3", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculate(tc.input)
			if result != tc.expected {
				t.Errorf("calculate(%q) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
