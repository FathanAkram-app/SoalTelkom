package main

import "fmt"

func main() {
	totalBelanja := 0
	statusMembership := ""

	fmt.Scan(&totalBelanja, &statusMembership)
	if statusMembership == "Yes" {
		if totalBelanja > 10000 {
			fmt.Println("Anda mendapat discount 5%, tambahan poin 10, dan free gift")

		} else if totalBelanja > 75000 {
			fmt.Println("Anda mendapat discount 5%", " dan tambahan poin 5")

		} else {
			fmt.Println("Anda mendapat tambahan poin 5")
		}

	}
}
