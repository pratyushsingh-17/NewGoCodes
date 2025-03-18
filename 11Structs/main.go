package main

import "fmt"

func main() {
	pratyush := User{"Pratyush", "Pratyush@nomiso.io", true, 24}
	fmt.Println(pratyush)
	fmt.Printf("Pratyush details are : %+v\n", pratyush)
	fmt.Printf("Name is %v age is %v", pratyush.Name, pratyush.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
