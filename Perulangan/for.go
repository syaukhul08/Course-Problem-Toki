package main

import (
	"fmt"
)

func main() {
	var awal, bil, i int
	jumlah := 0

	fmt.Scan(&awal)
	for i = 0; i < awal; i++ {
		fmt.Scan(&bil)
		jumlah += bil
	}
	fmt.Println(jumlah)
}
