package main

import (
	"fmt"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max, algo int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	fmt.Printf("Select Sort Algorithm\n")
	fmt.Printf("1 for bubblesort\n")
	fmt.Printf("2 for quicksort\n")
	fmt.Scanln(&algo)

	var sortFunc func([]int)

	if algo == 1 {
		sortFunc = bubblesort
	} else if algo == 2 {
		sortFunc = quicksort
	} else {
		fmt.Printf("Invalid selection\n")
		return
	}

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sortFunc(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
