package main

import "fmt"

func main() {
	defer fmt.Println("defer")
	defer fmt.Println("defer2")
	fmt.Println("Hello")
	myDefer()
}

func myDefer()  {
	for i := 0; i <5; i++ {
		defer fmt.Println(i)
	}
}
