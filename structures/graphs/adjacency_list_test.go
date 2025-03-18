package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjacencyListGraph_FindUsingBFS(t *testing.T) {
	graph := NewAdjacencyListGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	bfs, err := graph.FindUsingBFS(1, 3)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if bfs != true {
		t.Errorf("Expected true, got %v", bfs)
	}

	_, err = graph.FindUsingBFS(15, 1)
	if err == nil {
		t.Errorf("Expected error")
	}

	bfs, err = graph.FindUsingBFS(1, 15)
	if err != nil {
		t.Errorf("Expected error")
	}

	if bfs != false {
		t.Errorf("Expected false, got %v", bfs)
	}
}

func TestAdjacencyListGraph_FindUsingDFS(t *testing.T) {
	graph := NewAdjacencyListGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	bfs, err := graph.FindUsingDFS(1, 3)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if bfs != true {
		t.Errorf("Expected true, got %v", bfs)
	}

	_, err = graph.FindUsingDFS(15, 1)
	if err == nil {
		t.Errorf("Expected error")
	}

	bfs, err = graph.FindUsingDFS(1, 15)
	if err != nil {
		t.Errorf("Expected error")
	}

	if bfs != false {
		t.Errorf("Expected false, got %v", bfs)
	}
}

func TestAdjacencyListGraph_IsCyclic(t *testing.T) {
	cyclicGraph := NewAdjacencyListGraph()

	cyclicGraph.AddEdge(0, 1)
	cyclicGraph.AddEdge(0, 0)
	cyclicGraph.AddEdge(1, 2)
	cyclicGraph.AddEdge(2, 0)

	assert.Truef(t, cyclicGraph.IsCyclic(), "Graph is cyclic, but IsCyclic returned false")

	nonCyclicGraph := NewAdjacencyListGraph()

	nonCyclicGraph.AddEdge(0, 1)
	nonCyclicGraph.AddEdge(1, 2)
	nonCyclicGraph.AddEdge(2, 3)

	assert.False(t, cyclicGraph.IsCyclic(), "Graph is not cyclic, but IsCyclic returned true")
}

func TestAdjacencyListGraph_TopologicalSort(t *testing.T) {

}
