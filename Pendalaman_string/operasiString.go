package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2, s3, s4 string
	fmt.Scan(&s1, &s2, &s3, &s4)
	result1 := strings.Split(s1, s2)[0]
	result2 := strings.Split(s1, s2)[1]
	result := result1 + result2
	result1 = strings.Split(result, s3)[0]
	result2 = strings.Split(result, s3)[1]
	result = result1 + s3 + s4 + result2
	fmt.Print(result)
}
