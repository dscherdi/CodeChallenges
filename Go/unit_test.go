package main

import (
	"bytes"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(input)
	if !bytes.Equal(input, expected) {
		t.Errorf("reverseString(%q) == %q, expected %q", string(input), string(input), string(expected))
	}
}
