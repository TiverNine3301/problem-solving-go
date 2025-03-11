package problems

import (
	"fmt"
)

func CountMaxNumber() {
	var number int
	var currentMax int = number
	count := 0
	for {
		fmt.Scan(&number)
		if number == 0 {
			break
		}
		if number > currentMax {
			currentMax = number
			count = 1
		} else if number == currentMax {
			count++
		}
	}
	fmt.Println(count)
}
