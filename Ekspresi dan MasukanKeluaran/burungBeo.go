package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	line := scan.Text()
	fmt.Println(line)
}
