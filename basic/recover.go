package main

import "fmt"

func endAppRecover() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runAppRecover(error bool)  {
	defer endAppRecover()
	if error {
		panic("ERROR")
	} else {
		fmt.Println("Aplikasi berjalan")
	}
}

func main() {
	runAppRecover(true)
	fmt.Println("Adam")
}