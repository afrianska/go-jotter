package other

import "fmt"

var app = "1.0.0" // tidak bisa diakses dari luar package other | diawali huruf kecil
var App = "Golang" // bisa diakses karena diawali dengan huruf besar/kapital

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// fungsi ini tidak bisa diakses dari luar packagae other
func sayGoodbye(name string)  {
	fmt.Println("Bye ", name)
}