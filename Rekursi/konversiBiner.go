package main

import "fmt"

var n int

func konversiBiner(n int) string {
	if n == 1 {
		return "1"
	} else if n%2 == 1 {
		return konversiBiner(n/2) + "1"
	} else {
		return konversiBiner(n/2) + "0"
	}
}

func main() {
	fmt.Scan(&n)
	fmt.Println(konversiBiner(n))
}
