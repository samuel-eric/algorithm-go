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

	slice1 := makeRandomSlice(numItems, max)
	fmt.Println("Input slice")
	printSlice(slice1, 40)
	fmt.Println()

	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	fmt.Println("Normal bubble sort")
	start := time.Now()
	bubbleSort(slice1)
	checkSorted(slice1)
	fmt.Println(time.Since(start).String())

	fmt.Println()

	fmt.Println("Optimized bubble sort")
	start = time.Now()
	optimizedBubbleSort(slice2)
	checkSorted(slice2)
	fmt.Println(time.Since(start).String())
}
