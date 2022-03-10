package main

import (
	"fmt"
	"math"
)

func fungsi(a, b, k, x float64) float64 {
	if k == 0 {
		return x
	} else {
		return math.Abs(a*fungsi(a, b, k-1, x) + b)
	}
}

func main() {
	var a, b, k, x float64
	fmt.Scan(&a, &b, &k, &x)
	fmt.Print(fungsi(a, b, k, x))
}
