package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("\033[1;36mLCO web Request\033[0m")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\033[1;35mResponse is of type: %T\033[0m\n", response)
	fmt.Println("\033[1;32mThe response are the folowing:\033[0m", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
	fmt.Println("This is done")

	fmt.Printf("\033[1;35mResponse is of type: %T\033[0m\n", response)
	fmt.Println("\033[1;32mThe response are the folowing:\033[0m", response)
}
