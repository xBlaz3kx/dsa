package sorting

import (
	"fmt"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	InsertionSort(array)
	fmt.Println(array)
}
