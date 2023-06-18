package main

import "fmt"

func main() {
	counter := 0
	name := "haha"

	increment := func() {
		fmt.Println("Increment")
		counter++
		name := "hoho"
		fmt.Println(name)
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}