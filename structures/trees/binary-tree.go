package trees

import (
	"fmt"
	"log"

	"github.com/xBlaz3kx/dsa/structures/queue"
)

type BinaryTree[T any] struct {
	root *Node[T]
}

func NewBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

func (t *BinaryTree[T]) Print() {
	if t.root != nil {
		t.root.PrintNodes()
	}
}

func (t *BinaryTree[T]) Insert(value T) {
	if t.root != nil {
		t.root.AddNode(value)
		return
	}

	t.root = &Node[T]{Value: value}
}

func (t *BinaryTree[T]) FindBFS(value any) bool {
	q := queue.NewQueue[*Node]()

	q.Push(t.root.LeftChild)
	q.Push(t.root.RightChild)

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

func (t *BinaryTree[T]) BFS() {
	q := queue.NewQueue[*Node]()

	if t.root != nil {
		q.Push(t.root.LeftChild)
		q.Push(t.root.RightChild)
	}

	for !q.IsEmpty() {
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
	return t.root.HasValueDFS(value)
}

type Node[T any] struct {
	Value      T
	LeftChild  *Node[T]
	RightChild *Node[T]
}

func (n *Node[T]) AddNode(value T) {
	// Rule: Left child value should be always lesser than right child value
	if n.Value < value {
		if n.RightChild == nil {
			// Put as right child
			n.RightChild = &Node{Value: value}
			return
		}

		// Add to right node as a child
		n.RightChild.AddNode(value)
	} else {
		if n.LeftChild == nil {
			n.LeftChild = &Node{Value: value}
			return
		}

		// Put as left child's node
		n.LeftChild.AddNode(value)
	}
}

func (n *Node[T]) HasValueDFS(value T) bool {

	// Check if this nodes' value matches the searched value
	if n.Value != value {

		if n.Value > value {
			// Check left child recursively
			if n.LeftChild != nil {
				return n.LeftChild.HasValueDFS(value)
			}

			return false
		} else {
			// Check right child recursively
			if n.RightChild != nil {
				return n.RightChild.HasValueDFS(value)
			}

			return false
		}
	}

	return true
}

func (n *Node[T]) PrintNodes() {
	if n.RightChild != nil {
		n.RightChild.PrintNodes()
	}

	if n.LeftChild != nil {
		n.LeftChild.PrintNodes()
	}

	fmt.Println(n.Value)
}
