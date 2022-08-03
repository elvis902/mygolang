package main

import "fmt"

func main() {
	fmt.Println("Wecome to my Maps in Golang")

	languages := make(map[string]string)
	languages["CPP"] = "C plus plus"
	languages["JS"] = "JavaScript"
	languages["KT"] = "Kotlin"

	fmt.Println("The languages are ", languages)
	fmt.Println("The shorts CPP means ", languages["CPP"])

	delete(languages, "KT")
	fmt.Println("The languages are ", languages)

	//looping through the map, it is extraa here, will revisited later

	for key, value := range languages {
		fmt.Printf("The key is %v and the value is %v\n", key, value)
	}

}
