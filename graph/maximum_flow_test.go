package graph

import "testing"

func TestFordFulkerson(t *testing.T) {
	g := [][]int{
		{0, 3, 2, 0},
		{0, 0, 5, 2},
		{0, 0, 0, 3},
		{0, 0, 0, 0},
	}
	want := 5
	got := FordFulkerson(g, 0, 3)
	if want != got {
		t.Errorf("Incorrect ForkFulkerson algorithm: got: %v, want:%v", got, want)
	}
}
