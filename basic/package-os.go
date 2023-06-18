package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("argument:")
	fmt.Println(args)
	
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname : ", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}