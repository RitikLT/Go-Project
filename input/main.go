package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	fmt.Println("Welcome to user input --->")
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the rating for our Pizzas")

// 	//comma ok || err err
// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanx for rating", input)
// 	fmt.Printf("Type of this rating is: %T", input)

// 	fmt.Println("Welcome to out restraurant")
// 	readerone := bufio.NewReader(os.Stdin)
// 	fmt.Println("Plz rate the restaurant:")

// 	inputekd, _ := readerone.ReadString('\n')
// 	fmt.Println("thanx for the rate ourself", inputekd)

// 	// Scanning for a string input
// 	var inputekd string
// 	fmt.Scanln(&inputekd)
// 	fmt.Println("Thanks for rating our restaurant:", inputekd)

// }

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to user input --->")
	fmt.Println("Enter the rating for our Pizzas")

	// ReadString will block until the delimiter is entered
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)

	fmt.Println("Welcome to our restaurant")
	fmt.Println("Please rate the restaurant:")

	// Scanning for a string input
	var inputekd string
	fmt.Scanln(&inputekd)
	fmt.Println("Thanks for rating our restaurant:", inputekd)
}
