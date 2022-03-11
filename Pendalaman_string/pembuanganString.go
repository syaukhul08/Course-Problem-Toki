package main

import (
	"fmt"
	"strings"
)

var s, t string

func main() {
	fmt.Scan(&s, &t)
	for strings.Contains(s, t) {
		s = strings.Replace(s, t, "", 1)
	}
	fmt.Println(s)
}
