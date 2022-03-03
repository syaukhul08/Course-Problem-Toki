package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n == 1 {
		fmt.Println("ya")
	}

	for n%2 == 0 {
		n = n / 2

		if n == 1 {
			fmt.Println("ya")
			break
		}
	}

	if n != 1 {
		fmt.Println("bukan")
	}
}
