package sorting

// BubbleSort sorts an array using two nested loops to compare and swap elements
func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j <= len(array)-1; j++ {
			if array[i] >= array[j] {
				// Swap
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}
}
