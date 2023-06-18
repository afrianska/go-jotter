package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	fmt.Println("Hello", filter(name))
// }

// Use type declarations
type Filter func(string) string 

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string  {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Afrian", spamFilter)

	// Or
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}