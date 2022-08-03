package main

import "fmt"

func main() {
	fmt.Println("Structs for golang")

	//structs in golang donot support inheritance
	elvis := User{"Elvis", "elvis@go.dev", 22, true}
	fmt.Println("The details of the new user is ", elvis)
	fmt.Printf("The details of the new user in a good formated way is %+v\n", elvis)
	fmt.Printf("The name of the user is %v and the age of the user is %v\n", elvis.Name, elvis.Age)

	elvis.GetStatus()
	elvis.NewMain()
	fmt.Printf("The email of the user is %v\n", elvis.Email)

}

//If we wana export any properties from struct then they must start with capital letter
type User struct {
	Name    string
	Email   string
	Age     int
	IsValid bool
}

func (u User) GetStatus() {
	fmt.Println("Is the user valid: ", u.IsValid)
}

// The user is passed as a copy in the argument if we donot use pointer
func (u User) NewMain() {
	u.Email = "test@go.dev"
	fmt.Println("The new mail is: ", u.Email)
}
