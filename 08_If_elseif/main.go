package main

import "fmt"

func main() {

	nilai := 0 //nilai variabel tampung lngsung di inisiai 0

	fmt.Printf("Masukkan nilai kamu :  ")
	fmt.Scanln(&nilai) // kalau menggunakan scanln, tanpa harus menulis format tipe data

	//if else dengan banyak kondisi
	if nilai >= 10 {
		fmt.Println("Nilai kamu Sempurna!, Kamu lulus!")
	} else if nilai >= 8 {
		fmt.Println("Kamu Lulus")
	} else if nilai == 7 {
		fmt.Println("Harus belajar lagi")
	} else {
		fmt.Println("Kamu tidak Lulus!")
	}
}
