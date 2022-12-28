package main

import "fmt"

func main() {
	var a, b int
	var h bool
	fmt.Scan(&a)
	fmt.Scan(&b)
	h = false
	for b != 0 {
		t := b % 10
		if t == a {
			h = true
			break
		}
		b = (b - t) / 10
	}
	fmt.Println(h)
}
