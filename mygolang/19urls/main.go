package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://elvisrahman.com:1000/about?fname=elvis&hobby=badminton"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println("My url is: ", myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Printf("The type returned by url.Parse is : %T\n", result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of qparams is: %T\n", qparams)
	//so qparam is acting as an key value pair
	fmt.Println(qparams["fname"])

	for _, val := range qparams {
		fmt.Println("The value of param is: ", val)
	}

	//creating an url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "elvisbusiness.com:2000",
		Path:     "listall",
		RawQuery: "scale=3000&location=india",
	}

	myNewUrl := partsOfUrl.String()
	fmt.Println(myNewUrl)
}
