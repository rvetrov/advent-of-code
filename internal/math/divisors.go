package math

func posGCD(a, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	return posGCD(b, a%b)
}

func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return posGCD(a, b)
}

func LCM(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a / GCD(a, b) * b
}
