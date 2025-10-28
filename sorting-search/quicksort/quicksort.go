package main

func partition(slice []int) int {
	low := 0
	hi := len(slice) - 1

	pivot := slice[hi]
	i := low
	for j := 0; j < hi; j++ {
		if slice[j] <= pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}
	slice[i], slice[hi] = slice[hi], slice[i]
	return i
}

func quicksort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	p := partition(slice)

	quicksort(slice[:p])
	quicksort(slice[p+1:])
}
