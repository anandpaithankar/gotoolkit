package math

// MinInt ... Returns a min of two integers.
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MaxInt ... Returns a max of two integers.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
