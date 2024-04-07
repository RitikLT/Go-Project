package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointers")

	// var ptr *int
	// fmt.Println("Value of pointer: ", ptr)

	mypointer := 23
	var ptr = &mypointer
	fmt.Println("Value of the actual parameter", ptr)
	fmt.Println("Value of the actual parameter", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The value of new pointer is:", *ptr)

}
