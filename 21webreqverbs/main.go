package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostJson()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content length ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJson()  {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename" : "Golang",
			"price" : 0,
			"platform" : "lco.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err!= nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest()  {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstName", "Pratyush")
	data.Add("lastName", "Singh")
	data.Add("email", "Pratyush@go.in")
	data.Add("contact", "8271099873")

	response, err := http.PostForm(myurl, data)

	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()
	content , _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}