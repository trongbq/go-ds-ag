package sorting

const succeed = "\u2713"
const failed = "\u2717"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualFloatArray(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] { // Careful with float check
			return false
		}
	}
	return true
}

func EqualRuneArray(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] { // Careful with float check
			return false
		}
	}
	return true
}
