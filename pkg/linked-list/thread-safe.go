package linked_list

import (
	"errors"
	"fmt"
	"sync"
)

type ThreadSafeLinkedList[T comparable] struct {
	head  *Element[T]
	mutex sync.Mutex
}

func NewThreadSafeLinkedList[T comparable]() *ThreadSafeLinkedList[T] {
	return &ThreadSafeLinkedList[T]{
		head:  nil,
		mutex: sync.Mutex{},
	}
}

func (l *ThreadSafeLinkedList[T]) AddElement(val T) {
	node := &Element[T]{
		Value: val,
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

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

func (l *ThreadSafeLinkedList[T]) GetElements() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	next := l.head

	for next != nil {
		fmt.Println(next.Value)
		next = next.next
	}
}

func (l *ThreadSafeLinkedList[T]) RemoveElement(val T) *ThreadSafeLinkedList[T] {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.head.Value == val {

	}

	if l.head.next != nil {

	}

	return l
}

func (l *ThreadSafeLinkedList[T]) FindElement(val T) (*Element[T], error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	element := l.head

	for element != nil {
		if element.Value == val {
			return element, nil
		}

		element = element.next
	}

	return nil, errors.New("element not found")
}
