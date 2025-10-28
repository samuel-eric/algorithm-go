package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)

	for {
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

		index, numTests := linearSearch(slice, target)
		if index == -1 {
			fmt.Printf("Target %d not found, %d tests\n", target, numTests)
		} else {
			fmt.Printf("values[%d] = %d, %d tests\n", index, target, numTests)
		}
	}
}
