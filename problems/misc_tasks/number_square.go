package misk_tasks

import (
	"fmt"
)

func ProcessNumber() {
	var number float64
	fmt.Scan(&number)

	for {
		if number <= 0 {
			fmt.Printf("число %2.2f не подходит\n", number)
			break
		} else if number > 10000 {
			fmt.Printf("%e\n", number)
			break
		}

		result := (number * number)
		fmt.Printf("%.4f\n", result)
		break
	}
}
