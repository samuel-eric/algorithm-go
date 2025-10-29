package main

// Perform binary search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearch(slice []int, target int) (index, numTests int) {
	l := 0
	r := len(slice) - 1

	for l <= r {
		numTests++
		m := l + (r-l)/2
		if slice[m] < target {
			l = m + 1
		} else if slice[m] > target {
			r = m - 1
		} else {
			return m, numTests
		}
	}
	return -1, numTests
}

// Perform binary search.
// Return the target's leftmost location in the slice (if there is duplicate) and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearchLeftmost(slice []int, target int) (index, numTests int) {
	l := 0
	r := len(slice)

	for l < r {
		numTests++
		m := l + (r-l)/2
		if slice[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if slice[l] == target {
		return l, numTests
	} else {
		return -1, numTests
	}
}
