package main

import "fmt"

func main() {
	// constant tidak bisa dirubah nilai nya, tidak di deklarasikan pun juga tidak dipermasalahkan golang
	const (
		nama    = "Ragel Listiyono"
		ganteng = true
		umur    = 25
	)

	//penulisan gaya horizontal
	const nama1, ganteng1, umur1 = "Ragel", true, 25

	fmt.Println(nama)
	fmt.Println(ganteng)
	fmt.Println(umur)
	fmt.Println()
	fmt.Println(nama1)
	fmt.Println(ganteng1)
	//fmt.Println(umur1)

}
