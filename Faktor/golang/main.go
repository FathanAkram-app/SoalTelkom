package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := 1; i < a+1; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
