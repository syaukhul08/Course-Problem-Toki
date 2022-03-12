package main

import (
	"fmt"
	"math"
)

var a, b, k, x float64

func fungsi(x float64) float64 {
	if k == 0 {
		return x
	} else {
		x = math.Abs(a*x + b)
		k--
		return fungsi(x)
	}
}

func main() {
	fmt.Scan(&a, &b, &k, &x)
	fmt.Print(fungsi(x))

}
