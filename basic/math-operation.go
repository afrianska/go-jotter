package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	// augmented assignment
	var i = 12
	i *= 10 // i = i + 10
	fmt.Println(i)

	// Unay operator
		var k = 12
	k-- // k = k + 1
	fmt.Println(k)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}