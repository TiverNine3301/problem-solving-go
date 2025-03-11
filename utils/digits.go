package utils

import "fmt"

// GetDigits возвращает множество цифр числа.
func GetDigits(num int) map[int]struct{} {
	digits := make(map[int]struct{})
	strNum := fmt.Sprintf("%d", num)
	for _, digit := range strNum {
		digits[int(digit-'0')] = struct{}{}
	}
	return digits
}

// FindIntersection возвращает пересечение двух множеств цифр.
func FindIntersection(set1, set2 map[int]struct{}) []int {
	var intersection []int
	for digit := range set1 {
		if _, ok := set2[digit]; ok {
			intersection = append(intersection, digit)
		}
	}
	return intersection
}
