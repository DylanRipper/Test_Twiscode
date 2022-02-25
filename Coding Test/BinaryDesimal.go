package main

import (
	"fmt"
	"math"
)

func main() {
	var bin int
	var dec int
	var menu int
	fmt.Print(
		"===========================\n",
		"Pilih Menu\n",
		"1. Binary to Decimal\n",
		"2. Decimal to Binary\n",
		"===========================\n")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		fmt.Print("Input Binary Number : ")
		fmt.Scanln(&bin)
		var save int
		count := 0.0
		for bin != 0 {
			save = bin % 10
			dec += save * int(math.Pow(2.0, count))
			bin = bin / 10
			count++
		}
		fmt.Printf("Hasil : %d", dec)

	case 2:
		fmt.Print("Input Decimal Number : ")
		fmt.Scan(&dec)
		var save int
		count := 1
		for dec != 0 {
			save = dec % 2
			dec = dec / 2
			bin += save * count
			count *= 10
		}
		fmt.Printf("Hasil : %d", bin)
	}
}
