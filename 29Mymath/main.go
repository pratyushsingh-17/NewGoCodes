package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	// fmt.Println("Sum : ", mynumberOne+int(mynumberTwo))

	// Random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	// Random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
