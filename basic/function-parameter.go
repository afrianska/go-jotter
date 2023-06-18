package main

import "fmt"

func sayHi(name string) {
	fmt.Println("Hello", name)
}

func main() {
	sayHi("Jacob")
	sayHi("Adam")
	sayHi("Moses")
}