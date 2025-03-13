package math_tasks

import "fmt"

func ToBinary() {
	var N int
	fmt.Scan(&N)

	if N == 0 {
		fmt.Println(0)
		return
	}

	arr := []int{}

	for N > 0 {
		remainder := N % 2
		arr = append(arr, remainder)
		N = N / 2
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}
}
