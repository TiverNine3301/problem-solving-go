package conditional_structures

import "fmt"

func GetFirstDigit() {
	var number int
	fmt.Scan(&number)

	for number >= 10 {
		number /= 10
	}

	fmt.Println(number)
}
