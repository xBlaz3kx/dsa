package sorting

// Sorts an array using the quick sort algorithm. The array is sorted based on the pivot, which is the last element in the array.
// Elements of the array, smaller than the pivot, are moved to the left subarray, while elements greater than the pivot are moved to the right subarray.
// The left and right subarrays are then sorted recursively, untill the entire array is sorted.
func QuickSort(array []int, low, high int) {
	if low < high {
		pivot := partition(array, low, high)

		// Sort the left side of the pivot
		QuickSort(array, low, pivot-1)

		// Sort the right side of the pivot
		QuickSort(array, pivot+1, high)
	}
}

func partition(array []int, low, high int) int {
	// Pivot is the last element in the array
	pivot := array[high]

	// i is the index of the smaller element
	i := low - 1

	for j := low; j <= high-1; j++ {
		// If current element is smaller than the pivot, swap the pivot with the current element
		if array[j] < pivot {
			i++
			swap(array, i, j)
		}
	}

	swap(array, i+1, high)
	return i + 1
}

func swap(array []int, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
