package main

import "fmt"

func main() {
	var n, terkecil, terbesar int
	terkecil = 2147483647
	terbesar = -2147483647

	for true {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if n < terkecil {
			terkecil = n
		}

		if n > terbesar {
			terbesar = n
		}
	}
	fmt.Print(fmt.Sprintf("%d %d", terbesar, terkecil))
}
