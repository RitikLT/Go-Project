package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the picture of if...else")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is equal to zero")
	} else {
		fmt.Println("number is not equal to zero")
	}

	if num := 11; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is not less than 10")
	}
}
