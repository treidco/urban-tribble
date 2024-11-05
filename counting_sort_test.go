package main

import (
	"slices"
	"testing"
)

func TestCountingsort(t *testing.T) {
	input := []int{8, 3, 3, 2, 6, 1, 9, 3, 11, 8, 4, 2, 5, 11}
	expected := []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 8, 8, 9, 11, 11}
	max := 16

	counts := count(input, max)
	counted := addCounts(counts)
	result := sortSlice(input, counted)

	if !slices.Equal(result, expected) {
		t.Errorf("Expected %d, found %d", expected, result)
	}
}

func TestCounting(t *testing.T) {
	slice := []int{8, 3, 3, 2, 6, 1, 9, 3, 11, 8, 4, 2, 5, 11}
	expected := []int{0, 1, 2, 3, 1, 1, 1, 0, 2, 1, 0, 2, 0, 0, 0, 0}
	max := 15

	result := count(slice, max)

	if !slices.Equal(result, expected) {
		t.Errorf("Expected %d, found %d", expected, result)
	}

}

func TestAddCounts(t *testing.T) {
	counts := []int{0, 1, 2, 3, 0, 1, 1, 1}
	expected := []int{0, 1, 3, 6, 6, 7, 8, 9}

	result := addCounts(counts)

	if !slices.Equal(result, expected) {
		t.Errorf("Expected %d, found %d", expected, result)
	}

}

func TestSortSlice(t *testing.T) {
	input := []int{8, 3, 3, 2, 6, 1}
	//counts := []int{0, 1, 1, 2, 0, 0, 1, 0, 1}
	counted := []int{0, 1, 2, 4, 4, 4, 5, 5, 6}
	expected := []int{1, 2, 3, 3, 6, 8}

	result := sortSlice(input, counted)

	if !slices.Equal(result, expected) {
		t.Errorf("Expected %d, found %d", expected, result)
	}

}
