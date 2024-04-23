package sorting

// MergeSort sorts an array using the merge sort algorithm. It is a divide and conquer algorithm
// that divides the array into two halves, sorts the two halves and then merges them. Repeat this
// process until the array is sorted.
func MergeSort(array []int, left, right int) {
	if left < right {

		middle := (left + right) / 2
		MergeSort(array, left, middle)
		MergeSort(array, middle+1, right)

		mergeSortMerge(array, left, right, middle)
	}
}

func mergeSortMerge(array []int, left, right, middle int) {
	for i := left; i <= right; i++ {

	}
}
