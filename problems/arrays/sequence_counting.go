package arrays

import "fmt"

func CountNonNegatives() {
	var n int8
	var count int
	fmt.Scan(&n)

	for i := int8(0); i < n; i++ {
		var number int
		fmt.Scan(&number)

		if number >= 0 {
			count++
		}
	}
	fmt.Println(count)
}
