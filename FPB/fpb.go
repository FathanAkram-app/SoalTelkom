package main

import "fmt"

func main() {
	// code by Fauzein, Shua
	var num1, num2, bnum, bfact uint
	fmt.Printf("Faktor Persekutuan Terbesar (FPB)\n\n")
	fmt.Print("Angka yang difaktorkan (2) : ")
	fmt.Scan(&num1, &num2)
	if num1 < num2 {
		bnum = num2
	} else {
		bnum = num1
	}
	for i := uint(1); i < bnum; i++ {
		if num1%i == 0 && num2%i == 0 {
			bfact = i
		}
	}
	fmt.Println(bfact)
}
