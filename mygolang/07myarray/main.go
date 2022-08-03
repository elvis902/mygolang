package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of arrays in golang")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "orange"

	fmt.Println("The list of Fruits is ", fruitList)
	fmt.Println("The length of the list of Fruits is ", len(fruitList))

	//We can directly initialize in the following way
	var vegList = [3]string{"potato", "tomato", "cucumber"}
	fmt.Println("The vegy list is ", vegList)
	fmt.Println("The length of the vegy list is ", len(vegList))
}
