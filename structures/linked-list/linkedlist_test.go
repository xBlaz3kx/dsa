package linked_list

import (
	"fmt"
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

	list.GetElements()

	element, err := list.FindElement(11)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(element.Value)

	_, err = list.FindElement(10)
	if err != nil {
		fmt.Println(err)
	}
}
