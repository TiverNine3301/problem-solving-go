package cycles

import "fmt"

func ShowSquares() {
	for n := 1; n < 11; n++ {
		number := n * n
		fmt.Println(number)
	}
}
