package main

func countingsortCustomer(input []Customer, max int) []Customer {
	return sortCustomerSlice(input, addCounts(countCustomer(input, max)))
}

func countCustomer(input []Customer, max int) []int {
	output := make([]int, max+1)
	for i := 0; i < len(input); i++ {
		output[input[i].numPurchases] += 1
	}
	return output
}

func sortCustomerSlice(input []Customer, counted []int) []Customer {
	output := make([]Customer, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		output[counted[input[i].numPurchases]-1] = input[i]
		counted[input[i].numPurchases]--
	}
	return output
}
