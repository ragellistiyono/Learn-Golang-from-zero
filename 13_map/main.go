package main

import "fmt"

func main() {

	//inisiasi nilai map dengan horizontal
	nilai := map[string]int{"Ragel": 100, "Messi": 90}

	fmt.Println(nilai)
	fmt.Println(nilai["Ragel"])
	fmt.Println("=====================")

	//vertikal
	nilai1 := map[string]string{
		"Ragel": "Ganteng",
		"Messi": "Goat",
	}
	fmt.Println(nilai1)
	fmt.Println("=====================")

	// for range pada map
	iniGoat := map[string]string{
		"Ragel Listiyono": "Goat",
		"Lionel Messi":    "Goat",
		"Adele Adkins":    "Goat",
	}

	for real, bos := range iniGoat { //kalau mau menampung variabel yg tdk perlu gunakan underscore _
		fmt.Println(real, ":", bos)
	}
	fmt.Println("=====================")
	//mengahpus item map dengan fungsi delete()
	delete(iniGoat, "Adele Adkins") //delete, lalu masukkan variabel, lalu key nya

	for real, bos := range iniGoat { //kalau mau menampung variabel yg tdk perlu gunakan underscore _
		fmt.Println(real, ":", bos)
	}
	fmt.Println("=====================")

	//kombinasi slice dan map

	makanan := []map[string]string{
		{"Nama makanan": "Rendang", "Asal": "Sumatra"},
		{"Nama Minuman": "Es Teh", "Asal": "Jatim"},
		{"Nama Pembuat": "Ragel", "Asal": "Surabaya"},
	}
	for _, cok := range makanan {
		fmt.Println(cok["Nama makanan"], cok["Nama Minuman"], cok["Nama Pembuat"], cok["Asal"])
	}
}
