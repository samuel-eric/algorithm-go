package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	slice := make([]int, numItems)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		slice[i] = random.Intn(max)
	}
	return slice
}

// Print at most numItems items.
func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}
