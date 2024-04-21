package main

import "fmt"

func main() {
	fmt.Println("Welcome to chapter of structure")

	Ritik := User{"Ritik", "Ritik@go.dev", true, 23}
	fmt.Println(Ritik)
	fmt.Printf("Ritik details are: %+v\n", Ritik)
	fmt.Printf("Name is %v and email is %v.\n", Ritik.Name, Ritik.Email)
	Ritik.GetStatus()
	Ritik.GetEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) GetEmail() {
	u.Email = "Test@Go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
