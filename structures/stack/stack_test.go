package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	val, err := stack.Pop()
	if *val != 4 || err != nil {
		t.Error("Expected 4")
	}

	val, err = stack.Pop()
	if *val != 3 || err != nil {
		t.Error("Expected 3")
	}

	val, err = stack.Pop()
	if *val != 2 || err != nil {
		t.Error("Expected 2")
	}

	val, err = stack.Pop()
	if *val != 1 || err != nil {
		t.Error("Expected 1")
	}

	val, err = stack.Pop()
	if err == nil {
		t.Error("Expected error")
	}
}
