package linked_list

import (
	"errors"
	"fmt"
)

// Singly linked list
type SinglyLinkedList[T comparable] struct {
	head *Element[T]
}

func NewLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

type Element[T comparable] struct {
	Value T
	next  *Element[T]
}

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

func (l *SinglyLinkedList[T]) GetElements() {
	next := l.head

	for next != nil {
		fmt.Println(next.Value)
		next = next.next
	}
}

func (l *SinglyLinkedList[T]) RemoveElement(val T) *SinglyLinkedList[T] {
	if l.head.Value == val {

	}

	if l.head.next != nil {

	}

	return l
}

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
