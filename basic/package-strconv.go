package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String to apa yg kita mau
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(2000000, 16)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}