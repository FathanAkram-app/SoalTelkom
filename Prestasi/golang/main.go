package main

import "fmt"

func main() {
	var predicate string

	total1, total2, total3, repeat := 0, 0, 0, true

	for repeat {
		fmt.Scan(&predicate)
		if predicate == "Sufficient" {
			total1 += 1
		} else if predicate == "Good" {
			total2 += 1
		} else if predicate == "Excellent" {
			total3 += 1
		}
		if predicate == "STOP" {
			repeat = false
		}
	}
}
