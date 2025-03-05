package main

import (
	"fmt"
)

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

// solution3 находит наибольшее число на отрезке [a, b], кратное 7.
func solution3() {
	var a, b int
	fmt.Print("Введите два числа (a и b, где a <= b): ")
	if _, err := fmt.Scan(&a, &b); err != nil || a > b {
		fmt.Println("Ошибка: введите корректные числа.")
		return
	}

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Println("Наибольшее число, кратное 7:", i)
			return
		}
	}
	fmt.Println("NO")
}

// solution4 склоняет слово "корова" в зависимости от числа n.
func solution4() {
	var num int
	fmt.Print("Введите число n (0 < n < 100): ")
	if _, err := fmt.Scan(&num); err != nil || num <= 0 || num >= 100 {
		fmt.Println("Ошибка: введите число от 1 до 99.")
		return
	}

	var word string
	switch {
	case num%10 == 1 && num%100 != 11:
		word = "korova"
	case num%10 >= 2 && num%10 <= 4 && (num%100 < 10 || num%100 >= 20):
		word = "korovy"
	default:
		word = "korov"
	}

	fmt.Printf("%d %s\n", num, word)
}

// solution5 выводит все степени двойки, не превышающие N.
func solution5() {
	var N int
	fmt.Print("Введите число N: ")
	if _, err := fmt.Scan(&N); err != nil || N <= 0 {
		fmt.Println("Ошибка: введите корректное натуральное число.")
		return
	}

	degree := 1
	for degree <= N {
		fmt.Print(degree, " ")
		degree *= 2
	}
	fmt.Println()
}

// solution6 определяет, каким по счету числом Фибоначчи является число A.
func solution6() {
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

func main() {
	solution1()
	solution3()
	solution4()
	solution5()
	solution6()

	// Для solution2
	var n int
	fmt.Print("Введите число для вычисления цифрового корня: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: введите корректное число.")
		return
	}
	fmt.Println("Цифровой корень:", solution2(n))
}
