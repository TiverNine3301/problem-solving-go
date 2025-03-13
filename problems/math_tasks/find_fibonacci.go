package math_tasks

import "fmt"

func FibonacciIndex(A int) int {
	if A < 0 {
		fmt.Print(-1)
	}

	a, b := 0, 1
	index := 0

	for a <= A {
		if a == A {
			return index
		}
		a, b = b, a+b
		index++
	}
	return -1
}
