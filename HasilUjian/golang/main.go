package amin

import "fmt"

func main() {
	var n int
	fmt.Println("Berapa jumlah siswa yang nilainya akan diproses?")
	fmt.Scan(&n)

	n1 := 0
	n2 := 0
	n3 := 0
	n_passed := 0
	n_failed := 0
	avg := 0
	for i := 1; i < n; i++ {
		fmt.Scan(n1, n2, n3)
		avg = (n1 + n2 + n3) / 3
		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed = n_passed + 1
		} else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed = n_failed + 1
		}
	}

	fmt.Println("Jumlah siswa lolos seleksi admistrasi", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi admistrasi", n_failed)
}
