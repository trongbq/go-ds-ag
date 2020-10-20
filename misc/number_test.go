package number

import "testing"

func TestReverse(t *testing.T) {
	want := -4317
	got := Reverse(-7134)
	if got != want {
		t.Errorf("Reverse test error: want: %v, got: %v", want, got)
	}
}
