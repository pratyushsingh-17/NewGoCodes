package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	fmt.Println(days)

	// for i:=0; i<len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("Index is %v value is %v\n", index, day)
	// }

	rougeValue := 1

	// for rougeValue <10{
	// 	fmt.Println("Value is ", rougeValue)
	// 	rougeValue++
	// }

	for rougeValue <10{
		if rougeValue == 2{
			goto lco
		}
		if rougeValue == 5{
			rougeValue++
			continue
		}else if rougeValue ==8{
			break
		}
		fmt.Println("Value is ", rougeValue)
		rougeValue++
	}
	lco:
		fmt.Println("Jumping at google.com")

}
