package main

import "fmt"

type Address struct {
	City, State, Country string
}

func main() {	
	/*
		Secara default golang menggunakan konsep pass by value
		Jadi datanya telah di assign ke variable lain, maka data tersebut
		sebenarnya di duplicate/copy. Seperti di bawah ini:
	*/
	address1 := Address{"Metro", "Lampung", "Indonesia"}
	address2 := address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
	
	/*
	Pointer digukanan untuk mengakali hal tersebut dan membuatnya
	menjadi pass by reference. Seperti di bawah ini:
	*/
	address3 := Address{"Fullham", "London", "England"}
	address4 := &address3
	address4.City = "Jakarta"
	fmt.Println(address3)
	fmt.Println(address4)

	/*
		Jika merubah keseluruhan maka variable reference tidak berubah.
	*/ 
	address5 := Address{"Fullham", "London", "England"}
	address6 := &address5
	
	address6 = &Address{"Kirby", "Liverpool", "England"}

	fmt.Println(address5)
	fmt.Println(address6)

	/*
		Untuk merubah semu maka perlu *
	*/ 
	address7 := Address{"Fullham", "London", "England"}
	address8 := &address7
	
	*address8 = Address{"Kirby", "Liverpool", "England"} // menambahkan * dan tidak menggunakan & (tidak untuk merubah keseluruhan value dalam struct)

	fmt.Println(address7)
	fmt.Println(address8)

	/*
		Untuk membuat data baru kosong
	*/ 
	alamat1 := new(Address)
	alamat2 := alamat1
	
	alamat2.Country = "Georgia"

	fmt.Println(alamat1) // alamat 1 berubah
	fmt.Println(alamat2)
}