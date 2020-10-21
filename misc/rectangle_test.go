package misc

import "testing"

func TestIntersect(t *testing.T) {
	data := []struct {
		rectA Rect
		rectB Rect
		want  Rect
	}{
		{
			rectA: Rect{1, 2, 3, 4},
			rectB: Rect{5, 3, 2, 4},
			want:  Rect{0, 0, 0, 0},
		},
		{
			rectA: Rect{1, 2, 3, 4},
			rectB: Rect{2, 3, 2, 4},
			want:  Rect{2, 3, 2, 3},
		},
		{
			rectA: Rect{1, 1, 4, 5},
			rectB: Rect{2, 5, 2, 4},
			want:  Rect{2, 5, 2, 1},
		},
	}

	for _, c := range data {
		if got := Intersect(c.rectA, c.rectB); got != c.want {
			t.Errorf("Intersect test failed: input: %+v and %+v, want: %+v, got %+v", c.rectA, c.rectB, c.want, got)
		}
	}
}
