package calc

// Mul in a slow version
func Mul(a, b int) int {
	acc := 0
	for i := 0; i < b; i++ {
		acc = Add(acc, a)
	}

	return acc
}
