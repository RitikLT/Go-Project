package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Welcome to the chapter of slices")
	var fruitList = []string{"Apple", "Peach", "Tomato"}
	fmt.Printf("This is the type %T\n", fruitList)

	//We can add as much value in this type slices
	fruitList = append(fruitList, "Mango", "Grapes", "Banana")
	fmt.Println("The final list of fruit is:", fruitList)
	fmt.Println(strings.Join(fruitList, ", "))
	fmt.Println(fruitList[1:5])

	highScores := make([]int, 4)
	highScores[0] = 121
	highScores[1] = 221
	highScores[2] = 543
	highScores[3] = 110
	// highScores[4] = 100
	highScores = append(highScores, 444, 553)
	fmt.Println("Here is the scores list:", highScores)
	sort.Ints(highScores)

	var course = []string{"javascript", "python", "ruby", "HTML", "CSS"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
