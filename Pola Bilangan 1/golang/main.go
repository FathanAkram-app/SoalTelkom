package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		for h := 1; h < a+1; h++ {
			fmt.Print(h)
		}
		fmt.Println()
	}
}
