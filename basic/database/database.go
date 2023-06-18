package database

import "fmt"

var connection string

func init() {
	fmt.Println("check ini dipanggil")
	connection = "MongoDB"
}

func GetDB() string {
	return connection
}