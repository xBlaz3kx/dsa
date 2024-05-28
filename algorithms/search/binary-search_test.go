package search

import "testing"

func TestBinarySearch(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}

	index, found := BinarySearch(array, 22)
	if index != 6 || !found {
		t.Error("Expected 6, got ", index)
	}

	_, found = BinarySearch(array, 0)
	if found {
		t.Error("Element should not be found")
	}
}
