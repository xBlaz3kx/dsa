package sorting

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
