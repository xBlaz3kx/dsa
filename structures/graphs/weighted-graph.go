package graph

import linkedlist "github.com/xBlaz3kx/dsa/structures/linked-list"

type WeightedEdge struct {
	Weight int
	Node   int
}

// Graph implementation with adjacency list technique.
// Adjacency list consist of the map and list of elemets, linked to the node.
type WeightedGraph struct {
	Nodes []*linkedlist.SinglyLinkedList[WeightedEdge]
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		Nodes: []*linkedlist.SinglyLinkedList[WeightedEdge]{},
	}
}
func (g *WeightedGraph) AddEdge(node int, child WeightedEdge) {
	if g.Nodes[node] == nil {
		g.Nodes[node] = linkedlist.NewSinglyLinkedList[WeightedEdge]()
	}

	g.Nodes[node].AddElement(child)
}

func (g *WeightedGraph) RemoveEdge(node int, edge int) {
	if g.Nodes[node] == nil {
		// do nothing
		return
	}

	// g.Nodes[node].RemoveElement(child)
}

func (g *WeightedGraph) Djikstra(startingNode, targetNode int) {

}

func (g *WeightedGraph) BellmanFord(startingNode, targetNode int) {

}
