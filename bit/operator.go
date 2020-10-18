package bit

func Add(a, b int64) int64 {
	var sum int64
	k := int64(1)
	ta := a
	tb := b
	carrying := int64(0)
	for ta != 0 || tb != 0 {
		ka := ta & k
		kb := tb & k
		sum |= ka ^ kb ^ carrying
		carrying = ((ka & kb) | (ka & carrying) | (kb & carrying)) << 1
		k <<= 1
		ta <<= 1
		tb <<= 1
	}
	return sum
}
