package linked_list

import (
	"errors"
)

type List[T comparable] interface {
	AddElement(val T)
	GetElements() []T
	RemoveElement(val T) error
	FindElement(val T) (*T, error)
}

// Singly linked list
type SinglyLinkedList[T comparable] struct {
	head *Element[T]
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

type Element[T comparable] struct {
	Value T
	next  *Element[T]
	prev  *Element[T]
}

// AddElement adds an element to the end of a linked list
func (l *SinglyLinkedList[T]) AddElement(val T) {
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
		if nextElement.next == nil {
			nextElement.next = node
			return
		}

		nextElement = nextElement.next
	}
}

// GetElements returns values of elements in a linked list.
func (l *SinglyLinkedList[T]) GetElements() []T {
	elements := []T{}

	element := l.head
	for element != nil {
		// Add element to the list
		elements = append(elements, element.Value)

		// Move to the next element
		element = element.next
	}

	return elements
}

// RemoveElement removes an element from a linked list
func (l *SinglyLinkedList[T]) RemoveElement(val T) error {
	// If head is nil, return the list
	if l.head == nil {
		return errors.New("list is empty")
	}

	element := l.head
	var previousElement *Element[T]

	for element != nil {
		if element.Value == val {
			if previousElement != nil {
				previousElement.next = element.next
			}
		}

		previousElement = element
		element = element.next
	}

	return nil
}

// FindElement finds an element in a linked list
func (l *SinglyLinkedList[T]) FindElement(val T) (*Element[T], error) {
	element := l.head
	for element != nil {
		if element.Value == val {
			return element, nil
		}

		element = element.next
	}

	return nil, errors.New("element not found")
}
