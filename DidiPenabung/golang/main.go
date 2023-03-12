package main

import "fmt"

func main() {

	var uangMasuk, tertinggi, terendah, jumlah int
	var repeat bool
	repeat = true
	i := 0
	for repeat {
		fmt.Scan(&uangMasuk)
		i += 1
		jumlah += uangMasuk
		if uangMasuk < terendah {
			terendah = uangMasuk
		}

		if uangMasuk > tertinggi {
			tertinggi = uangMasuk
		}
	}
	fmt.Println("Jumlah = ", jumlah)
	fmt.Println("Rata - rata = ", jumlah/i)
	fmt.Println("Uang tertinggi = ", tertinggi)
	fmt.Println("Uang terendah = ", terendah)
}
