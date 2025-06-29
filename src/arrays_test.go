package main

import (
	"fmt"
	"testing"
)

func TestNothing(t *testing.T) {
	t.Run("DummyTest", func(t *testing.T) {
		t.Logf("Im a dummy test")
	})
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "butsadbut",
			needle:   "sad",
			want:     3,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Running strStr(%s, %s)", tt.haystack, tt.needle)
		t.Run(testName, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
