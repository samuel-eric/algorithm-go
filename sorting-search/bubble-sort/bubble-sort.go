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

func optimizedBubbleSort(slice []int) {
	n := len(slice)

	for n >= 1 {
		newn := 0
		for i := 1; i < n; i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				newn = i
			}
		}
		n = newn
	}
}

func cocktailShakerSort(slice []int) {
	swapped := true
	for swapped {
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

		swapped = false
		for i := len(slice) - 1; i > 0; i-- {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				swapped = true
			}
		}
	}
}
