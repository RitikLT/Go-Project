package main

import "fmt"

func main() {
	var username string = "Ritik"
	fmt.Println(username)
	fmt.Printf("The variable is type of:%T \n\n", username)

	var smallVal bool = true
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of: %T \n\n", smallVal)

	var isEnabled int64 = 122323232
	fmt.Println(isEnabled)
	fmt.Printf("The type of: %T \n", isEnabled)
}
