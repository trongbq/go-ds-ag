package comsearch

import (
	"reflect"
	"testing"
)

func TestSubset(t *testing.T) {
	want := [][]int{
		{},
		{0},
		{1},
		{2},
		{0, 1},
		{0, 2},
		{1, 2},
		{0, 1, 2},
	}
	got := Subset(2)

	if len(got) != len(want) {
		t.Fatalf("Incorrect subset algorithms: got: %v, want: %v", got, want)
	}

	for i := 0; i < len(got); i++ {
		f := false
		for j := 0; j < len(want); j++ {
			if reflect.DeepEqual(got[i], want[j]) {
				f = true
			}
		}
		if !f {
			t.Fatalf("Incorrect subset algorithms: got: %v, want: %v", got, want)
		}
	}
}

func TestPermutation(t *testing.T) {
	want := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}
	got := Permuation(2)

	if len(got) != len(want) {
		t.Fatalf("Incorrect permutation algorithms: got: %v, want: %v", got, want)
	}

	for i := 0; i < len(got); i++ {
		f := false
		for j := 0; j < len(want); j++ {
			if reflect.DeepEqual(got[i], want[j]) {
				f = true
			}
		}
		if !f {
			t.Fatalf("Incorrect permutation algorithms: got: %v, want: %v", got, want)
		}
	}
}

func TestCountAllPathInGraph(t *testing.T) {
	g := [][]int{
		{0, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 0},
	}
	want := 7
	got := CountAllPathInGraph(g, 0, 2)

	if got != want {
		t.Fatalf("Incorrect CountAllPathInGraph algorithms: got: %v, want: %v", got, want)
	}
}
