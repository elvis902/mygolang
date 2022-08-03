package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to mytime with Golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createDate := time.Date(2018, time.November, 13, 04, 16, 07, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("Monday 02-01-2006"))
}
