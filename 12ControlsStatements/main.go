package main

import "fmt"

func main() {
	fmt.Println("If else")

	loginCount := 10
	var res string

	if loginCount < 10{
		res = "regular"
	}else if loginCount >10{
		res = "Watch out"
	}else{
		res = "Exactly 10"
	}
	fmt.Println(res)

	if loginCount%2==0 {
		fmt.Println("Even")
	}else{
		fmt.Println("odd")
	}

	if num:=3; num<10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is more than 10")

	}


}
