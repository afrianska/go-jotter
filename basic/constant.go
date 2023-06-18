package main

import "fmt"

func main() {
	const firstName string = "Afrian"
	const lastName = "EP"
	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		aName string = "ssss"
		bName 		 = "rrrr"
	)

	fmt.Println(aName)
	fmt.Println(bName)
}