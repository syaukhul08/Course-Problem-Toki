package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	if number > 0 {
		fmt.Println("positif")
	} else if number == 0 {
		fmt.Println("nol")
	} else {
		fmt.Println("negatif")
	}
}
