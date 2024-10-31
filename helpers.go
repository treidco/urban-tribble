package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	var slice = make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = rand.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if numItems < len(slice) {
		fmt.Println(slice[:numItems])
	} else {
		fmt.Println(slice)
	}
}

func checkSorted(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}
