package main

import "fmt"

func main() {
	var a, i int
	fmt.Scan(&a)
	i = 0
	for a != 0 {
		t := a % 10
		if i < t {
			i = t
		}
		a = (a - t) / 10
	}
	fmt.Println(i)
}
