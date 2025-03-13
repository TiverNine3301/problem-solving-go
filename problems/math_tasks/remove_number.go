package math_tasks

import (
	"fmt"
)	

func Delete_num() {
	var number, digit int
	fmt.Scan(&number)
	fmt.Scan(&digit)

	multiplier := 1
	result := 0

	for number
}

// var number int
// var digit int

// // Ввод числа и цифры
// fmt.Scan(&number)
// fmt.Scan(&digit)

// result := 0
// multiplier := 1

// // Обрабатываем каждую цифру числа
// for number > 0 {
// 	currentDigit := number % 10 // Получаем последнюю цифру
// 	number = number / 10        // Убираем последнюю цифру из числа

// 	// Если цифра не равна удаляемой, добавляем её в результат
// 	if currentDigit != digit {
// 		result += currentDigit * multiplier
// 		multiplier *= 10
// 	}
// }

// // Выводим результат
// fmt.Println(result)
