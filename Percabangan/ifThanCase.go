package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scan(&n)
	switch len(n) {
	case 1:
		fmt.Println("satuan")
	case 2:
		fmt.Println("puluhan")
	case 3:
		fmt.Println("ratusan")
	case 4:
		fmt.Println("ribuan")
	case 5:
		fmt.Println("puluhribuan")
	}

}
