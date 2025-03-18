package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")
	var fruits [4]string
	fruits[0]="Apple"
	fruits[1]="Tomato"
	fruits[3]="Peach"

	fmt.Println("Fruitlist is ", fruits)
	fmt.Println("length ", len(fruits))

	var vegetables = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Length", len(vegetables))
	fmt.Println("List", vegetables)
	fmt.Printf("Type of veg is %T\n ", vegetables)
}
