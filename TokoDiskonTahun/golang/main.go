package main

import "fmt"

func main() {

	var tahun, harga int
	fmt.Scan(&tahun)
	fmt.Scan(&harga)

	cd := tahun % 100
	a := (tahun - cd) / 1000
	fmt.Print("Besar diskon: ", a*cd, "%\n")

	fmt.Println("Jumlah yang dibayar: ", harga-int(float64(harga)*float64(a*cd)/100))

}
