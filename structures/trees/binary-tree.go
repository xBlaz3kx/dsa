package trees

import (
	"fmt"
	"log"

	"github.com/xBlaz3kx/dsa/structures/queue"
	stack2 "github.com/xBlaz3kx/dsa/structures/stack"
)

type BinaryTree[T Comparable] struct {
	root *Node[T]
}

func NewBinaryTree[T Comparable]() *BinaryTree[T] {
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

// FindBFS traverses the binary tree using TraverseBFS.
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
			if (*tempNode).Value.Compare(value) == 0 {
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

// TraverseBFS traverses the binary tree using BFS.
func (t *BinaryTree[T]) TraverseBFS() {
	q := queue.NewQueue[*Node[T]]()
	// traversalOrder := queue.NewQueue[*Node[T]]()

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

func (t *BinaryTree[T]) FindDFS(value T) bool {
	stack := stack2.NewStack[*Node[T]]()

	if t.root != nil {
		comparison := t.root.Value.Compare(value)
		if comparison == 0 {
			return true
		}

		if t.root.RightChild != nil {
			stack.Push(t.root.RightChild)
		}
		if t.root.LeftChild != nil {
			stack.Push(t.root.LeftChild)
		}
	}

	for !stack.IsEmpty() {
		pop, err := stack.Pop()
		if err != nil {
			continue
		}

		tempNode := pop
		if tempNode != nil {
			log.Println((*tempNode).Value)

			comparison := (*tempNode).Value.Compare(value)
			if comparison == 0 {
				return true
			}

			if (*tempNode).RightChild != nil {
				stack.Push((*tempNode).RightChild)
			}

			if (*tempNode).LeftChild != nil {
				stack.Push((*tempNode).LeftChild)
			}
		}
	}

	return false
}
