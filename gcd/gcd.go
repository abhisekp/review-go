package gcd

import "math"

func GCD(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	if b == 0 {
		return a
	}

	if a == 0 {
		return b
	}

	high := max(a, b)
	low := min(a, b)

	return GCD(low, high%low)
}
