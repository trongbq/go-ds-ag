package bit

import "testing"

func TestSwapBit(t *testing.T) {
	want := int64(11)
	got := SwapBit(73, 1, 6)
	if want != got {
		t.Errorf("SwapBit test error: want %v, got %v", want, got)
	}
}

func TestCountBit(t *testing.T) {
	want := 3
	got := CountBit(13)
	if want != got {
		t.Errorf("CountBit test error: want %v, got %v", want, got)
	}
}

func TestParity(t *testing.T) {
	data := []struct {
		input int64
		want  int
	}{
		{
			input: 11,
			want:  1,
		},
		{
			input: 10,
			want:  0,
		},
	}

	for _, c := range data {
		r := Parity(c.input)
		if r != c.want {
			t.Errorf("Parity test error: want %v, got %v", c.want, r)
		}
	}
}

func TestParityOpt(t *testing.T) {
	data := []struct {
		input int64
		want  int
	}{
		{
			input: 11,
			want:  1,
		},
		{
			input: 10,
			want:  0,
		},
	}

	for _, c := range data {
		r := ParityOpt(c.input)
		if r != c.want {
			t.Errorf("ParityOpt test error: want %v, got %v", c.want, r)
		}
	}
}
