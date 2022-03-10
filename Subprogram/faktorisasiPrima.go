package main

import "fmt"

func faktor(n int) {
	for i := 2; i <= n; i++ {
		sum := 0
		for n%i == 0 {
			sum++
			n /= i
		}
		if sum > 1 {
			fmt.Printf("%d^%d", i, sum)
		}
		if sum == 1 {
			fmt.Printf("%d", i)
		}
		dev := n / i
		if sum != 0 && dev != 0 {
			fmt.Printf(" x ")
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	faktor(n)
}
