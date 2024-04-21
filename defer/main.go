package main

import "fmt"

func main() {
	defer fmt.Println("Welcome to the new class!!")
	defer fmt.Println("one")
	defer fmt.Println("Two")
	fmt.Println("Hello world!!")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
