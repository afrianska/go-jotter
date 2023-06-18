package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(value int)  {
	defer logging()
	fmt.Println("run app")
	result := 10/value
	fmt.Println("rusult", result)
}

func main() {
	runApp(5)
}