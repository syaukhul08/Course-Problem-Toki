package main

import "fmt"

func main() {
	result := 0
	for {
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			break
		}
		result += number
	}
	fmt.Println(result)
}
