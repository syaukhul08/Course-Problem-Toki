package main

import "fmt"

func main() {
	var baris, kolom int
	fmt.Scanf("%d %d", &baris, &kolom)

	var arr [105][105]int

	for i := 0; i < baris; i++ {
		for j := 0; j < kolom; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
	for i := 0; i < kolom; i++ {
		for j := baris - 1; j >= 0; j-- {
			fmt.Print(arr[j][i], " ")
		}
		if i+1 != baris {
			fmt.Println()
		}
	}
}
