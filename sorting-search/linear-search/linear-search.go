package main

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	for i, el := range slice {
		numTests++
		if el == target {
			index = i
			return index, numTests
		}
	}
	return -1, numTests
}
