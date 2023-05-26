package controller

import "fmt"

type Person struct {
	nama  string
	umur  int
	email string
}

func Cetak() {
	var p1 Person
	p1.nama = "Azis"
	p1.umur = 9
	p1.email = "azis@gmail.com"
	fmt.Println(p1.nama, " ", p1.umur, " ", p1.email)
}
