package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to working with file in Golang")
	file, err := os.Create("./mytextfile.txt")

	if err != nil {
		panic(err)
	}

	content := "This is the content to be inserted in the text file - elvis.rahman"

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("The length is: ", length)
	defer file.Close()

	readFile("./mytextfile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("The content of the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
