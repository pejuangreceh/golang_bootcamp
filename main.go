package main

import (
	"fmt"
	"golang_project/controller"
	"math"
)

type Logger interface {
	Log(msg string)
}

type UseCase struct {
	logger Logger
}

func (uc UseCase) UseCase() {
	//uc.logger.Log(msg string)
}

func main() {
	//map string key
	studentAges := map[string]int{
		"Yabes": 23,
		"Dea":   19,
	}
	//map integer key
	testMap := map[int]int{
		1: 23,
		4: 24,
	}
	//map kosong
	var map1 map[string]int
	map1 = make(map[string]int, 5)
	//map kosong tapi tanpa make
	map3 := map[string]int{}
	//isi ke map kosong
	map1["satu"] = 1
	map3["tiga"] = 3
	type Person struct {
		Name  string `json:"nama"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}
	var p = new(Person)
	p.Name = "Elloy"
	p.Age = 7
	p.Email = "email.com"
	var test int64 = 7
	var coba = "A"
	//array
	var angka = [4]int{1, 2, 3, 5}
	//slice
	numbers := []int{1, 2, 5, 3, 4}
	//var num = *int
	//ptr = &num
	//*ptr = 12
	fmt.Println("Halo bosqu", rune(test))
	fmt.Println("ASCII = ", []byte(coba))
	fmt.Println("RUNE = ", []rune(coba))
	fmt.Println("ANGKA ARRAY = ", angka[0])
	fmt.Println("Struct awal = ", p.Name)
	controller.Cetak()
	fmt.Println(&angka[2])
	controller.Create(3)
	fmt.Println(" Slice = ", numbers)
	//fmt.Println(" value Num = ", ptr)
	//fmt.Println(" addressS Num = ", *ptr)
	result1 := applyOperation("Increment by 1", 10, incrementBy1)
	result2 := applyOperation("nambah 2 bang", 10, incrementBy2)

	fmt.Println("hasil 1 = ", result1)
	fmt.Println("hasil 2 = ", result2)
	fmt.Println("Umur Yabes = ", studentAges["Yabes"])
	fmt.Println("Umur Setya = ", studentAges["Setya"])
	fmt.Println("Map ke 1 = ", testMap[1])
	fmt.Println("Map ke 4 = ", testMap[4])
	fmt.Println("Map kosong pertama = ", map1["satu"])
	fmt.Println("Map kosong pertama = ", map3)
	//perulangan
	for _, number := range numbers {
		fmt.Println(number)
	}
	//interface implementasi
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	//uc := UseCase{}
	//	main bottom
}

// ini namanya alias -> function kosongan buat parameter di applyOperation sebagai callback
type Increment func(int) int

func applyOperation(name string, num int, operation Increment) int {
	fmt.Println("operation = ", name)
	return operation(num)
}
func incrementBy1(val int) int {
	return val + 1
}
func incrementBy2(val int) int {
	return val + 2
}

type Shape interface {
	Area() float64
}

//func (c Circle) Area() float64 {
//	return 3.14 * c.Radius * c.Radius
//}

//contoh interface

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	ctrl := controller.New()
	ctrl2 := controller.New()

	ex := controller.Default()
	ex2 := controller.Default()
	fmt.Println(ctrl, " == ", ctrl2, " hasilnya ", ctrl == ctrl2)
	fmt.Println(ex, " == ", ex2, " hasilnya ", ex == ex2)
}
