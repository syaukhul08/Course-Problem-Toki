package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
