package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang functions")
	greeter()

	result := adder(3, 2)
	fmt.Println("The value of our result is: ", result)

	proRes, myMessage, _ := proAdder(4, 5, 2, 6, 7)
	fmt.Println("The value of proResult is: ", proRes)
	fmt.Println("The message from proResult is: ", myMessage)

}

func greeter() {
	fmt.Println("Hello there !")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

//We can take numerous arguments in slice
//We can also return multiple items from function
func proAdder(values ...int) (int, string, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi from proAdder function", "just random"
}
