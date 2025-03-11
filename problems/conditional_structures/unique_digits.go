package problems

import "fmt"

func AreDigitsUnique() {
	var number int
	fmt.Scan(&number)

	var first int = number / 100
	var mid int = (number / 10) % 10
	var last int = number % 10

	switch {
	case first == mid:
		fmt.Println("NO")
	case first == last:
		fmt.Println("NO")
	case mid == last:
		fmt.Println("NO")
	default:
		fmt.Println("YES")
	}
}
