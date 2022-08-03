package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// World, One, Two  -> defer stack of main func
// Hello, 43210, Two, One, World -> the final output,
// 0, 1, 2, 3, 4 -> defer stack of myDefer func

//defer statements always execute at last in reverse order (LIFO)

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
