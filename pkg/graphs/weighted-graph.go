package graph

import linkedlist "github.com/xBlaz3kx/dsa/pkg/linked-list"

// Graph implementation with adjacency list technique.
// Adjacency list consist of the map and list of elemets, linked to the node.
type WeightedGraph struct {
	Nodes []*linkedlist.SinglyLinkedList[int]
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		Nodes: []*linkedlist.SinglyLinkedList[int]{},
	}
}
func (g *WeightedGraph) AddEdge(node, child int) {
	if g.Nodes[node] == nil {
		g.Nodes[node] = linkedlist.NewLinkedList[int]()
	}

	g.Nodes[node].AddElement(child)
}

func (g *WeightedGraph) RemoveEdge(node, child int) {
	if g.Nodes[node] == nil {
		// do nothing
		return
	}

	g.Nodes[node].RemoveElement(child)
}

func (g *WeightedGraph) Djikstra(startingNode, targetNode int) {

}

func (g *WeightedGraph) BellmanFord() {

}

func (g *WeightedGraph) TopologicalSort() []int {
	return make([]int, 0)
}
