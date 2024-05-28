package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	MergeSort(array, 0, len(array))

	if !reflect.DeepEqual([]int{1, 2, 3, 6, 11, 15, 22, 31, 50}, array) {
		t.Errorf("Expected [1, 2, 3, 6, 11, 15, 22, 31, 50], got %v", array)
	}
}
