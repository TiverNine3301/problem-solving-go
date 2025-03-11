package problems

import "fmt"

func DeclineCow() {
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
