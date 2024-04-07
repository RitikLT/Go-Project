package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "Potato"

	fmt.Println("There is the list of fruit: ", len(fruitList))
	// fmt.Println("There is the list of fruit: ", strings.Split(fruitList))

	// for _, fruit := range fruitList {
	// 	fmt.Println(fruit)
	// }

	for i, fruit := range fruitList {
		if i != len(fruitList)-1 {
			fmt.Println(fruit, ", ")
		} else {
			fmt.Println(fruit)
		}
	}

	var vegList = [3]string{"potato", "tomato", "mashroom"}
	fmt.Println("Here is the list of vegetable", vegList)

	var listNumber = [4]int{1, 2, 3, 4}
	fmt.Println(len(listNumber))
	fmt.Println("Here is the list of number:", listNumber)
	for j, number := range listNumber {
		if j != len(listNumber)-1 {
			fmt.Println(number, ", ")
		} else {
			fmt.Println(number)
		}
	}
}
