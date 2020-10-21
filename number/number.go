package number

import (
	"math"
)

// Reverse return a number in reverse order
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

// Palindrome checks if a number is a palindrome.
// Algorithm tries to compare most significant digit to least significant digit
// and truncate both two digit each round.
// Time complexity is O(n) and space complexity is O(1).
func Palindrome(a int) bool {
	if a < 0 {
		return false
	}

	numDigits := int(math.Floor(math.Log10(float64(a)))) + 1
	mask := int(math.Pow10(numDigits - 1))
	for a != 0 {
		if (a / mask) != (a % 10) {
			return false
		}
		a %= mask // Eliminate most significant digit
		a /= 10   // Eliminate least significant digit
		mask /= 100
	}
	return true
}
