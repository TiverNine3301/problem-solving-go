package cycles

import (
	"fmt"
)

func HandleNumberInput() {
	var number int
	for {
		fmt.Scan(&number)
		if number < 10 {
			continue
		} else if number > 100 {
			break
		}
		fmt.Println(number)
	}
}
