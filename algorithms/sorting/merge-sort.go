package sorting

// MergeSort sorts an array using the merge sort algorithm. It is a divide and conquer algorithm
// that divides the array into two halves, sorts the two halves and then merges them. This process is repeated recursively
// until the array is sorted.
func MergeSort(array []int, left, right int) {
	if left < right {

		middle := (left + right) / 2

		// Divide the array into two halves and sort them
		MergeSort(array, left, middle)
		MergeSort(array, middle+1, right)

		// Merge the results
		mergeSortMerge(array, left, right, middle)
	}
}

func mergeSortMerge(array []int, left, right, middle int) {

	// todo
	for i := left; i <= right; i++ {

	}

}
