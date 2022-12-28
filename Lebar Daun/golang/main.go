package main

import "fmt"

func main() {
	var n, num, hasil int
	hasil = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if hasil < num {
			hasil = num
		}
	}
	fmt.Println(hasil)
}
