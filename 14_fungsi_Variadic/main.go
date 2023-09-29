package main

import "fmt"

func theGoat(nmr ...int) int {

	total := 0

	for _, cok := range nmr {
		fmt.Println(cok)
		total += cok
	}

	return total

}

func main() {
	totals := theGoat(10, 20, 30, 40)
	fmt.Println("Total : ", totals)

}
