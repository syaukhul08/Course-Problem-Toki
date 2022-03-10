package main

import "fmt"

func main() {

	var arr [105][105]int
	var arr1 [105][105]int
	var arr2 [105][105]int
	var x, y, z int

	fmt.Scanf("%d %d %d", &x, &y, &z)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Scanf("%d", &arr1[i][j])
		}
	}
	for i := 0; i < y; i++ {
		for j := 0; j < z; j++ {
			fmt.Scanf("%d", &arr2[i][j])
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			arr[i][j] = 0
		}
	}

	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			for k := 0; k < y; k++ {
				arr[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			if j == 0 {
				fmt.Printf("%d", arr[i][j])
			} else {
				fmt.Printf(" %d", arr[i][j])
			}
		}
		fmt.Println("\n")
	}
}
