package bit

func Add(a, b int64) int64 {
	var sum int64
	k := int64(1)
	carrying := int64(0)
	ta := a
	tb := b
	for ta != 0 || tb != 0 {
		ka := a & k
		kb := b & k
		sum |= ka ^ kb ^ carrying

		carrying = ((ka & kb) | (ka & carrying) | (kb & carrying)) << 1
		k <<= 1
		ta >>= 1
		tb >>= 1
	}
	return sum | carrying
}

// a * b = (a)2 * b = (... + (0|1)*2^1 + (0|1)*2^0) * b
func Multiply(a, b int64) int64 {
	var mul int64
	for a != 0 {
		if (a & 1) == 1 {
			mul = Add(mul, b)
		}
		a >>= 1
		b <<= 1
	}
	return mul
}
