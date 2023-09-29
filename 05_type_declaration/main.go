package main

import "fmt"

func main() {

	type NamaLengkap string
	type umur int
	type tampan bool

	var jeneng NamaLengkap = "Ragel Listiyono"
	var umurCok umur = 25
	var paras tampan = true

	fmt.Println("Nama \t: ", jeneng)
	fmt.Println("Umur lur :", umurCok)
	fmt.Println("Ganteng ta ?", paras)
}
