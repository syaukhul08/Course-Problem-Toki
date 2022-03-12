package main

import "fmt"

func palindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	result := palindrome(s)
	if result == true {
		fmt.Println("YA")
	} else {
		fmt.Println("BUKAN")
	}
}
