func mySqrt(n int) int {
	if n == 0 {
		return n
	}
	x := 1 << ((bits.Len32(uint32(n)) + 1) / 2)
	for {
		next := (x + n/x) >> 1
		if next >= x {
			return x
		}
		x = next
	}
}