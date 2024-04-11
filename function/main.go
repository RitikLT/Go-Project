package main

import "fmt"

func main() {
	fmt.Println("Welcome to the chapter of function")
	// greeter()
	// greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Results are:", proResult)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Welcome to the Golang")
}
