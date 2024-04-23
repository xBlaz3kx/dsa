package main

import "fmt"

func sortingMain() {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	bubbleSort(array)
	fmt.Println(array)

	array = []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	insertionSort(array)
	fmt.Println(array)

	array = []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	mergeSort(array, 0, len(array))
	fmt.Println(array)

	array = []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	// quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func insertionSort(array []int) {
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

func mergeSort(array []int, left, right int) {
	if left < right {

		middle := left + (left-right)/2
		mergeSort(array, left, middle)
		mergeSort(array, middle+1, right)

		mergeSortMerge(array, left, right, middle)
	}
}

func mergeSortMerge(array []int, left, right, middle int) {

}

func quickSort(array []int, low, high int) {
	if low < high {
		pivot := partition(array, low, high)

		quickSort(array, low, pivot-1)
		quickSort(array, pivot+1, high)
	}
}

func swap(array []int, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
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

func bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j <= len(array)-1; j++ {
			if array[i] >= array[j] {
				// Swap
				fmt.Println(array[i], array[j])
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}
}
