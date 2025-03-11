package problems

import "fmt"

func Minimum() {
	var n int
	fmt.Print("Введите количество чисел: ")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Ошибка: введите корректное натуральное число.")
		return
	}

	numbers := make([]int, n)
	minValue := 0
	fmt.Print("Введите числа: ")
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&numbers[i]); err != nil {
			fmt.Println("Ошибка: введите корректное число.")
			return
		}
		if i == 0 || numbers[i] < minValue {
			minValue = numbers[i]
		}
	}

	count := 0
	for _, num := range numbers {
		if num == minValue {
			count++
		}
	}

	fmt.Println("Количество минимальных элементов:", count)
}
