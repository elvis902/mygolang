package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class on Pointers in Golang")

	// var ptr *int
	// fmt.Println("The value of a default pointer : ", ptr)

	myNumber := 10
	var ptr = &myNumber
	fmt.Println("The value of ptr is ", ptr)
	fmt.Println("The value of inside the ptr is ", *ptr)

	*ptr = *ptr + 1
	fmt.Println("The new value of myNumber is ", myNumber)
}
