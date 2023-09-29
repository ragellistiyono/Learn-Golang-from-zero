package main

import "fmt"

func main() {
	nama1 := "Ragel"
	nama2 := "Listiyono"

	hasilNama := nama1 == nama2
	fmt.Println(hasilNama)

	fmt.Println("================")

	angka1, angka2 := 10, 20

	fmt.Println(angka1 > angka2)
	fmt.Println(angka1 < angka2)
	fmt.Println(angka1 == angka2)
	fmt.Println(angka1 >= angka2)
	fmt.Println(angka1 != angka2)
	fmt.Println("===Operator boolean===")

	matik, ipa := 80, 90

	lulusMatik, lulusIpa := matik > 70, ipa > 70

	lulus := lulusMatik && lulusIpa
	fmt.Println(lulus)

	//penulisan lainnya
	fmt.Println(matik >= 70 && ipa >= 70)

}
