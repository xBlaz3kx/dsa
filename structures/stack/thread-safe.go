package stack

import (
	"errors"
	"sync"
)

// ThreadSafeStack is a thread-safe stack implementation
type ThreadSafeStack[T any] struct {
	elements []T
	mutex    sync.Mutex
}

// NewThreadSafeStack creates a new thread-safe stack
func NewThreadSafeStack[T any]() *ThreadSafeStack[T] {
	return &ThreadSafeStack[T]{elements: []T{}, mutex: sync.Mutex{}}
}

// Push adds an element to the stack
func (s *ThreadSafeStack[T]) Push(element T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.elements = append(s.elements, element)
}

// Pop removes an element from the stack
func (s *ThreadSafeStack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	index := len(s.elements) - 1
	retElement := s.elements[index]
	s.elements = s.elements[:index]
	return &retElement, nil
}

// Size returns the size of the stack
func (s *ThreadSafeStack[T]) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.elements)
}

// Peek returns the top element of the stack without removing it
func (s *ThreadSafeStack[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	index := len(s.elements) - 1
	return &s.elements[index], nil
}

// IsEmpty checks if the stack is empty
func (s *ThreadSafeStack[T]) IsEmpty() bool {
	return s.Size() == 0
}
