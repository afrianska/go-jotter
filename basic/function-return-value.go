package main

import "fmt"

func callName(name string) string {

	if name == "" {
		call := "Sorry, with sir?"
		return call
	} else {
		call := "Hi, " + name
		return call
	}
}

func main() {
	call := callName("Abraham")
	fmt.Println(call)

	fmt.Println(callName(""))
}