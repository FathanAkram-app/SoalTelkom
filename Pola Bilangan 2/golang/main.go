package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := 1; i < a+1; i++ {
		for h := 0; h < a; h++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
