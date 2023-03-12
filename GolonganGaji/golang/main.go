package main

import "fmt"

func main() {
	var jam, total, jamLembur, gajiTemp int
	var golongan string
	var repeat bool
	repeat = true
	for repeat {
		fmt.Scan(&golongan, &jam, &jamLembur)
		if golongan == "A" {
			gajiTemp = (75000 * jam) + (90000 * jamLembur)
			total += gajiTemp
			fmt.Println(gajiTemp)
		} else if golongan == "B" {
			gajiTemp = (125000 * jam) + (70000 * jamLembur)
			total += gajiTemp
			fmt.Println(gajiTemp)
		} else if golongan == "Z" {
			fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan ", total)
			repeat = false
		} else {
			fmt.Println("Golongan tidak dikenali")
		}

	}
}
