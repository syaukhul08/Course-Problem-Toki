package main

import (
	"fmt"
)

func main() {
	var n, s int
	var arr [1001]int

	for i := 0; i < 1001; i++ {
		arr[i] = 0
	}

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		arr[s]++
	}

	max := 0
	modus := 0

	for i := 0; i <= 1000; i++ {
		if max <= arr[i] && modus < i {
			max = arr[i]
			modus = i
		}
	}

	fmt.Println(modus)
}
