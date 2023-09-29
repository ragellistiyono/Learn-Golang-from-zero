package main

import "fmt"

func main() {

	//pembuatan slice sama halnya dengan array, bedanya hanya jumlah elemen tidak di isi
	nama := []string{"Ragel", "Listiyono", "Messi", "Thiago"}
	fmt.Println(nama[0], nama[2]) // kalau pemanggilan variabel seperti ini tanda [] tidak iku ter print
	fmt.Println(nama)             // tanda [] akan ter print

	fmt.Println("===================")

	newNama := nama[0:3] // dimulai dari index ke-0, mengambil 3 elemen dari index ke-0
	fmt.Println(newNama)
	fmt.Println("===================")

	xNama := nama[0:3] // dimulai dari index ke-0, mengambil 3 elemen dari index ke-0
	yNama := nama[1:4] // dimulai dari index ke-1 dan elemen yang di ambil 3 elemen dimulai dari index ke-1

	xxNama := xNama[1:2] //variabel baru merubah nilai dr variabel lama
	yyNama := yNama[0:1] //variabel baru merubah nilai dr variabel lama

	fmt.Println(nama)
	fmt.Println(xNama)
	fmt.Println(yNama)
	fmt.Println(xxNama)
	fmt.Println(yyNama)

	//dirubah nilai elemennya
	yyNama[0] = "Ronaldo"

	fmt.Println(nama)
	fmt.Println(xNama)
	fmt.Println(yNama)
	fmt.Println(xxNama)
	fmt.Println(yyNama)
	fmt.Println("===================")
	//len dan cap
	fmt.Println(len(xNama))
	fmt.Println(cap(xNama))
	fmt.Println("===================")
	//append digunakan untuk menambahkan elemen di akhir
	cNama := append(nama, "Wiranto")
	fmt.Println(nama)
	fmt.Println(cNama)

}
