package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	//Logn operation
	fmt.Printf("Worker %d done\n", id)

}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		//time.Sleep(1 * time.Second / 10)
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All worker finished")
}
