package main

import (
	"fmt"
)

var str string
var n int

func main() {
	fmt.Scan(&str)
	fmt.Scan(&n)
	for i := 0; i < len(str); i++ {
		charNow := int(str[i] - 97)
		charNext := (charNow+n)%26 + 97
		fmt.Print(string(charNext))
	}
}
