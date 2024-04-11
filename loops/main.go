package main

import "fmt"

func main() {
	fmt.Println("Let's handsUp with loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	roogueValue := 1
	for roogueValue < 10 {
		if roogueValue == 5 {
			roogueValue++
			continue
		} else if roogueValue == 8 {
			break
		}
		fmt.Println("value is:", roogueValue)
		roogueValue++
	}
}
