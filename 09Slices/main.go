package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to videos on slices")

	var fruits =[]string{"Apple", "Tomato", "Peach"}
	fmt.Printf("type of fruitlist is %T\n", fruits )
	fruits= append(fruits, "mango", "Banana")
	fmt.Println(fruits)
	fmt.Println(fruits[1:])
	fmt.Println(fruits[1:3])
	fmt.Println(fruits[:3])

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 456
	highScore[2] = 678
	highScore[3] = 789
	// highScore[4] = 890
	highScore = append(highScore, 123, 1234, 2345, 4567)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove val from slice based on index
	var courses = []string{"Reactjs", "Javascript", "swift", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
