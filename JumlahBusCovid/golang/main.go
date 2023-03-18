package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(jumlahBus(n, m), "bus")
}

func jumlahBus(penumpang, kapasitas int) int {
	bus := penumpang / kapasitas
	if penumpang%kapasitas > 0 {
		bus++
	}
	return bus
}
