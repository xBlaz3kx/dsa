package linked_list

import (
	"errors"
	"fmt"
)

// Doubly linked list
type DoublyLinkedList[T comparable] struct {
	head *Element[T]
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) AddElement(val T) {
	node := &Element[T]{
		Value: val,
	}

	if l.head == nil {
		l.head = node
		return
	}

	// Loop through all other elements and check if they have a pointer to the next element
	nextElement := l.head
	for nextElement != nil {

		// If the next element has no next element, it is the end of the list
		if nextElement.next == nil {
			// Add this element to the previous node
			nextElement.next = node
			// Add previous element to the current node
			node.prev = nextElement
			return
		}

		nextElement = nextElement.next
	}
}

func (l *DoublyLinkedList[T]) InsertElement(previousNodeValue, val T) {
	node := &Element[T]{
		Value: val,
	}

	if l.head == nil {
		l.head = node
		return
	}

	// Loop through all other elements and check if they have a pointer to the next element
	nextElement := l.head
	for nextElement != nil {

		// If the next element has no next element, it is the end of the list
		if nextElement.next == nil {
			// Add this element to the previous node
			nextElement.next = node
			// Add previous element to the current node
			node.prev = nextElement
			return
		}

		nextElement = nextElement.next
	}
}

func (l *DoublyLinkedList[T]) GetElements() {
	next := l.head

	for next != nil {
		fmt.Println(next.Value)
		next = next.next
	}
}

func (l *DoublyLinkedList[T]) RemoveElement(val T) *DoublyLinkedList[T] {
	if l.head.Value == val {

	}

	if l.head.next != nil {

	}

	return l
}

func (l *DoublyLinkedList[T]) FindElement(val T) (*Element[T], error) {
	element := l.head

	for element != nil {
		if element.Value == val {
			return element, nil
		}

		element = element.next
	}

	return nil, errors.New("element not found")
}
