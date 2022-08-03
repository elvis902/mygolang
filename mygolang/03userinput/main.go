package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome, Enter the rating here :")
	reader := bufio.NewReader(os.Stdin)

	//comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("You have given a rating of : ", input)
	fmt.Printf("The type of the rating is : %T", input)
}
