package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNo := rand.Intn(6)
	fmt.Println("Value of dice is ", diceNo)

	switch diceNo {
	case 1:
		fmt.Println("Dice val 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll again")
	default:
		fmt.Println("???????")
	}


}
