package main

import "fmt"

const JwtToken = "sdjflkdjskfjldks" //It is public because name starts with a capital letter

func main() {
	var username string = "elvis"
	fmt.Println(username)
	fmt.Printf("The type of the variable is : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of the variable is : %T \n", isLoggedIn)

	var totalUsers uint8 = 255
	fmt.Println(totalUsers)
	fmt.Printf("The type of the variable is : %T \n", totalUsers)

	var totalValue int32 = -255
	fmt.Println(totalValue)
	fmt.Printf("The type of the variable is : %T \n", totalValue)

	var approxValue float64 = 3444.0943243435434435
	fmt.Println(approxValue)
	fmt.Printf("The type of the variable is : %T \n", approxValue)

	//implicit type
	var myPhone = "iPhone"
	fmt.Println(myPhone)
	fmt.Printf("The type of the variable is : %T \n", myPhone)

	//no vor style
	myPhoneNumber := 25454654
	fmt.Println(myPhoneNumber)
	fmt.Printf("The type of the variable is : %T \n", myPhoneNumber)

	myPhoneNumber = 76788565
	fmt.Println(myPhoneNumber)
	fmt.Printf("The type of the variable is : %T \n", myPhoneNumber)

	fmt.Println(JwtToken)
	fmt.Printf("The type of the variable is : %T \n", JwtToken)
}
