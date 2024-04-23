package queue

import (
	"errors"
	"sync"
)

// ThreadSafeQueue creates a queue structure that is thread-safe.
type ThreadSafeQueue[T any] struct {
	elements []T
	mutex    sync.Mutex
}

func NewThreadSafeQueue[T any]() *ThreadSafeQueue[T] {
	return &ThreadSafeQueue[T]{elements: []T{}, mutex: sync.Mutex{}}
}

// Push pushes the element in front of the queue.
func (q *ThreadSafeQueue[T]) Push(element T) {
	// Push in front of the queue
	var newElementsArray []T

	q.mutex.Lock()
	defer q.mutex.Unlock()

	newElementsArray = append(newElementsArray, element)
	newElementsArray = append(newElementsArray, q.elements...)

	q.elements = newElementsArray
}

// Pop removes an element from the end of the queue. If there are no elements, an error is returned.
func (q *ThreadSafeQueue[T]) Pop() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	defer q.mutex.Unlock()
	q.mutex.Lock()
	// Get last element
	lastElementIndex := len(q.elements) - 1
	element := q.elements[lastElementIndex]

	// Remove the element
	q.elements = q.elements[:lastElementIndex]

	return &element, nil
}

// Size returns the size of the queue
func (q *ThreadSafeQueue[T]) Size() int {
	return len(q.elements)
}

// IsEmpty checks if the queue is empty
func (q *ThreadSafeQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}
