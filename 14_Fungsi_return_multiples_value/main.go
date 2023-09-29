package main

import "fmt"

func theGoat() (string, string, int) { // penulisan multiple return value
	return "Lionel", "Messi", 10
}

func main() {
	namaDepan, namaBelakang, nomorPunggung := theGoat()
	fmt.Println(namaDepan, namaBelakang, nomorPunggung)
}
