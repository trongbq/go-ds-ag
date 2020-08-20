package graph

import "testing"

func TestTopologicalSorting(t *testing.T) {
	data := []int{7, 1, 2, 3, 6, 5, 4}
	g := prepareDAG()
	arr, err := TopologicalSorting(g)
	if err != nil {
		t.Errorf("Error with topological sorting: %v", err)
	}
	if !equal(data, arr) {
		t.Errorf("Incorrect algorithm find cycles, Got: %v, Want: %v", arr, data)
	}
}

func TestTopologicalSortingWithCycle(t *testing.T) {
	g := prepareDirectedGraphWithCycle()
	_, err := TopologicalSorting(g)
	if err == nil {
		t.Errorf("Fail to return error!")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}
