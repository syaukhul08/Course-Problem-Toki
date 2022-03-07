package main

import "fmt"

func main() {
	var n, x int

	number := make([]int, 0)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		number = append(number, x)
	}

	for _, num := range number {
		y := 0
		for i := 2; i < (num / 3); i++ {
			if num%i == 0 {
				y += 1
			}
			if y > 2 {
				break
			}
		}
		if num == 1 || y <= 2 {
			fmt.Println("YA")
		} else {
			fmt.Println("BUKAN")
		}
	}
}
