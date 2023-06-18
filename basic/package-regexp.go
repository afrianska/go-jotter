package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("ezo"))
	fmt.Println(regex.MatchString(" "))

	fmt.Println(regex.FindAllString("Eko edo egi ego elo eto", 10))
}