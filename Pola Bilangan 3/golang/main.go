package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	for i := 1; i < a+1; i++ {
		for h := 1; h < a+1; h++ {
			if h == i || h == (a+1)-i {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
