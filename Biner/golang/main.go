package main

import (
	"fmt"
)

func main() {
	var a string
	var num int
	fmt.Scan(&num)
	for num != 0 {
		a += fmt.Sprintf("%d", num%2)
		num = num / 2
	}
	if len(a) == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(a)
	}
}
