package main

import (
	"fmt"
	"time"
)

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	fmt.Println("Input slice")
	printSlice(slice, 40)
	fmt.Println()

	fmt.Println("Lomuto Partition quicksort")
	start := time.Now()
	quicksort(slice)
	fmt.Println(time.Since(start).String())
	checkSorted(slice)
}
