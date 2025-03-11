package arrays

import "fmt"

func ProcessArray() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		var num uint8
		fmt.Scan(&num)
		workArray[i] = uint8(num)
	}

	for i := 0; i < 3; i++ {
		var idx1, idx2 int
		fmt.Scan(&idx1, &idx2)
		workArray[idx1], workArray[idx2] = workArray[idx2], workArray[idx1]
	}

	for _, num := range workArray {
		fmt.Printf("%d ", num)
	}
}
