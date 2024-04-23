package search

import (
	"log"
	"sort"
)

// BinarySearch searches for a number in a sorted array using the binary search algorithm
func BinarySearch(array []int, num int) (index int, found bool) {

	// Ensure array is sorted
	sort.Ints(array)

	i := 0
	size := len(array)

	for {
		i = (i+size)/2 - 1

		log.Printf("i: %d", i)

		// Check if number matches
		if num == array[i] {
			return i, true
		}

		if i == len(array)-1 || i == 0 {
			return -1, false
		}

		if num > array[i] {
			i = i + 1
		} else {
			size = i
		}
	}
}
