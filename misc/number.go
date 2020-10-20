package number

import "math"

func Reverse(a int) int {
	var r int
	t := int(math.Abs(float64(a)))
	for t != 0 {
		r = r*10 + t%10
		t /= 10
	}

	if a < 0 {
		return -r
	}
	return r
}
