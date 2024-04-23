package sorting

func QuickSort(array []int, low, high int) {
	if low < high {
		pivot := partition(array, low, high)

		QuickSort(array, low, pivot-1)
		QuickSort(array, pivot+1, high)
	}
}

func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
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
