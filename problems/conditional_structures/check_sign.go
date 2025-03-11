package conditional_structures

import "fmt"

func ShowNumberSign() {
	var digit int
	fmt.Print("Введите любое число: ")
	fmt.Scan(&digit)

	switch {
	case digit < 0:
		fmt.Println("Число отрицательное")
	case digit > 0:
		fmt.Println("Число положительное")
	case digit == 0:
		fmt.Println("Ноль")
	}
}
