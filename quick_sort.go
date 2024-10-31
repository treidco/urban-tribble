package main

func quicksort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	pivot := partition(slice)
	quicksort(slice[0:pivot])
	quicksort(slice[pivot+1:])
}

func partition(slice []int) int {
	pivot := slice[len(slice)-1]
	// temp pivot index
	i := 0
	for j := 0; j <= len(slice)-2; j++ {
		if slice[j] <= pivot {
			swap := slice[i]
			slice[i] = slice[j]
			slice[j] = swap
			i++
		}
	}
	// swap the element at the pivot index and the last element
	hi := slice[len(slice)-1]
	slice[len(slice)-1] = slice[i]
	slice[i] = hi

	return i // the pivot index
}
