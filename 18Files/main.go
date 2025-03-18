package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")

	content := "This needs to go in a file - learnCodeOnline.in"

	file, err := os.Create("./nyDir.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length : ", length)
	defer file.Close()
	readfile("./nyDir.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error){
	if err!=nil {
		panic(err)
	}
}