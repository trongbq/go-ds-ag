package graph

import (
	"testing"
)

func TestBuildGraph(t *testing.T) {
	g := NewGraph(5, false)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(5, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 3)
	g.AddEdge(4, 3)

	if g.M != 7 {
		t.Error("Invalid number of edges")
	}
}

func TestCheckEdge(t *testing.T) {
	g := NewGraph(5, false)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(5, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 3)
	g.AddEdge(4, 3)

	if g.CheckEdge(5, 4) == false {
		t.Error("Incorrect check edge")
	}

	if g.CheckEdge(1, 4) == true {
		t.Error("Incorrect check edge")
	}
}

func prepare() Graph {
	g := NewGraph(6, false)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 6)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5)
	g.AddEdge(5, 4)
	g.AddEdge(4, 3)
	return g
}

func prepareDAG() Graph {
	g := NewGraph(7, true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(5, 4)
	g.AddEdge(6, 5)
	g.AddEdge(7, 1)
	g.AddEdge(7, 6)
	return g
}

func prepareDirectedGraphWithCycle() Graph {
	g := NewGraph(7, true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(5, 4)
	g.AddEdge(5, 2)
	g.AddEdge(6, 5)
	g.AddEdge(7, 1)
	g.AddEdge(7, 6)
	return g
}

func prepareWeightedDirectedGraph() Graph {
	g := NewGraph(6, true)
	g.addWeightedEdge(1, 2, 17)
	g.addWeightedEdge(1, 3, 2)
	g.addWeightedEdge(2, 3, 1)
	g.addWeightedEdge(2, 4, 11)
	g.addWeightedEdge(3, 5, 21)
	g.addWeightedEdge(3, 6, 2)
	g.addWeightedEdge(5, 4, 1)
	g.addWeightedEdge(5, 2, 6)
	g.addWeightedEdge(6, 5, 3)
	return g
}

func prepareWeightedGraph() Graph {
	g := NewGraph(6, false)
	g.addWeightedEdge(1, 2, 17)
	g.addWeightedEdge(1, 3, 2)
	g.addWeightedEdge(2, 3, 1)
	g.addWeightedEdge(2, 4, 11)
	g.addWeightedEdge(3, 5, 21)
	g.addWeightedEdge(3, 6, 2)
	g.addWeightedEdge(5, 4, 1)
	g.addWeightedEdge(5, 2, 6)
	g.addWeightedEdge(6, 5, 3)
	return g
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !equal(v, b[i]) {
			return false
		}
	}
	return true
}
