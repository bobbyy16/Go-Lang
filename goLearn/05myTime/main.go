package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create a date
	createdDate := time.Date(2008, time.December, 16, 8, 37, 55, 00, time.UTC)
	fmt.Println(createdDate)
}
