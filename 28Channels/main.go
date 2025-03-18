package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 1)

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myCh)/
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		// myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
