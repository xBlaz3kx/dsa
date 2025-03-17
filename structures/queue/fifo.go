package queue

import (
	"errors"
)

// Queue represents a fifo queue
type Queue[T any] struct {
	elements []T
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{elements: []T{}}
}

// Push pushes the element in front of the queue.
func (q *Queue[T]) Push(element T) {
	// Push in front of the queue
	var newElementsArray []T

	newElementsArray = append(newElementsArray, element)
	newElementsArray = append(newElementsArray, q.elements...)

	q.elements = newElementsArray
}

// Pop removes an element from the end of the queue. If there are no elements, an error is returned.
func (q *Queue[T]) Pop() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	// Get last element
	lastElementIndex := len(q.elements) - 1
	element := q.elements[lastElementIndex]

	// Remove the element
	q.elements = q.elements[:lastElementIndex]
	return &element, nil
}

// Peek checks the latest element from the queue.
func (q *Queue[T]) Peek() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	// Get last element
	lastElementIndex := len(q.elements) - 1
	element := q.elements[lastElementIndex]

	return &element, nil
}

// Size returns the size of the queue
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
