package main

import "fmt"

func main()  {
	person := map[string]string {
		"name" : "Adam",
		"city" : "Metro",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["city"])
	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Hahah"
	book["author"] = "jajaj"
	book["genre"] = "lalala"
	book["release"] = "mama"
	fmt.Println(len(book))

	delete(book, "genre")

	fmt.Println(len(book))
	fmt.Println(book["title"])
}