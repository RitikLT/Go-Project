package main

import "fmt"

func main() {
	fmt.Println("Welcome to chapter of structure")

	Ritik := User{"Ritik", "Ritik@go.dev", true, 23}
	fmt.Println(Ritik)
	fmt.Printf("Ritik details are: %+v\n", Ritik)
	fmt.Printf("Name is %v and email is %v.", Ritik.Name, Ritik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
