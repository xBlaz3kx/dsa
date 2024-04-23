package stack

import (
	"errors"
	"sync"
)

type ThreadSafeStack[T any] struct {
	elements []T
	mutex    sync.Mutex
}

func NewThreadSafeStack[T any]() *ThreadSafeStack[T] {
	return &ThreadSafeStack[T]{elements: []T{}, mutex: sync.Mutex{}}
}

func (s *ThreadSafeStack[T]) Push(element T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.elements = append(s.elements, element)
}

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

func (s *ThreadSafeStack[T]) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.elements)
}

func (s *ThreadSafeStack[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	index := len(s.elements) - 1
	return &s.elements[index], nil
}

func (s *ThreadSafeStack[T]) IsEmpty() bool {
	return s.Size() == 0
}
