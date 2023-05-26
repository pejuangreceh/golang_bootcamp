package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}
func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		//wait manual pake time sleep
		//time.Sleep(1 * time.Second / 10)
		go func() {
			counter.Increment()
			wg.Done()
		}()
		fmt.Println("Counter : ", counter.value)
	}
	wg.Wait()
	fmt.Println("Counter Value", counter.value)
}
