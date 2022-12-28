package main

import (
	"fmt"
)

func main() {
	var a, terkecil, terbesar, ratarata, t, c int
	terkecil = 2147483647
	terbesar = -2147483647
	ratarata = 0
	a = 1
	c = 0
	for true {
		fmt.Scan(&t)
		if t < terkecil {
			terkecil = t
		}

		if t > terbesar {
			terbesar = t
		}
		ratarata += t

		if a == 0 && t == 0 {
			break
		}
		c++

		a = t

	}
	fmt.Println(terbesar)
	fmt.Println(terkecil)
	fmt.Println(float32(ratarata) / (float32(c)))

}
