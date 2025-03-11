package arrays

import "fmt"

func FilterEvenIndices() {

	var n int8

	fmt.Scan(&n)

	array := make([]int8, 0, n)

	var nubmer int8

	for i := int8(0); i < n; i++ {
		fmt.Scan(&nubmer)
		array = append(array, nubmer)
	}

	var ambulant_index int8 = 2

	for i := 0; i < len(array); i++ {
		if int8(i)%ambulant_index == 0 {
			fmt.Print(array[i], " ")
		}
	}

}
