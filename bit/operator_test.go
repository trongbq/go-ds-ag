package bit

import "testing"

func TestAdd(t *testing.T) {
	want := int64(34)
	got := Add(25, 9)
	if got != want {
		t.Errorf("Add test error: want: %v, got: %v", want, got)
	}
}

func TestMultiply(t *testing.T) {
	want := int64(225)
	got := Multiply(25, 9)
	if got != want {
		t.Errorf("Multiply test error: want: %v, got: %v", want, got)
	}
}
