package main

import (
	"fmt"
)

func main() {

	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println("masing-masing ", n/m)
	fmt.Println("bersisa ", n%m)
}
