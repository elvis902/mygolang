package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Wecome to the study on pointers")

	//If we donot provide the size then we are declaring a slice in golang
	//slice is heavily used in golang
	var fruitList = []string{"Apple", "Banana", "Peach"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Watermelon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 876
	highScores[1] = 675
	highScores[2] = 113
	highScores[3] = 965
	//highScores[4] = 321

	highScores = append(highScores, 321, 554)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value form slice based on index
	var courses = []string{"react.js", "python", "kotlin", "java", "golang", "c++"}
	var index int = 2
	fmt.Println(courses)
	courses = append(courses[:index], courses[index+1:]...)
	//After removing index from slice courses
	fmt.Println(courses)
}
