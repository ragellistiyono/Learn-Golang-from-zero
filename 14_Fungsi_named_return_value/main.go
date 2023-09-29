package main

import "fmt"

//Penulisan named return multiple value dengan inisiasi nama variabel langsung di block kode
func theGoat() (namaDepan, namaBelakang string, nomorPunggung int) {
	namaDepan = "Lionel"
	namaBelakang = "Messi"
	nomorPunggung = 10

	return
}

func main() {

	//pemanggilang langsung
	fmt.Println(theGoat())
	fmt.Println("=================")

	//pemanggilaan dengan nama variabel metode type inference dan penampungan variabel _
	//penamaan variabel tidak harus sama dengan variabel dlm funcion multiple value
	_, namaBelakang, nomorPunggung := theGoat()
	fmt.Println(namaBelakang)
	fmt.Println(nomorPunggung)
	fmt.Println("=================")

	//penggunaan for pada named reurn multiple value
	for i := 0; i < 5; i++ {
		namaDepan, namaBelakang, nomorPunggung := theGoat()
		fmt.Println(namaDepan, namaBelakang, nomorPunggung)
	}
	fmt.Println("=================")

}
