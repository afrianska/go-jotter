package main

import "fmt"

// interface kosong
func random() interface{} {
	return false
}

func main() {
	/*
	Di bawah ini penggungaan type assertion
	Tapi jika salah dalah menentukan type datanya
	maka golang akan panik dan menghentikan app
	*/
	// var result interface{} = random()
	// var resultString string = result.(string) //type assertion dari interface kosong
	// fmt.Println((resultString))

	/*
	Untuk mengakali hal tersebut bisa menggunakan\
	switch seperti di bawah
	*/

	resultWithSwitch := random()
	switch value := resultWithSwitch.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("unknown")
	}
}