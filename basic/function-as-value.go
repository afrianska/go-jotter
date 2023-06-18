package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye lady " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Neza"))
}