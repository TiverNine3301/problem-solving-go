package math_tasks

import "fmt"

func FibonacciIndex() {
	var A int
	fmt.Print("Введите число A (> 1): ")
	if _, err := fmt.Scan(&A); err != nil || A <= 1 {
		fmt.Println("Ошибка: введите число больше 1.")
		return
	}

	a, b := 1, 1
	index := 2
	for b < A {
		a, b = b, a+b
		index++
	}

	if b == A {
		fmt.Println("Число", A, "является", index, "числом Фибоначчи.")
	} else {
		fmt.Println(-1)
	}
}
