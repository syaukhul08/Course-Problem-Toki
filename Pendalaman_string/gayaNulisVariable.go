package main

import (
	"fmt"
	"strconv"
	"strings"
)

var a, res string
var x []string

func gayaNulis(word string) string {
	if strings.Contains(a, "_") {
		x = strings.Split(a, "_")
		for id, word := range x {
			if id != 0 {
				word = strings.Title(word)
			}
			res += word
		}
	} else {
		for _, l := range a {
			i := strconv.QuoteRune(l)
			if i == strings.ToUpper(i) {
				res += "_"
				i = strings.ToLower(i)
			}
			res += i
		}
	}
	return res
}

func main() {
	fmt.Scan(&a)
	fmt.Print(gayaNulis(res))
}
