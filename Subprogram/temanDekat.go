package main

import (
	"fmt"
)

func check(a, b, c int) int {
	var x int
	x = (a - b)
	if x < 0 {
		x *= -1
	}
	y := 1
	for i := 1; i <= c; i++ {
		y *= x
	}
	return y
}

func main() {
	var x, d int
	fmt.Scanf("%d %d", &x, &d)
	var arr1 [1025]int
	var arr2 [1025]int
	q := 0
	for i := 1; i <= x; i++ {
		fmt.Scanf("%d %d", &arr1[i], &arr2[i])
	}

	p := check(arr1[1], arr1[2], d) + check(arr2[1], arr2[2], d)

	for i := 1; i <= x; i++ {
		for j := i + 1; j <= x; j++ {
			a := check(arr2[i], arr2[j], d) + check(arr1[i], arr1[j], d)
			if a > q {
				q = a
			} else if a < p {
				p = a
			}
		}
	}
	fmt.Printf("%d %d", p, q)
	return
}
