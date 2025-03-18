package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[string]()
	// Add some elements to the queue: 123, 124443, 1234, 4444, 44313
	q.Push("123")
	q.Push("124443")
	q.Push("1234")
	q.Push("4444")
	q.Push("44313")

	val, err := q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "123", *val)

	val, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "124443", *val)

	val, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "1234", *val)

	val, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "4444", *val)

	val, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "44313", *val)

	val, err = q.Pop()
	assert.ErrorContains(t, err, "queue is empty")
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue[string]()
	// Add some elements to the queue: 123, 124443, 1234, 4444, 44313
	q.Push("123")
	q.Push("124443")
	q.Push("1234")
	q.Push("4444")
	q.Push("44313")

	val, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "123", *val)

	val, err = q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "123", *val)

	val, err = q.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "123", *val)

	val, err = q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "124443", *val)
}
