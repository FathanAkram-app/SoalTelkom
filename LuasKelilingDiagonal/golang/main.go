package main

import (
	"fmt"
	"math"
)

func main() {
	var panjang, lebar float64
	fmt.Scan(&panjang)
	fmt.Scan(&lebar)
	fmt.Println("Luas: ", int(panjang*lebar))
	fmt.Println("Keliling: ", int(panjang+lebar+panjang+lebar))
	fmt.Println("Panjang Diagonal: ", int(math.Sqrt((panjang*panjang)+(lebar*lebar))))
}
