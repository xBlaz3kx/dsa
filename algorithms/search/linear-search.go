package search

// LinearSearch searches for a number in an array using the linear search algorithm
func linearSearch(array []int, element int) (index int, found bool) {
	for i, v := range array {
		if v == element {
			return i, true
		}
	}

	return -1, false
}
