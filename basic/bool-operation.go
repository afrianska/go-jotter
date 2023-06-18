package main

import "fmt"

func main()  {
	var nilaiAkhir = 90
	var absensi = 85

	var lulusNilaiakhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusNilaiakhir && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80 )
}