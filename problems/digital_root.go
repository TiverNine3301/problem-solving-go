package problems

// solution2 вычисляет цифровой корень числа.
func solution2(n int) int {
	if n == 0 {
		return 0
	}
	if n%9 == 0 {
		return 9
	}
	return n % 9
}
