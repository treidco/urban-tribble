package main

import (
	"slices"
	"testing"
)

func TestCustomerCountingsort(t *testing.T) {
	input := makeCustomersFromInts([]int{8, 3, 3, 2, 6, 1, 9, 3, 11, 8, 4, 2, 5, 11})
	expected := []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 8, 8, 9, 11, 11}
	max := 16

	counts := countCustomer(input, max)
	counted := addCounts(counts)
	result := sortCustomerSlice(input, counted)
	resultInts := makeIntsFromCustomers(result)

	if !slices.Equal(resultInts, expected) {
		t.Errorf("Expected %v, found %v", expected, result)
	}
}

func TestCustomerCounting(t *testing.T) {
	intSlice := []int{8, 3, 3, 2, 6, 1, 9, 3, 11, 8, 4, 2, 5, 11}
	expected := []int{0, 1, 2, 3, 1, 1, 1, 0, 2, 1, 0, 2, 0, 0, 0, 0}
	max := 15

	slice := makeCustomersFromInts(intSlice)

	result := countCustomer(slice, max)

	if !slices.Equal(result, expected) {
		t.Errorf("Expected %d, found %d", expected, result)
	}

}
