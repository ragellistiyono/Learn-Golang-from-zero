package main

import "fmt"

type diTolak func(string) bool

func namaDiTolak(nama string, tolak diTolak) {
	if tolak(nama) {
		fmt.Println("Selamat Datang", nama)
	} else {
		fmt.Println("Tidak punya Akses", nama)
	}
}

func main() {
	tolak := func(nama string) bool {
		return nama == "Ragel"
	}

	namaDiTolak("Ragel", tolak)
	namaDiTolak("septian", tolak)
}
