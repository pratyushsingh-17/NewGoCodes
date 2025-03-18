package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	// mut.RLock()
	var score = []int{0}
	// mut.RUnlock()

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)


	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	
	
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)


	go func (wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("Read")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
	fmt.Println(score)
}
