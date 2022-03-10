package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var arr [100001]int
	var num [1001]int

	var modus int

	for i := 0; i < n+1; i++ {
		fmt.Scanf("%d", &arr[i])
		num[arr[i]]++
		if num[arr[i]] >= num[modus] {
			if num[arr[i]] > num[modus] {
				modus = arr[i]
			} else {
				if arr[i] > modus {
					modus = arr[i]
				}
			}
		}
	}
	fmt.Print(modus)
}
