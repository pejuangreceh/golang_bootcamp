package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	for i := 1; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "three"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recevied", msg1)
		case msg2 := <-c2:
			fmt.Println("recevied", msg2)
		case msg3 := <-c3:
			fmt.Println("recevied", msg3)
		}
	}
}
