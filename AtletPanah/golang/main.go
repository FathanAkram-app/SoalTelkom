package main

import "fmt"

func main() {
	var b1, b2, b3, h1, h2, h3, terbesar int
	fmt.Scan(&b1, &b2, &b3)
	for b1 != 0 && b2 != 0 && b3 != 0 {
		h1 += b1
		h2 += b2
		h3 += b3
		fmt.Scan(&b1, &b2, &b3)

	}
	if h1 > h2 && h1 > h3 {
		terbesar = h1
	} else if h2 > h1 && h2 > h3 {
		terbesar = h2
	} else if h3 > h2 && h3 > h1 {
		terbesar = h3
	}
	fmt.Print(h1, " ", h2, " ", h3, " ", terbesar)

}
