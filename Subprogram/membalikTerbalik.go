package main

import "fmt"

func reverse(x int) int {
	res := 0

	temp := x

	for temp > 0 {
		res = (res * 10) + (temp % 10)
		temp = temp / 10
	}
	return res
}

func main() {
	var a, b, c int
	fmt.Scanf("%d %d", &a, &b)

	c = reverse(a) + reverse(b)
	c = reverse(c)
	fmt.Println(c)

}
