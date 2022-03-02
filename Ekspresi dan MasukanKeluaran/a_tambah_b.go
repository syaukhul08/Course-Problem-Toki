package main

import (
	"fmt"
)

func main() {

	var a, b int
	c := 0

	fmt.Print("Masukkan A : ")
	fmt.Scan(&a)

	fmt.Print("Masukkan B : ")
	fmt.Scan(&b)

	c = a + b
	fmt.Print("A + B = ", c)
}
