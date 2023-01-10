package main

import "fmt"

func main() {
	var gg string
	var n, m int
	fmt.Scan(&gg)
	fmt.Scan(&n)
	fmt.Scan(&m)
	for i := n; i <= m; i++ {
		if gg == "genap" {
			if i%2 == 0 {
				fmt.Print(fmt.Sprintf("%d ", i))
			}
		} else {
			if i%2 != 0 {
				fmt.Print(fmt.Sprintf("%d ", i))
			}
		}
	}
}
