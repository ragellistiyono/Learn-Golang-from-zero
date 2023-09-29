package main

import "fmt"

func main() {
	//penulisan variabel langsung di inisiasikan tipe datanya
	var gel1 = "Ragel Listiyono "

	//penulisan variabel dengan tipe data dan di inisiasi nilainya
	var gel2 string = "Ragel Listiyono"

	//penulisan variabel dengan type inference atau penulisan variabel tanpa tipe data dan kata var
	gel3 := "Ralisto Reddington"

	//penulisan multi variabel
	gel4, gel5, gel6 := "Ragel", 25, "Badminton"

	//penulisan vaariael underscore _ berguna untuk menampung nilai yang tdk diperlukan

	gel7, _, gel9 := "Ralisto", 69, "Mancing"

	fmt.Printf("Nama Saya :%s\n", gel1) //penulisan fmt.printf menggunakan format tipe datanya
	fmt.Println("Nama saya : ", gel1)   //penulisan fmt.println tanpa menuliskan format tipedata
	fmt.Println(gel2)
	fmt.Println(gel3)
	fmt.Printf("Nama %s, Umur %d, Hoby %s\n", gel4, gel5, gel6)
	fmt.Printf("Nama %s, Hoby %s\n", gel7, gel9)
}
