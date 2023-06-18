package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "admin", "Put your database user")
	var pass *string = flag.String("pass", "admin", "Put your database pass")
	var number *int = flag.Int("number", 123, "Put your database pass")

	flag.Parse()

	fmt.Println("host : ", *host)
	fmt.Println("user : ", *user)
	fmt.Println("pass : ", *pass)
	fmt.Println("number : ", *number)
}