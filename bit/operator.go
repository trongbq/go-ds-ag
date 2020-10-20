package bit

func Add(a, b uint64) uint64 {
	var sum uint64
	k := uint64(1)
	carrying := uint64(0)
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
func Multiply(a, b uint64) uint64 {
	var mul uint64
	for a != 0 {
		if (a & 1) == 1 {
			mul = Add(mul, b)
		}
		a >>= 1
		b <<= 1
	}
	return mul
}

func Divide(a, b uint64) uint64 {
	var quotient uint64
	// should be 64, but then if b > 1, then bk is larger than uint64, so 32 for running purpose
	k := 32
	bk := b << k
	for a >= b {
		// try to find max `k` that 2^k < a, decreasing...
		for bk > a {
			bk >>= 1
			k--
		}

		quotient += 1 << k
		a -= bk
	}
	return quotient
}

// if b is even then a^b = a^(b/2)^2 = a^(b/4)^2^2 = ...
// if b is odd then a^b = a * a^(b-1)
func Power(a, b uint64) uint64 {
	p := uint64(1)
	for b != 0 {
		if (b & 1) == 1 {
			// two cases
			// 1. when b is an odd number (on each iteration)
			// 2. when b reaches to 1
			p *= a
		}
		a *= a
		b >>= 1
	}
	return p
}
