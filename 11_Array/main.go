package main

import "fmt"

func main() {

	//penulisan array dengan format penulisan variabel
	var nama [3]string

	nama[0] = "Lionel"
	nama[1] = "Andres"
	nama[2] = "Messi"

	fmt.Println(nama[0], nama[1], nama[2])
	fmt.Println("========================")

	//penulsian array horizontal dengan type inference dan langsung inisiasi nilai
	nama1 := [4]string{"Lionel", "Andres", "Messi", "Ragel"}

	fmt.Println("Nama : ", nama1)
	fmt.Println("========================")

	//penulisan array vertikal dengan type inference dan langsung inisiasi nilai
	nama2 := [3]string{
		"Listiyono",
		"Ragel",
		"Messi",
	}
	fmt.Println(nama2)
	fmt.Println("========================")

	//Penulisan array tanpa jumlah elemen dengan [...]
	nama3 := [...]string{
		"Ragel",
		"Listiyono",
	}
	fmt.Println(nama3)
	fmt.Println("========================")

	//array multi dimensi dengan penulisan variabel menggunakan type inference
	nama4 := [3][2]int{{1, 2}, {3, 6}, {7, 9}}

	fmt.Println(nama4)
	fmt.Println("========================")

	//Perulangan array menggunakan for
	nama5 := [...]string{"Lionel", "Andres", "Messi"}

	for i := 0; i < len(nama5); i++ {
		fmt.Println("nama : ", i, nama5[i])
	}
	fmt.Println("========================")

	//Perulangan aray dengan for-range | Lebih simpel menurut saya

	nama6 := [4]string{"Thiago", "Messi", "Chiro", "Ragel"}

	for _, namanya := range nama6 { //kalau pengen membuang nilai index, gunaakan underscore
		fmt.Println(namanya)
	}
	fmt.Println("========================")

	//membuat array dengan keyword 'make'
	nama7 := make([]string, 2)

	nama7[0] = "Ragel"
	nama7[1] = "Listiyono"

	fmt.Println(nama7)
}
