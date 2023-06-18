package main

import "fmt"

func name()(firstname, midlename, lastname string){
	firstname = "Afrian"
	midlename = "Eko"
	lastname = "Prasetyo"

	return
}

func main() {
	a, b, c := name()
	fmt.Println(a, b, c)
}