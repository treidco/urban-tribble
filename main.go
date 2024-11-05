package main

import (
	"fmt"
	"time"
)

type Customer struct {
	id           string
	numPurchases int
}

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
	fmt.Printf("3 for countingsort\n")
	fmt.Scanln(&algo)

	if algo == 1 {
		basicExercise(numItems, max, bubblesort)
	} else if algo == 2 {
		basicExercise(numItems, max, quicksort)
	} else if algo == 3 {
		countingsortExercise(numItems, max)
	} else {
		fmt.Printf("Invalid selection\n")
		return
	}

}

func basicExercise(numItems int, max int, sortFunc func([]int)) {
	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	begin := time.Now().UnixMilli()
	sortFunc(slice)
	end := time.Now().UnixMilli()
	printSlice(slice, 40)
	fmt.Printf("Slice sorted in %d milliseconds\n", end-begin)

	// Verify that it's sorted.
	checkSorted(slice)
}

func countingsortExercise(numItems int, max int) {
	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	begin := time.Now().UnixMilli()

	slice = countingsort(slice, max)

	end := time.Now().UnixMilli()
	printSlice(slice, 40)
	fmt.Printf("Slice sorted in %d milliseconds\n", end-begin)

	// Verify that it's sorted.
	checkSorted(slice)
}
