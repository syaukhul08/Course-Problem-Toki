package main

import "fmt"

func main() {
	var words []string
	for {
		var word string
		_, err := fmt.Scanf("%s", &word)
		if err != nil {
			break
		}
		words = append(words, word)
	}
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
}
