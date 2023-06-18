package main

import "fmt"

type Alamat struct {
	City, State, Country string
}

/*
	penggunaan function di bawah ini tidak merubah data address
	Karena data yg dirubah adalah data duplicate
	data asli tidak berubah
*/
func ChangeAddressToIndonesia(address Alamat) {
	address.Country = "Indonesia"
}

/*
	Untuk merubahnya harus menggunakan pointer
*/
func ChangeAddressToIndonesiaPointer(address *Alamat) {
	address.Country = "Indonesia"
}

func main() {

	address := Alamat{"Metro", "Lampung", "Singapore"}
	ChangeAddressToIndonesia(address)
	fmt.Println(address) // Data tidak berubah
	
	address2 := Alamat{"Metro", "Lampung", "Singapore"}
	ChangeAddressToIndonesiaPointer(&address2)
	fmt.Println(address2) // Data berubah

}