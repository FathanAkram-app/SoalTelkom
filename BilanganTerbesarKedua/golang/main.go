package main

import "fmt"

func main() {
	var n, temp1, temp2 int
	fmt.Scan(&n)
	for n != 0 {
		if n > temp1 {
			temp2 = temp1
			temp1 = n
		} else if temp2 < n {
			temp2 = n
		}
		fmt.Scan(&n)

	}
	fmt.Println(temp2)
}
