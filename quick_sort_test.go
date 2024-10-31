package main

import (
	"slices"
	"testing"
)

func TestPartition(t *testing.T) {
	slice := []int{8, 3, 7, 2, 6}
	expected := []int{3, 2, 6, 8, 7}
	expectedPartition := 2
	pivot := partition(slice)

	if pivot != expectedPartition {
		t.Errorf("Expected %d, found %d", expectedPartition, pivot)
	}
	if !slices.Equal(slice, expected) {
		t.Errorf("Slice not partitioned as expected")
	}

}

func TestPartitionWithSmallArray(t *testing.T) {
	slice := []int{3, 2, 7}
	expected := []int{3, 2, 7}
	expectedPartition := 2
	pivot := partition(slice)

	if pivot != expectedPartition {
		t.Errorf("Expected %d, found %d", expectedPartition, pivot)
	}
	if !slices.Equal(slice, expected) {
		t.Errorf("Slice not partitioned as expected, %d", slice)
	}
}

func TestPartitionWithTwoElementArray(t *testing.T) {
	slice := []int{3, 2}
	expected := []int{2, 3}
	expectedPartition := 0
	pivot := partition(slice)

	if pivot != expectedPartition {
		t.Errorf("Expected %d, found %d", expectedPartition, pivot)
	}
	if !slices.Equal(slice, expected) {
		t.Errorf("Slice not partitioned as expected, %d", slice)
	}
}

func TestQuicksort(t *testing.T) {
	slice := []int{8, 3, 7, 2, 6}
	expected := []int{2, 3, 6, 7, 8}

	quicksort(slice)

	if !slices.Equal(slice, expected) {
		t.Errorf("Quicksort failed, result was %d", slice)
	}
}

func TestQuicksortWithSmallArray(t *testing.T) {
	slice := []int{3, 2, 7}
	expected := []int{2, 3, 7}

	quicksort(slice)

	if !slices.Equal(slice, expected) {
		t.Errorf("Quicksort failed, result was %d", slice)
	}
}

func TestQuicksortWithTwoElementArray(t *testing.T) {
	slice := []int{3, 2}
	expected := []int{2, 3}

	quicksort(slice)

	if !slices.Equal(slice, expected) {
		t.Errorf("Quicksort failed, result was %d", slice)
	}
}
