package main

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "   +0 123",
			want: 0,
		},
		{
			s:    "-91283472332",
			want: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
