package main

import (
	"fmt"
)

//Количество минимумов
// Найдите количество минимальных элементов в последовательности.

// Входные данные

// Вводится натуральное число N, а затем N целых чисел последовательности.

// Выходные данные

// Выведите количество минимальных элементов последовательности.

func solution1() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	minValue := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
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

	fmt.Println(count)
}

func main() {
	solution1()
}
