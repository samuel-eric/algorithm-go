package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	start := time.Now()
	sorted := countingSort(slice, max)
	fmt.Println("Time:", time.Since(start).String())
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
