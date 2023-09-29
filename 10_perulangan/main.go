package main

import "fmt"

func main() {
	/*
		for i := 0; i <= 3; i++ {
			fmt.Println("Angka", i)
		}
	*/

	var kalimat string
	var ulang int

	fmt.Print("Masukkan kalimat: ")
	fmt.Scanln(&kalimat)

	fmt.Print("Mau berapa kali di cetak : ")
	fmt.Scan(&ulang)

	//for dengan gaya while
	for i := 0; i < ulang; i++ {
		fmt.Println(kalimat)
	}

}
