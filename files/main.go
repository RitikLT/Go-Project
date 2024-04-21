package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This need to fo in files - LearnCodeOnline.in"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./mygofile.txt")

}

func readFile(filename string) {
	datatype, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("The readed file is:", string(datatype))
}
