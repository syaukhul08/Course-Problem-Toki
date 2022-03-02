package main

import (
	"fmt"
)

func main() {
	var a, t, luas float64
	fmt.Scanf("%f %f", &a, &t)
	luas = a * t / 2
	fmt.Printf("%.2f", luas)
}
