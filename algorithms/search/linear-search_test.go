package search

import "testing"

func TestLinearSearch(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	index, found := LinearSearch(array, 22)
	if !found || index != 5 {
		t.Errorf("Expected to find element")
	}

	index, found = LinearSearch(array, 1)
	if !found || index != 4 {
		t.Errorf("Expected to find element")
	}

	_, found = LinearSearch(array, 0)
	if found {
		t.Errorf("Element should not be present")
	}
}
