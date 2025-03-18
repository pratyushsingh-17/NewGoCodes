package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	languages := make(map[string]string)

	languages["Js"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of lang : " ,languages)
	fmt.Println("JS means : " ,languages["Js"])
	
	delete(languages, "RB")
	fmt.Println("List of lang : " ,languages)

	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
}