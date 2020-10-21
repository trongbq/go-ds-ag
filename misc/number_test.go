package number

import "testing"

func TestReverse(t *testing.T) {
	want := -4317
	got := Reverse(-7134)
	if got != want {
		t.Errorf("Reverse test error: want: %v, got: %v", want, got)
	}
}

func TestPalindrome(t *testing.T) {
	data := []struct {
		input int
		want  bool
	}{
		{
			input: 0,
			want:  true,
		},
		{
			input: 11011,
			want:  true,
		},
		{
			input: -1,
			want:  false,
		},
		{
			input: 3123,
			want:  false,
		},
		{
			input: 9987899,
			want:  true,
		},
	}

	for _, c := range data {
		if got := Palindrome(c.input); got != c.want {
			t.Errorf("Palindrom test failed: input: %v, want: %v, got: %v", c.input, c.want, got)
		}
	}
}
