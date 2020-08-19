package graph

import (
	"testing"
)

func TestFindCycle(t *testing.T) {
	data := [][]int{{2, 5, 4, 3},{1, 2, 5}}
	g := prepare()
	cycles := FindCycleInGraph(g, 1)
	if !equal2D(data, cycles) {
		t.Errorf("Incorrect algorithm find cycles, Got: %v, Want: %v", cycles, data)
	}
}

