package main

import "fmt"

func names() (string, string, int) {
	return "Hihang", "Hoheng", 26
}

func main() {
	firstName, _, age := names()
	fmt.Println(firstName)
	fmt.Println(age)
}