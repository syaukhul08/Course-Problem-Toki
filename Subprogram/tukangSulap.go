package main

import (
	"fmt"
)

func main() {

	var a = []int{}
	var b = []int{}
	var n, m, t, x, y int
	var p, q string

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		a = append(a, m)
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		b = append(b, m)
	}
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Scan(&p, &x, &q, &y)

		if p == "A" && q == "A" {
			a[x-1], a[y-1] = a[y-1], a[x-1]
		} else if p == "A" && q == "B" {
			a[x-1], b[y-1] = b[y-1], a[x-1]
		} else if p == "B" && q == "B" {
			b[x-1], b[y-1] = b[y-1], b[x-1]
		} else if p == "B" && q == "A" {
			b[x-1], a[y-1] = a[y-1], b[x-1]
		}
	}
	for _, bil := range a {
		fmt.Printf("%d ", bil)
	}
	fmt.Println()

	for _, bil := range b {
		fmt.Printf("%d ", bil)
	}
}
