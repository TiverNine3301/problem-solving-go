package main

import (
	"fmt"
	"problems"
)

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
