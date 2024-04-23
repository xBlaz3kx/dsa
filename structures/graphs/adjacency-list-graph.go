package graph

import linkedlist "github.com/xBlaz3kx/dsa/structures/linked-list"

// Graph implementation with adjacency list technique.
// Adjacency list consist of the map and list of elemets, linked to the node.
type AdjacencyListGraph struct {
	Nodes []*linkedlist.SinglyLinkedList[int]
}

func NewAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		Nodes: []*linkedlist.SinglyLinkedList[int]{},
	}
}
func (g *AdjacencyListGraph) AddEdge(node, child int) {
	if g.Nodes[node] == nil {
		g.Nodes[node] = linkedlist.NewSinglyLinkedList[int]()
	}

	g.Nodes[node].AddElement(child)
}

func (g *AdjacencyListGraph) RemoveEdge(node, child int) {
	if g.Nodes[node] == nil {
		// do nothing
		return
	}

	g.Nodes[node].RemoveElement(child)
}

func (g *AdjacencyListGraph) BFS(node int) {

}

func (g *AdjacencyListGraph) DFS(node int) {

}

func (g *AdjacencyListGraph) TopologicalSort() []int {
	return make([]int, 0)
}
