package problems

import "fmt"

func IsHappyTicket() {
	var number int
	fmt.Scan(&number)

	var one int = number / 100000
	var two int = (number / 10000) % 10
	var three int = (number / 1000) % 10
	var four int = (number / 100) % 10
	var five int = (number / 10) % 10
	var six int = number % 10

	result1 := one + two + three
	result2 := four + five + six

	if result1 == result2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
