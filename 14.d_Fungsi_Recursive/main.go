package main

import "fmt"

func fungsiRecursive(nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return nilai * fungsiRecursive(nilai-1)
	}
}

func main() {
	panggilRecursive := fungsiRecursive(5)

	fmt.Println(panggilRecursive)
}
