package p50

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	power := int64(n)
	if power < 0 {
		x = 1 / x
		power = -power
	}

	res := 1.0
	for power > 0 {
		if power&1 == 1 {
			res *= x
		}
		x *= x
		power >>= 1
	}
	return res
}
