package main

import "fmt"

type Customer3 struct {
	Name, Address string
	Age           int
}

func (customer Customer3) helooo(name string) {
	fmt.Println("Hello ",name , "my name is ", customer.Name)
}

func main() {
	budi := Customer3 {
		Name: "Budi",
		Address: "Cinere",
		Age: 25,
	}
	budi.helooo("Joko")
}