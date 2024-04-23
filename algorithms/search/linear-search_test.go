package search

import "testing"

func TestLinearSearch(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}
	index, found := linearSearch(array, 22)
	if !found || index != 5 {
		t.Errorf("Expected to find element")
	}
}
