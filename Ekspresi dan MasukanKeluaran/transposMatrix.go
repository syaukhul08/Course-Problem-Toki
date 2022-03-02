package main

import (
	"fmt"
)

func main() {

	var a, b, c, d, e, f, g, h, i int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Scanf("%d %d %d", &d, &e, &f)
	fmt.Scanf("%d %d %d", &g, &h, &i)

	fmt.Printf("%d %d %d\n", a, d, g)
	fmt.Printf("%d %d %d\n", b, e, h)
	fmt.Printf("%d %d %d\n", c, f, i)
}
