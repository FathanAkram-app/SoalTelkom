package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for i := 1; i < a+1; i++ {
		if a%i == 0 {
			b += 1
		}

	}
	if b > 2 {
		fmt.Println(false)
	} else if b <= 2 {
		fmt.Println(true)
	}

}
