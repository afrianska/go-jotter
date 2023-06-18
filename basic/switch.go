package main

import "fmt"

func main()  {
	name := "Noahaaaa"

	switch name {
	case "Noah":
		fmt.Println("Hi Noah")
	case "Jonah":
		fmt.Println("Hi Jonah")
	default:
		fmt.Println("Hi Adam")
	}

	// short statemen
	switch  length := len(name); length > 5 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	// tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("karakter lebeih dari 10")
	case length > 5:
		fmt.Println("karakter lebeih dari 5")
	default:
		fmt.Println("nama kurang dari 5")
	}
}