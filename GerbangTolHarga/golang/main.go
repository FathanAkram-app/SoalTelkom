package main

import "fmt"

func main() {
	var a string
	var countA, countB, countC, price, j int
	price = 0
	fmt.Scan(&j)

	for i := 0; i < j; i++ {
		fmt.Scan(&a)
		if a == "A" {
			price += 10000
			countA++
		} else if a == "B" {
			price += 23000
			countB++
		} else if a == "C" {
			price += 45000
			countC++
		} else {
			break
		}
	}
	fmt.Print(fmt.Sprintf("%d %d %d %d", price, countA, countB, countC))

}
