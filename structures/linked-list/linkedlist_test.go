package linked_list

import (
	"reflect"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.AddElement(2)
	list.AddElement(3)
	list.AddElement(4)
	list.AddElement(5)
	list.AddElement(19)
	list.AddElement(12)
	list.AddElement(11)

	elements := list.GetElements()

	// Why no pass?
	if reflect.DeepEqual(elements, []int{2, 3, 4, 5, 19, 12, 11}) {
		t.Errorf("Expected [2, 3, 4, 5, 19, 12, 11], got %v", elements)
	}

	element, err := list.FindElement(11)
	if err != nil {
		t.Error("Expected to find the element, got error: ", err)
	}

	if element.Value != 11 {
		t.Errorf("Expected 11, got %v", element.Value)
	}

	_, err = list.FindElement(10)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
