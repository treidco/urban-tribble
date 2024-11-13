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
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSliceOfCustomers(numItems, max)
	printCustomerSlice(slice, 40)
	fmt.Println()

	begin := time.Now().UnixMilli()
	sortedSlice := countingsortCustomer(slice, max)
	end := time.Now().UnixMilli()
	printCustomerSlice(sortedSlice, 40)
	fmt.Printf("Slice sorted in %d milliseconds\n", end-begin)

	checkCustomersSorted(sortedSlice)
}
