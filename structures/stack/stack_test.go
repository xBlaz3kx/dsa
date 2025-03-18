package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	assert.Equal(t, 4, stack.Size())

	val, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 4, *val)

	assert.Equal(t, 3, stack.Size())

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)

	assert.Equal(t, 2, stack.Size())

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)

	assert.Equal(t, 1, stack.Size())

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)

	val, err = stack.Pop()
	assert.ErrorContains(t, err, "stack is empty")

	assert.Equal(t, 0, stack.Size())
	assert.True(t, stack.IsEmpty())
}

func TestStackPeek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	val, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 4, *val)

	val, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 4, *val)

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 4, *val)

	val, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)
}
