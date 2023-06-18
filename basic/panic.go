package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	defer endApp()
	if error {
		panic("ERROR")
	} else {
		fmt.Print("Aplikasi berjalan")
	}

}

func main() {
	runApp(false)
}