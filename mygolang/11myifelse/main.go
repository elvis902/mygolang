package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in Golang")

	var loginCount int
	_, err := fmt.Scanf("%d", &loginCount)
	if err != nil {
		fmt.Println("There was an error while reading input", err)
		return
	}

	if loginCount < 10 {
		fmt.Println("Regular User")
	} else if loginCount > 10 {
		fmt.Println("Watch out")
	} else {
		fmt.Println("Exactly 10 login count")
	}

	if 9%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// initializing and checking the condition at the same time, we can see like this
	//statement in many places while working with Golang
	if num := 4; num < 10 {
		fmt.Println("The num is less than 10")
	} else {
		fmt.Println("The num is NOT less than 10")
	}
}
