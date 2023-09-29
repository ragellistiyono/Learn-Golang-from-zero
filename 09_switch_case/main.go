package main

import "fmt"

func main() {

	nilai := 0

	fmt.Printf("Masukkan Nilai : ")
	fmt.Scanln(&nilai)

	/*switch case bisa menggunakan gaya if else, dengan penulisan variabel seteleah switch di hilangkan
	dan ditulis pada bagian case contoh :
		switch {
		case (nilai < 3 ) && ( nilai <= 4):
			fmt.Println("Tidak Lulus")
		}
	*/
	switch nilai {
	case 10:
		fmt.Println("Nilai Sempurna")
	case 9:
		fmt.Println("Nilai Hampir Sempurna")
	case 8:
		fmt.Println("Nilai mendekati Sempurna")
	case 7, 6, 5: //case banyak kondisi
		fmt.Println("Harus Belajar lagi")
	default:
		{ // kurung kurawal {} berguna untuk membuat banyak statement
			fmt.Println("Hanya masukkan Angka")
			fmt.Println("Coba jalankan lagi program ini")
		}
	}

}
