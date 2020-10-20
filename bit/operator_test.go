package bit

import "testing"

func TestAdd(t *testing.T) {
	want := uint64(34)
	got := Add(25, 9)
	if got != want {
		t.Errorf("Add test error: want: %v, got: %v", want, got)
	}
}

func TestMultiply(t *testing.T) {
	want := uint64(225)
	got := Multiply(25, 9)
	if got != want {
		t.Errorf("Multiply test error: want: %v, got: %v", want, got)
	}
}

func TestDivide(t *testing.T) {
	data := []struct {
		input struct {
			a uint64
			b uint64
		}
		want uint64
	}{
		{
			input: struct {
				a uint64
				b uint64
			}{
				a: 11,
				b: 2,
			},
			want: 5,
		},
		{
			input: struct {
				a uint64
				b uint64
			}{
				a: 87,
				b: 5,
			},
			want: 17,
		},
	}
	for _, c := range data {
		if got := Divide(c.input.a, c.input.b); got != c.want {
			t.Errorf("Divide test error: want: %v, got: %v", c.want, got)
		}
	}
}

func TestPower(t *testing.T) {
	data := []struct {
		input struct {
			a uint64
			b uint64
		}
		want uint64
	}{
		{
			input: struct {
				a uint64
				b uint64
			}{
				a: 4,
				b: 12,
			},
			want: 16777216,
		},
		{
			input: struct {
				a uint64
				b uint64
			}{
				a: 3,
				b: 5,
			},
			want: 243,
		},
	}
	for _, c := range data {
		if got := Power(c.input.a, c.input.b); got != c.want {
			t.Errorf("Power test error for %v^%v: want: %v, got: %v", c.input.a, c.input.b, c.want, got)
		}
	}
}
