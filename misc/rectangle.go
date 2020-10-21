package misc

type Rect struct {
	X, Y, Width, Height int
}

// Intersect return the rectable formed by two input rectables
func Intersect(a, b Rect) Rect {
	if !IsIntersect(a, b) {
		return Rect{0, 0, 0, 0}
	}
	return Rect{
		X:      max(a.X, b.X),
		Y:      max(a.Y, b.Y),
		Width:  min(a.X+a.Width, b.X+b.Width) - max(a.X, b.X),
		Height: min(a.Y+a.Height, b.Y+b.Height) - max(a.Y, b.Y),
	}
}

func IsIntersect(a, b Rect) bool {
	return a.X <= b.X+b.Width && b.X <= a.X+a.Width &&
		a.Y <= b.Y+b.Height && b.Y <= a.Y+a.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
