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
