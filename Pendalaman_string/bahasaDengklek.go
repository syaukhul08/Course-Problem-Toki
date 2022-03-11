package main

import (
	"fmt"
	"strconv"
	"strings"
)

var kata, result string

func bahasa(word string) string {

	for _, word := range kata {
		i := strconv.QuoteRune(word)
		if i == strings.ToUpper(i) {
			result += strings.ToLower(i)
		} else {
			result += strings.ToUpper(i)
		}
	}
	return result
}

func main() {
	fmt.Scan(&kata)
	fmt.Print(bahasa(string(result)))
}
