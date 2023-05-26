package main

import "fmt"

func printOddNumber(ch chan int) {
	for i := 1; i < 20; i = i + 2 {
		println("  ", i)
	}
	ch <- 1
}

func printEvenNumber(ch chan int) {
	for i := 2; i < 20; i = i + 2 {
		println(i)
	}
	ch <- 2

}
func printNumber(ch chan int, id int) {
	for i := 1; i < 11; i = i + 2 {
		println("  ", i)
	}
	ch <- id
}

func main() {
	//inisialisiasi channel
	//ch := make(chan int)
	//menerima data dari channel
	//data := <-ch
	//mengirim data ke channel
	//ch <- data
	ch := make(chan int)
	go printOddNumber(ch)
	go printEvenNumber(ch)
	fmt.Println("Finished oddeven", <-ch)
	fmt.Println("Finished oddeven", <-ch)

	for i := 0; i < 11; i++ {
		go printNumber(ch, i)
	}
	for i := 0; i < 11; i++ {
		fmt.Println("Finished ", <-ch)
	}

	//message := make(chan string)
	//go func() {
	//	message <- "Hello World"
	//}()
	//
	//receivedMsg := <-message
	//fmt.Println(receivedMsg)
}
