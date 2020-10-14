package bit

import "testing"

func TestSwapBit(t *testing.T) {
	want := int64(11)
	got := SwapBit(73, 1, 6)
	if want != got {
		t.Errorf("SwapBit test error: want %v, got %v", want, got)
	}
}
