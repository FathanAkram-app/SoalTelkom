package main

// one line input
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	var a int
// 	var s string
// 	var h [3]int

// 	scanner := bufio.NewScanner(os.Stdin)
// 	if scanner.Scan() {
// 		s = scanner.Text()
// 		a = len(s)
// 		for i := 0; i < (a/2)+1; i++ {
// 			if s[i*2] == 97 {
// 				h[0] += 1
// 			} else if s[i*2] == 98 {
// 				h[1] += 1
// 			} else if s[i*2] == 99 {
// 				h[2] += 1
// 			}

// 		}
// 		fmt.Println("Tipe A = ", h[0])
// 		fmt.Println("Tipe B = ", h[1])
// 		fmt.Println("Tipe C = ", h[2])

// 	}

// }

import "fmt"

func main() {
	var a string
	var countA, countB, countC int
	for true {
		fmt.Scan(&a)
		if a == "A" {
			countA++
		} else if a == "B" {
			countB++
		} else if a == "C" {
			countC++
		} else {
			break
		}
	}
	fmt.Println("Tipe A = ", countA)
	fmt.Println("Tipe B = ", countB)
	fmt.Println("Tipe C = ", countC)
}
