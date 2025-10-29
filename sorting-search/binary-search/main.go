package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	quicksort(slice)
	printSlice(slice, 40)

	for {
		fmt.Println()

		var input string
		fmt.Printf("Target: ")
		fmt.Scanln(&input)

		if input == "" {
			fmt.Println("Exiting...")
			break
		}

		target, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: input is not integer")
			continue
		}

		fmt.Println()

		fmt.Println("Normal binary search")
		start := time.Now()
		index, numTests := binarySearch(slice, target)
		if index == -1 {
			fmt.Printf("Target %d not found, %d tests\n", target, numTests)
		} else {
			fmt.Printf("values[%d] = %d, %d tests\n", index, target, numTests)
		}
		fmt.Println("Time:", time.Since(start).String())

		fmt.Println()

		fmt.Println("Leftmost binary search")
		start = time.Now()
		index, numTests = binarySearchLeftmost(slice, target)
		if index == -1 {
			fmt.Printf("Target %d not found, %d tests\n", target, numTests)
		} else {
			fmt.Printf("values[%d] = %d, %d tests\n", index, target, numTests)
		}
		fmt.Println("Time:", time.Since(start).String())
	}
}
