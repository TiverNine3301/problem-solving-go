package problems

import (
	"fmt"
)

func FindCommonDigits() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	digits1 := getDigits(num1)
	digits2 := getDigits(num2)

	intersection := findIntersection(digits1, digits2)

	for _, digit := range intersection {
		fmt.Printf("%d ", digit)
	}
}

func getDigits(num int) map[int]struct{} {
	digits := make(map[int]struct{})
	strNum := fmt.Sprintf("%d", num)
	for _, digit := range strNum {
		digits[int(digit-'0')] = struct{}{}
	}
	return digits
}

func findIntersection(set1, set2 map[int]struct{}) []int {
	var intersection []int
	for digit := range set1 {
		if _, ok := set2[digit]; ok {
			intersection = append(intersection, digit)
		}
	}
	return intersection
}
