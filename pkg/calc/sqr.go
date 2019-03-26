package calc

import "math"

// Sqr multiples number with itself
func Sqr(a int) int {
	return Mul(a, a)
}

// Sqr multiples number with itself
func SqrNative(a int) int {
	return a * a
}

// Sqr multiples number with itself
func SqrMath(a int) int {
	return int(math.Pow(float64(a), 2.0))
}
