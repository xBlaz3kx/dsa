package stack

import (
	"errors"
)

// Stack is a generic stack implementation
type Stack[T any] struct {
	elements []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: []T{}}
}

// Push adds an element to the stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes the last element from the stack
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.elements) - 1
	retElement := s.elements[index]
	s.elements = s.elements[:index]
	return &retElement, nil
}

// Size returns the size of the stack
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

// Peek returns the last element from the stack
func (s *Stack[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.elements) - 1
	return &s.elements[index], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
