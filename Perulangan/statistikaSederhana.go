package main

import "fmt"

func main() {
	var n, max, min int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var number int
		fmt.Scanf("%d", &number)
		if i == 0 {
			max = number
			min = number
		}
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	fmt.Printf("%d %d", max, min)
}
