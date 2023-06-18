package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Afrian Eko", "Afrian"))
	fmt.Println(strings.Contains("Afrian Eko", "Haha"))

	fmt.Println(strings.Split("Afrian Eko Prasetyo", " "))

	fmt.Println(strings.ToLower("Afrian Eko Prasetyo"))
	fmt.Println(strings.ToUpper("Afrian Eko Prasetyo"))
	fmt.Println(strings.Trim("   Afrian Eko Prasetyo  ", " "))
	fmt.Println(strings.ReplaceAll("haha haha haha hihi", "haha", "moo"))
}