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
	// Create temporary slices for the two halves
	leftArray := make([]int, middle-left+1)
	rightArray := make([]int, right-middle)

	copy(leftArray, array[left:middle+1])
	copy(rightArray, array[middle+1:right+1])

	i, j, k := 0, 0, left

	// Merge the temporary slices back into the original array
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			array[k] = leftArray[i]
			i++
		} else {
			array[k] = rightArray[j]
			j++
		}
		k++
	}

	// Copy any remaining elements of leftArray
	for i < len(leftArray) {
		array[k] = leftArray[i]
		i++
		k++
	}

	// Copy any remaining elements of rightArray
	for j < len(rightArray) {
		array[k] = rightArray[j]
		j++
		k++
	}
}
