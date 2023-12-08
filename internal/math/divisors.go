package math

func GCD(a, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
