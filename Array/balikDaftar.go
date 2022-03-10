package main

import "fmt"

func main() {
	var arr [100]int

	n := 0
	for {
		var x int
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		arr[n] = x
		n++
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
