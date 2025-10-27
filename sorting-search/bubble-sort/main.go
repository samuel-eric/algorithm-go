package main

import "fmt"

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	bubbleSort(slice)
	printSlice(slice, 40)

	checkSorted(slice)
}
