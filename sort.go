package main

func bubbleSort(slice []int) {
	swap := false
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			swap = true
			sink := slice[i]
			slice[i] = slice[i+1]
			slice[i+1] = sink
		}
	}
	if swap == true {
		bubbleSort(slice)
	}
}
