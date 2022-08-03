package main

import "fmt"

func main() {
	fmt.Println("All about loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("The index is %v and the value is %v\n", index, value)
	// }

	//It is like a while loop
	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 8 {
			break
		}

		if rougueValue == 3 {
			rougueValue++
			continue
		}

		if rougueValue == 7 {
			goto elvis
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

elvis:
	fmt.Println("This is the elvis point")

}
