package math

func AbsInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func AbsFloat64(x float64) float64 {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
