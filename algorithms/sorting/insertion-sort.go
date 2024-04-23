package sorting

// InsertionSort sorts an array using the insertion sort algorithm
func InsertionSort(array []int) {

	// Start at first element
	for i := 1; i < len(array)-1; i++ {
		temp := array[i]
		j := i - 1

		// Check the already sorted part of array and insert the element there
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = temp
	}
}
