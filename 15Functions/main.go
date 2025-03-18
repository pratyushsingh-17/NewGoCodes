package main

import "fmt"

func main() {
	// greet()
	// fmt.Println("Functions")
	// greetTwo()

	// res := adder(3,4)
	// fmt.Println("Result " , res)

	result :=proAdder(1,2,3,4,5)
	fmt.Println(result)

}
func adder(val1 int, val2 int) int {
	return val1 + val2
}
func proAdder(val... int) int {
	total := 0
	for _, val := range val{
		total = val+total
	}
	return total
}
func greet()  {
	fmt.Println("Namaste")
}
func greetTwo()  {
	fmt.Println("another one")
}
