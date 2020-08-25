package comsearch

import (
	"reflect"
	"testing"
)

func TestSubset(t *testing.T) {
	want := [][]int{
		{},
		{1},
		{2},
		{3},
		{1, 2},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}
	got := Subset(3)

	if len(got) != len(want) {
		t.Errorf("Incorrect subset algorithms: got: %v, want: %v", got, want)
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
