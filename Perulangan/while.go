package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var words[] string

	for range words {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()

		fmt.Scan(words)
	}

	for range words {
		fmt.Println(words)
	}

}
