package main

import "fmt"

type Man struct {
	Name string
}

// Method tanpa pointer
func (man Man) MarriedNoPointer() {
	man.Name = "Mr. " + man.Name
}

// Menggunakan pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	noPointer := Man{"Eko"}
	withPointer := Man{"Afrian"}
	
	noPointer.MarriedNoPointer()
	withPointer.Married()

	fmt.Println(noPointer.Name)
	fmt.Println(withPointer.Name)
}