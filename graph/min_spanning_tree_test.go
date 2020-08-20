package graph

import (
	"reflect"
	"testing"
)

func TestMinimumSpanningTreeWithPrim(t *testing.T) {
	data := map[int]int{
		1: Null,
		2: 5,
		3: 1,
		4: 5,
		5: 6,
		6: 3,
	}
	g := prepareWeightedGraph()
	parent := Prim(g, 1)
	if !reflect.DeepEqual(data, parent) {
		t.Errorf("Incorrect Prim algorithm, Got: %v, Want: %v", parent, data)
	}
}
