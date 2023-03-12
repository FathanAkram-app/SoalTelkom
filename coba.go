package main

import "fmt"

func main() {
	// var berenti, makanan string
	// var lompatan int

	// for true {
	// 	fmt.Println("mau berenti?")
	// 	fmt.Scan(&berenti)
	// 	if berenti == "iya" {
	// 		// keluar dari loop
	// 		break
	// 	}
	// 	// user ditanya ingin makan apa
	// 	fmt.Println("mau makan apa?")
	// 	// input makanan si user
	// 	fmt.Scan(&makanan)
	// 	// user makan makanan
	// 	makan(makanan)

	// 	fmt.Println("mau lompat seberapa jau? (cm)")
	// 	fmt.Scan(&lompatan)
	// 	lompat(lompatan)

	// }

	var x int
	x = 12
	fmt.Println(x)
	tambahsatu(&x)
	fmt.Println(x)

}

func tambahsatu(x *int) {
	*x += 1
}

// fungsi makanan
func makan(makanan string) {
	fmt.Println("anda telah memakan " + makanan)

}

// fungsi lompat
func lompat(lompatan int) {
	lompatSekarang := fmt.Sprintf("Anda melompat sejauh %dcm", lompatan)
	fmt.Println(lompatSekarang)
}
