package cycles

import (
	"fmt"
	"github.com/TiverNine3301/problem-solving-go/utils"
)

func FindCommonDigits() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	digits1 := utils.GetDigits(num1)
	digits2 := utils.GetDigits(num2)

	intersection := utils.FindIntersection(digits1, digits2)

	for _, digit := range intersection {
		fmt.Printf("%d ", digit)
	}
}
