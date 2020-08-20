package graph

import (
	"testing"
)

func TestDfsTraverseUndirectedGraph(t *testing.T) {
	g := prepare()
	parent, records := Dfs(g, 1)

	// Verify parent data
	parentData := []struct {
		child  int
		parent int
	}{
		{6, 1},
		{2, 1},
		{5, 2},
		{4, 5},
		{3, 4},
	}
	for _, v := range parentData {
		if parent[v.child] != v.parent {
			t.Errorf("Parent of %v is not %v but it is %v\n", v.child, v.parent, parent[v.child])
		}
	}

	// Verify records data
	recordData := []struct {
		v   int
		in  int
		out int
	}{
		{1, 1, 12},
		{2, 4, 11},
		{3, 7, 8},
		{4, 6, 9},
		{5, 5, 1},
		{6, 2, 3},
	}

	for _, d := range recordData {
		r := records[d.v]
		if r.In != d.in && r.Out == d.out {
			t.Errorf("Record time in or/and out of %v is not %v but it is %v\n", d.v, r, records[d.v])
		}
	}
}

func TestFindCycle(t *testing.T) {
	data := [][]int{{2, 5, 4, 3}, {1, 2, 5}}
	g := prepare()
	cycles := FindCycleInGraph(g, 1)
	if !equal2D(data, cycles) {
		t.Errorf("Incorrect algorithm find cycles, Got: %v, Want: %v", cycles, data)
	}
}
