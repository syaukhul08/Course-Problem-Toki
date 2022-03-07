package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x int

	number := make([]int, 0)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		number = append(number, x)
	}

	for _, num := range number {
		prima := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				prima = false
			}
		}
		if num == 1 || prima == false {
			fmt.Println("BUKAN")
		} else {
			fmt.Println("YA")
		}
	}
}
