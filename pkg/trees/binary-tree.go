package trees

import (
	"fmt"
	"log"

	"github.com/xBlaz3kx/dsa/pkg/queue"
)

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Print() {
	if t.root != nil {
		t.root.PrintNodes()
	}
}

func (t *BinaryTree) Insert(value int) {
	if t.root != nil {
		t.root.AddNode(value)
		return
	}

	t.root = &Node{Value: value}
}

func (t *BinaryTree) FindBFS(value int) bool {
	q := queue.NewQueue[any]()

	q.Push(t.root.LeftChild)
	q.Push(t.root.RightChild)

	for !q.IsEmpty() {
		pop, err := q.Pop()
		if err != nil {
			continue
		}

		tempNode := pop.(*Node)
		if tempNode != nil {
			fmt.Println(tempNode.Value)
			if tempNode.Value == value {
				return true
			}
			if tempNode.RightChild != nil {
				q.Push(tempNode.RightChild)
			}
			if tempNode.LeftChild != nil {
				q.Push(tempNode.LeftChild)
			}
		}
	}

	return false
}

func (t *BinaryTree) BFS() {
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
			log.Println(tempNode.Value)

			if tempNode.RightChild != nil {
				q.Push(tempNode.RightChild)
			}
			if tempNode.LeftChild != nil {
				q.Push(tempNode.LeftChild)
			}
		}
	}
}

func (t *BinaryTree) FindDFS(value int) bool {
	return t.root.HasValueDFS(value)
}

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

func (n *Node) AddNode(value int) {
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

func (n *Node) HasValueDFS(value int) bool {

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

func (n *Node) PrintNodes() {
	if n.RightChild != nil {
		n.RightChild.PrintNodes()
	}

	if n.LeftChild != nil {
		n.LeftChild.PrintNodes()
	}

	fmt.Println(n.Value)
}
