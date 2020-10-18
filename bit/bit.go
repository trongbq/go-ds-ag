package bit

func SwapBit(a, i, j int64) int64 {
	// We only need to swap when value at i and j are different
	if (a>>i)&1 != (a>>j)&1 {
		// Since bit value is either 1 or 0,
		// so we only have to use XOR operator to flip bit at i and j
		mask := (int64(1) << i) | (int64(1) << j)
		return a ^ mask
	}
	return a
}

func CountBit(a int64) int {
	var c int
	for a != 0 {
		c += int(a & 1)
		a >>= 1
	}
	return c
}

// The parity of a binary word is 1 if the number of 1s in the word is odd,
// otherwise, it is 0
func Parity(a int64) int {
	var p int
	for a != 0 {
		p ^= int(a & 1)
		a >>= 1
	}
	return p
}

func ParityOpt(a int64) int {
	var p int
	for a != 0 {
		p ^= 1
		a &= (a - 1) // Drop lowest 1 bit
	}
	return p
}
