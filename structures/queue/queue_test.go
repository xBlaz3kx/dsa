package queue

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[string]()
	q.Push("123")
	q.Push("124443")
	q.Push("1234")
	q.Push("4444")
	q.Push("44313")

	fmt.Println(q)

	for !q.IsEmpty() {
		val, err := q.Pop()
		if err != nil {
			continue
		}
		fmt.Println(*val)
	}
}
