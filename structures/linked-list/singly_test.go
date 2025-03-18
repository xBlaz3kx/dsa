package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	// Add some elements; 2, 3, 4, 5, 19, 12, 11
	list.AddElement(2)
	list.AddElement(3)
	list.AddElement(4)
	list.AddElement(5)
	list.AddElement(19)
	list.AddElement(12)
	list.AddElement(11)

	elements := list.GetElements()
	assert.Equal(t, elements, []int{2, 3, 4, 5, 19, 12, 11})

	element, err := list.FindElement(11)
	assert.NoError(t, err)
	assert.Equal(t, element.Value, 11)

	_, err = list.FindElement(10)
	assert.Error(t, err)
}
