package main

import "fmt"

func main() {
	fmt.Println("Welcome the chapter of maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Here is the list of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	// delete(languages, "RB")
	// fmt.Println("Here is the list of all languages:", languages)

	//Loops are interesting in golangs
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	fmt.Println("Welcome to the chapter two of maps")

	langpreference := make(map[int]string)

	langpreference[1] = "Python"
	langpreference[2] = "Cascandra"
	langpreference[3] = "C++"
	langpreference[4] = "Angular"

	fmt.Println("Here is the lanaguages preference:", langpreference)
	fmt.Println("The first languages prefernece is:", langpreference[1])

	for _, value := range langpreference {
		fmt.Println("the stands for:", value)
	}

}
