// utils/utils.go
package utils

import "fmt"

func PrintMessage() { // Заглавная буква — функция экспортируется
	fmt.Println("Hello from utils!")
}

func printMessage() { // Строчная буква — функция не экспортируется
	fmt.Println("This is private")
}
