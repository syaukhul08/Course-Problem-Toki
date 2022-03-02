package main

import (
	"fmt"
)

func main() {
	var n float32
	fmt.Scan(&n)

	a := int(n)
	x := float32(a)

	if n == x {
		fmt.Println(int(n), int(n))
	} else if n > 0 {
		fmt.Println(int(n), int(n+1))
	} else if n < 0 {
		fmt.Println(int(n-1), int(n))
	}
}
