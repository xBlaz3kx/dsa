package trees

import (
	"fmt"
	"log"

	"github.com/xBlaz3kx/dsa/structures/queue"
)

type BinaryTree[T comparable] struct {
	root *Node[T]
}

func NewBinaryTree[T comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

// Insert adds a new node to the binary tree.
func (t *BinaryTree[T]) Insert(value T) {
	if t.root != nil {
		t.root.AddNode(value)
		return
	}

	t.root = &Node[T]{Value: value}
}

// FindBFS traverses the binary tree using BFS.
func (t *BinaryTree[T]) FindBFS(value T) bool {
	// Use the queue to keep track of the neighbor nodes that need to be visited.
	q := queue.NewQueue[*Node[T]]()

	q.Push(t.root.LeftChild)
	q.Push(t.root.RightChild)

	// While the queue has elements
	for !q.IsEmpty() {

		pop, err := q.Pop()
		if err != nil {
			continue
		}

		tempNode := pop
		if tempNode != nil {
			fmt.Println((*tempNode).Value)
			if (*tempNode).Value == value {
				return true
			}
			if (*tempNode).RightChild != nil {
				q.Push((*tempNode).RightChild)
			}
			if (*tempNode).LeftChild != nil {
				q.Push((*tempNode).LeftChild)
			}
		}
	}

	return false
}

// BFS traverses the binary tree using BFS.
func (t *BinaryTree[T]) BFS() {
	q := queue.NewQueue[*Node[T]]()

	if t.root != nil {
		q.Push(t.root.LeftChild)
		q.Push(t.root.RightChild)
	}

	for !q.IsEmpty() {
		// Pop an element from the queue (looping through all the adjacent nodes)
		pop, err := q.Pop()
		if err != nil {
			break
		}

		log.Println(pop)

		tempNode := pop
		if tempNode != nil {
			log.Println((*tempNode).Value)

			if (*tempNode).RightChild != nil {
				q.Push((*tempNode).RightChild)
			}
			if (*tempNode).LeftChild != nil {
				q.Push((*tempNode).LeftChild)
			}
		}
	}
}

func (t *BinaryTree[T]) FindDFS(value int) bool {
	// return t.root.HasValueDFS(value)
	return false
}
