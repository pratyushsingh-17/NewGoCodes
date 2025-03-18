package main

import "fmt"

func main() {
	patty := User{"Patty", "singhpratyush56312gmail.com", true, 12}
	patty.GetStatus()
	patty.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	if u.Status {
		fmt.Println("Online")
	}
}

func (u User) NewMail(){
	u.Email = "singh@gmail.com"
	fmt.Println("Email of user is : ", u.Email)
}