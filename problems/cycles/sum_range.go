package cycles

import "fmt"

func GetRangeSum() {
	var numberFirst int
	var numberLast int
	var result int
	fmt.Scan(&numberFirst, &numberLast)

	for i := numberFirst; i < numberLast+1; i++ {
		result = numberFirst + result
		numberFirst++
	}
	fmt.Println(result)
}
