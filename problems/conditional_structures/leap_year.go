package conditional_structures

import "fmt"

func IsYearLeap() {
	var number int
	fmt.Scan(&number)

	if number%400 == 0 {
		fmt.Println("YES")
	} else if number%4 == 0 && number%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
