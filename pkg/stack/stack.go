package stack

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: []T{}}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.elements) - 1
	retElement := s.elements[index]
	s.elements = s.elements[:index]
	return retElement, nil
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.elements) - 1
	return s.elements[index], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

type ThreadSafeStack[T any] struct {
	elements []T
	mutex    sync.Mutex
}

func NewThreadSafeStack[T any]() *ThreadSafeStack[T] {
	return &ThreadSafeStack[T]{elements: []T{}, mutex: sync.Mutex{}}
}

func (s *ThreadSafeStack[T]) Push(element interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.elements = append(s.elements, element)
}

func (s *ThreadSafeStack[T]) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	index := len(s.elements) - 1
	retElement := s.elements[index]
	s.elements = s.elements[:index]
	return retElement, nil
}

func (s *ThreadSafeStack[T]) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.elements)
}

func (s *ThreadSafeStack[T]) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	index := len(s.elements) - 1
	return s.elements[index], nil
}

func (s *ThreadSafeStack[T]) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.elements) == 0
}
