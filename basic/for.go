package main

import "fmt"

func main() {
	counter := 1
	for counter <= 5 {
		fmt.Println("Looping to", counter)
		counter++
	}

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Hitung", counter)
	}

	// for range
	names := []string{"Adam", "Noah", "Jonah"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	persons := make(map[string]string)
	persons["name"] = "adam"
	persons["city"] = "mbuh"

	for _, person := range persons {
		fmt.Println(person)
	}
}