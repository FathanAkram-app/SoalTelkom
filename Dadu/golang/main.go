package main

import "fmt"

func main() {
	var daduSatu, daduDua, daduTiga, sum, ganjil int

	for sum != 10 {
		fmt.Scan(&daduSatu)
		fmt.Scan(&daduDua)
		fmt.Scan(&daduTiga)
		sum = (daduSatu + daduDua + daduTiga)
		if daduSatu%2 != 0 && daduDua%2 != 0 && daduTiga != 0 {
			ganjil++
		}
	}
	fmt.Println(ganjil)
}
