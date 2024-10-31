package main

import "testing"

func TestMakeRandomSlice(t *testing.T) {

	numItems := 4
	maxSize := 8
	result := makeRandomSlice(numItems, maxSize)

	if len(result) != 4 {
		t.Errorf("Expected length: %d, got: %d ", numItems, len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] > maxSize {
			t.Errorf("Expected value less than %d, got: %d ", maxSize, result[i])
		}
	}
}
