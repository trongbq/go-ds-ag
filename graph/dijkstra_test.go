package graph

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	wantedParent := map[int]int{
		1: Null,
		2: 5,
		3: 1,
		4: 5,
		5: 6,
		6: 3,
	}
	wantedDist := []int{0, 13, 2, 8, 7, 4}
	g := prepareWeightedGraph()
	dist, parent := Dijkstra(g, 1)
	if !reflect.DeepEqual(wantedParent, parent) {
		t.Errorf("Incorrect Dijkstra algorithm, Got parent: %v, Want: %v", parent, wantedParent)
	}
	if !equal(dist, wantedDist) {
		t.Errorf("Incorrect Dijkstra algorithm, Got dist: %v, Want: %v", dist, wantedDist)
	}
}
