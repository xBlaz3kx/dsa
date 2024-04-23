package sorting

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	MergeSort(array, 0, len(array))
	fmt.Println(array)
}
