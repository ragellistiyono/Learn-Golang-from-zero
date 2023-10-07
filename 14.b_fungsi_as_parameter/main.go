package main

import (
	"fmt"
	"strings"
)

type diFilter func(string) string

func namaAnak(nama string, filter diFilter) {
	namaAnakIbu := filter(nama)
	fmt.Println(namaAnakIbu)
}

func xNama(nama string) string {
	if strings.Contains("fintyas anggraini", strings.ToLower(nama)) {
		return nama + " Adalah Anak Pertama"
	} else if strings.Contains(strings.ToLower(nama), "anita") {
		return nama + " Adalah Anak kedua"
	} else if strings.Contains(strings.ToLower(nama), "septian") {
		return nama + " Adalah Anak ketiga"
	} else if strings.Contains(strings.ToLower(nama), "ragel") {
		return nama + " Adalah Anak terakhir"
	} else {
		return "Silahkan pilih nama sesuai daftar pilihan"
	}
}

func main() {
	var nama string
	cok := `
			======================================
			Daftar nama anak Ibu :

			fintyas anggraini
			anita
			septian
			ragel
			====================================== `

	fmt.Println(cok)
	fmt.Println()
	fmt.Print("Masukkan pilihan sesuai nama : ")
	fmt.Scanln(&nama)

	namaAnak(nama, xNama)
}
