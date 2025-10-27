package main

func bubbleSort(slice []int) {
	for {
		swapped := false
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
