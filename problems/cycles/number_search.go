package cycles

import "fmt"

func FindFirstMatch() {
	var n, c, d int

	fmt.Scan(&n, &c, &d)
	for i := 0; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
