package main

import "fmt"

var jwtToken = 4876628

//Public 
const LoginToken string = "1234567890"

func main() {
	var username string = "Pratyush"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
	
	var smallFloat float32 = 255.123456123456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	//implicit type
	var website = "https://google.com"
	fmt.Println(website)

	//no var style
	noOfUsers := 300000.0
	fmt.Println(noOfUsers)

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)
}