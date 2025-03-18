package sorting

// InsertionSort sorts an array using the insertion sort algorithm
func InsertionSort(array []int) {
	// Start at first element
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i - 1

		for j = i - 1; j >= 0 && temp < array[j]; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}
