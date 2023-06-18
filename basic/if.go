package main

import "fmt"

func main()  {
	name := "Noah"

	if name == "Adam" {
		fmt.Println("Hi")
	} else if name == "Jacob" {
		fmt.Println("Hu")
	} else {		
		fmt.Println("Ho")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("short version")
	} else {		
			fmt.Println("version")
	}
}