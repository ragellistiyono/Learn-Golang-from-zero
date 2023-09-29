package main

import "fmt"

func namaGoat(namaLengkap string) string {
	return "I'am :" + namaLengkap
}

func main() {
	//Penulisan dengan ditampung menggunakan variabel
	hasil := namaGoat("Lionel Messi")
	fmt.Println(hasil)
	fmt.Println("=====================")

	//penulisan langsung
	println(namaGoat(" Lionel Andres Messi"))
}
