package main

func countingsort(input []int, max int) []int {
	return sortSlice(input, addCounts(count(input, max)))
}

func count(input []int, max int) []int {
	output := make([]int, max+1)
	for i := 0; i < len(input); i++ {
		output[input[i]] += 1
	}
	return output
}

func addCounts(input []int) []int {
	for i := 1; i < len(input); i++ {
		input[i] = input[i] + input[i-1]
	}
	return input
}

func sortSlice(input []int, counted []int) []int {
	output := make([]int, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		output[counted[input[i]]-1] = input[i]
		counted[input[i]]--
	}

	return output
}
