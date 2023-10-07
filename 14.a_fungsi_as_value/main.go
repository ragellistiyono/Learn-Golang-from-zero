package main

import "fmt"

func imTheGoat(nama string) string {
	return "Lionel Messi " + nama
}

func main() {
	result := imTheGoat("goat")

	fmt.Println(result)
}
