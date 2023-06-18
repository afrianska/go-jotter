package main

import "fmt"

func main() {
	var name string
	name = "Afrian"
	fmt.Println(name)

	// variable direct declaration
	var city = "Japan"
	fmt.Println(city)

	// variable without var
	address := "Hahaha"
	fmt.Println(address)

	// multiple variable
	var (
		firstName = "Afrian"
		lastName = "Prasetyo"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}