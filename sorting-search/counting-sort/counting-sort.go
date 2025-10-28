package main

func countingSort(slice []Customer, max int) []Customer {
	counts := make([]int, max)
	sorted := make([]Customer, len(slice))

	for _, customer := range slice {
		counts[customer.numPurchases] += 1
	}

	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i] + counts[i-1]
	}

	for j := len(slice) - 1; j >= 0; j-- {
		counts[slice[j].numPurchases] -= 1
		sorted[counts[slice[j].numPurchases]] = slice[j]
	}

	return sorted
}
