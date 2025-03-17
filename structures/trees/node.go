package trees

import "fmt"

type Comparable interface {
	// Compare returns the equality;
	// 0 if both node values are equal,
	// 1 if current node value is greater than the one provided
	// -1 if current node value is smaller than the one provided.
	Compare(other Comparable) int
}

// Node represents a single node in a binary tree. It contains a value and two children nodes.
type Node[T Comparable] struct {
	Value      T
	LeftChild  *Node[T]
	RightChild *Node[T]
}

// AddNode adds a new node to the binary tree.
// It will add the node as a left child if the value is lesser than the current node's value.
// Otherwise, it will add it as a right child.
func (n *Node[T]) AddNode(value T) {

	// Rule: Left child value should be always lesser than right child value
	if n.Value.Compare(value) < 0 {

		// If the node has no right child, add the new node as the right child
		if n.RightChild == nil {
			n.RightChild = &Node[T]{Value: value}
			return
		}

		// Add the node as the right child's node
		n.RightChild.AddNode(value)
	} else {

		// If the node has no left child, add the new node as the left child
		if n.LeftChild == nil {
			n.LeftChild = &Node[T]{Value: value}
			return
		}

		// Add the node as the left child's node
		n.LeftChild.AddNode(value)
	}
}

// HasValueDFS checks if the binary tree has a node with the given value using Depth first search
func (n *Node[T]) HasValueDFS(value T) bool {
	// Check if this nodes' value matches the searched value
	if n.Value.Compare(value) == 0 {
		return true
	}

	if n.Value.Compare(value) < 1 {
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

// HasValueBFS checks if the binary tree has a node with the given value using Breadth first search
func (n *Node[T]) HasValueBFS(value T) bool {
	// Check if this nodes' value matches the searched value
	if n.Value.Compare(value) == 0 {
		return true
	}

	if n.Value.Compare(value) < 1 {
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

func (n *Node[T]) PrintNodes() {
	if n.RightChild != nil {
		n.RightChild.PrintNodes()
	}

	if n.LeftChild != nil {
		n.LeftChild.PrintNodes()
	}

	fmt.Println(n.Value)
}
