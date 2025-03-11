package problems

import "fmt"

func ShowPowersOfTwo() {
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
