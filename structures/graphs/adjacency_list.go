package graph

import (
	"fmt"

	linkedlist "github.com/xBlaz3kx/dsa/structures/linked-list"
	"github.com/xBlaz3kx/dsa/structures/queue"
	stack2 "github.com/xBlaz3kx/dsa/structures/stack"
)

// Graph implementation with adjacency list technique.
// Adjacency list consist of the map and list of elemets, linked to the node.
type AdjacencyListGraph struct {
	// A list of nodes and their edges.
	nodes map[int]*linkedlist.SinglyLinkedList[int]
}

func NewAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		nodes: map[int]*linkedlist.SinglyLinkedList[int]{},
	}
}

// AddEdge adds a new edge to the graph.
func (g *AdjacencyListGraph) AddEdge(node, edge int) {
	// Creates a new linked list if the node does not exist.
	if g.nodes[node] == nil {
		g.nodes[node] = linkedlist.NewSinglyLinkedList[int]()
	}

	g.nodes[node].AddElement(edge)
}

// RemoveEdge removes an edge from the graph.
func (g *AdjacencyListGraph) RemoveEdge(node, edge int) error {
	if g.nodes[node] == nil {
		// do nothing
		return nil
	}

	return g.nodes[node].RemoveElement(edge)
}

// FindUsingBFS finds a node using Breadth First Search.
func (g *AdjacencyListGraph) FindUsingBFS(startNode, searchedNode int) (bool, error) {
	// Use a queue to keep track of the nodes needed to be visited.
	nodeQueue := queue.NewQueue[int]()
	visited := make(map[int]bool)

	// Verify that the start node exists
	if _, exists := g.nodes[startNode]; !exists {
		return false, fmt.Errorf("node %d does not exist", startNode)
	}

	visited[startNode] = true
	nodeQueue.Push(startNode)

	// While the queue is not empty
	for nodeQueue.Size() > 0 {

		// Pop the node from the queue
		currentNode, err := nodeQueue.Pop()
		if err != nil {
			return false, err
		}

		// List through neighbors of the current node
		edges := g.nodes[*currentNode]
		if edges == nil {
			continue
		}

		for _, node := range edges.GetElements() {
			// If the node has not been visited, add it to the queue
			if !visited[node] {
				visited[node] = true
				nodeQueue.Push(node)
			}

			// If the node is the searched node, return true
			if node == searchedNode {
				return true, nil
			}
		}
	}

	return false, nil
}

// FindUsingDFS finds a node using Depth First Search.
func (g *AdjacencyListGraph) FindUsingDFS(startNode, desiredNode int) (bool, error) {
	stack := stack2.NewStack[int]()
	visited := make(map[int]bool)

	// Verify that the start node exists
	if _, exists := g.nodes[startNode]; !exists {
		return false, fmt.Errorf("node %d does not exist", startNode)
	}

	visited[startNode] = true
	stack.Push(startNode)

	// While the stack is not empty
	for stack.Size() > 0 {

		// Pop the node from the stack
		currentNode, err := stack.Pop()
		if err != nil {
			return false, err
		}

		// Check if current node has any edges/neighbors
		neighbours := g.nodes[*currentNode]
		if neighbours == nil {
			continue
		}

		// List through node of the current node
		for _, node := range neighbours.GetElements() {
			if node == desiredNode {
				return true, nil
			}

			if !visited[node] {
				visited[node] = true
				stack.Push(node)
			}
		}
	}

	return false, nil
}

// IsCyclic checks if the graph is cyclic.
func (g *AdjacencyListGraph) IsCyclic() bool {
	stack := stack2.NewStack[int]()
	visited := make(map[int]bool)

	for key, _ := range g.nodes {
		stack.Push(key)
	}

	// While the stack is not empty
	for stack.Size() > 0 {

		// Pop the node from the stack
		currentNode, err := stack.Pop()
		if err != nil {
			return false
		}

		// Check if current node has any edges/neighbors
		neighbours := g.nodes[*currentNode]
		if neighbours == nil {
			continue
		}

		// List through edges of the current node
		for _, element := range neighbours.GetElements() {
			// If the element has been visited, the graph is cyclic
			if visited[element] {
				return true
			}

			visited[element] = true
			stack.Push(element)
		}
	}

	return false
}

// TopologicalSort sorts the graph topologically.
func (g *AdjacencyListGraph) TopologicalSort() []int {
	sortedList := []int{}

	stack := stack2.NewStack[int]()
	visited := make(map[int]bool)

	// While the stack is not empty
	for stack.Size() > 0 {

		// Pop the node from the stack
		val, err := stack.Pop()
		if err != nil {
			continue
		}

		// List through children of the current node
		for _, element := range g.nodes[*val].GetElements() {
			if !visited[element] {
				visited[element] = true
				stack.Push(element)
			}
		}
	}

	return sortedList
}
