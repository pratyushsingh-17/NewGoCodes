package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/todos?coursename=reactjs&paymentid=1234rfcv"

func main() {
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()

	fmt.Printf("Type of query are : %T\n", qparam)

	fmt.Println(qparam["coursename"])

	for _, value := range qparam{
		fmt.Println("Param : ", value)
	}

	// Parts of url

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/qwer",
		RawPath: "user=Pratyush",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
