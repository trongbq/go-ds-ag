package graph

import (
	"testing"
)

func TestTraverseUndirectedGraph(t *testing.T) {
	data := []struct {
		parent int
		child  int
	}{
		{1, 2},
		{1, 5},
		{1, 6},
		{2, 3},
		{5, 4},
	}
	g := prepare()
	parent := Bfs(g, 1)
	for _, v := range data {
		if parent[v.child] != v.parent {
			t.Errorf("Parent of %v is not %v but it is %v\n", v.child, v.parent, parent[v.child])
		}
	}
}

func TestFindPathConnected(t *testing.T) {
	g := prepare()
	parent := Bfs(g, 1)
	want := []int{1, 5, 4}
	got := FindPath(1, 4, parent)
	if !equal(want, got) {
		t.Errorf("Can not find correct path from %v to %v, want: %v, got: %v", 1, 4, want, got)
	}
}

func TestFindPathNotConnected(t *testing.T) {
	g := prepare()
	parent := Bfs(g, 1)
	want := []int{}
	got := FindPath(2, 4, parent)
	if !equal(want, got) {
		t.Errorf("%v can not link to %v but still got result, want: %v, got: %v", 1, 4, want, got)
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
