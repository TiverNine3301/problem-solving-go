package problems

import (
	"fmt"
)

func YearsToReachTarget() {
	var x, p, y int
	count := 0
	fmt.Scan(&x, &p, &y)
	for {
		if x < y {
			x += (x * p) / 100
			count++
			continue
		}
		fmt.Println(count)
		break
	}
}
