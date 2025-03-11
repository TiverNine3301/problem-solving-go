package problems

import "fmt"

func Max_multiple_of_7() {
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
