package arrays

import (
	"fmt"
)

func ProcessSlice() {
	var n int
	fmt.Scan(&n)

	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	fmt.Println(b[3])
}
