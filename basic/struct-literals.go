package main

import "fmt"

type Customer1 struct {
	Name, Address string
	Age           int
}

func main() {
	joko := Customer1 {
		Name:    "Joko",
		Address: "Infonesia",
		Age:     50,
	}
	fmt.Println(joko)

	// rentan error jika merubah posisi
	budi := Customer1 {"Budi", "Indonesia", 52}
	fmt.Println(budi)

}