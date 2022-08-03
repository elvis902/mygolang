package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to playing with web request in Golang using our own Node.Js server")
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}
func PerformGetRequest() {
	response, err := http.Get("http://localhost:8000/get")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("The status code is: ", response.StatusCode)
	fmt.Println("The length of the content is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("The bytecount is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"elvis",
			"age": "22",
			"profession":"Software Engineer"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "elvis")
	data.Add("lastname", "rahman")
	data.Add("email", "elvisrahman4@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
