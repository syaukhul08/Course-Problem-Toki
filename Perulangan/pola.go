package main

import "fmt"

func main() {
	var n, pola int
	fmt.Scanf("%d %d", &n, &pola)
	for i := 1; i <= n; i++ {
		if i%pola == 0 {
			fmt.Printf(" * ")
		} else {
			fmt.Print(" ", i)
		}
	}
}
