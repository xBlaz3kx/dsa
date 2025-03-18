package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertionSort(t *testing.T) {
	array := []int{11, 2, 3, 6, 1, 22, 15, 50, 31}

	InsertionSort(array)

	assert.Equal(t, []int{1, 2, 3, 6, 11, 15, 22, 31, 50}, array)
}
