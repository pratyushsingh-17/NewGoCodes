package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Lco web req")

	response, err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Response type : %T\n", response)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(data))
}