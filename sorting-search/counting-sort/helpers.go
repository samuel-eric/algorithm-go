package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id           string
	numPurchases int
}

// Make a slice containing Customer struct with random numPurchases field.
func makeRandomSlice(numItems, max int) []Customer {
	slice := make([]Customer, numItems)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		slice[i] = Customer{
			id:           fmt.Sprintf("C%d", i+1),
			numPurchases: random.Intn(max),
		}
	}
	return slice
}

// Print at most numItems items.
func printSlice(slice []Customer, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []Customer) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}

	fmt.Println("The slice is sorted")
}
