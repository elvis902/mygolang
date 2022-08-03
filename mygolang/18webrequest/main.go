package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web request using Golang")
	response, err := http.Get("https://reqres.in/api/products/3")

	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of the response is: %T\n", response.Body)
	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
