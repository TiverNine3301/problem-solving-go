package main

import (
	"fmt"

	"problem-solving-go/problems"
)

func main() {
	fmt.Print("Введите номер задачи: ")
	var task_number int
	fmt.Scan(&task_number)

	switch task_number {
	case 1:
		problems.Minimum()
	case 2:
		var n int
		fmt.Print("Введите число для вычисления цифрового корня: ")
		if _, err := fmt.Scan(&n); err != nil {
			fmt.Println("Ошибка: введите корректное число.")
			return
		}
		fmt.Println("Цифровой корень:", problems.Digital_root(n))
	case 3:
		problems.Max_multiple_of_7()
	case 4:
		problems.DeclineCow()
	case 5:
		problems.ShowPowersOfTwo()
	case 6:
		problems.FibonacciIndex()
	case 7:
		problems.ShowNumberSign()
	case 8:
		problems.AreDigitsUnique()
	case 9:
		problems.GetFirstDigit()
	case 10:
		problems.IsHappyTicket()
	case 11:
		problems.IsYearLeap()
	case 12:
		problems.ShowSquares()
	case 13:
		problems.GetRangeSum()
	case 14:
		problems.SumEligibleNumbers()
	case 15:
		problems.CountMaxNumber()
	case 16:
		problems.FindFirstMatch()
	case 17:
		problems.HandleNumberInput()
	case 18:
		problems.YearsToReachTarget()
	case 19:
		problems.FindCommonDigits()
	default:
		fmt.Print("Введите число.")
	}

}
