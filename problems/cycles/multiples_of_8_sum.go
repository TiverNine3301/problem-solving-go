package problems

import "fmt"

func SumEligibleNumbers() {
	var size int
	var sum int
	var number int

	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		fmt.Scan(&number)

		if number >= 10 && number <= 99 && number%8 == 0 {
			sum += number
		}
	}
	fmt.Println(sum)
}
