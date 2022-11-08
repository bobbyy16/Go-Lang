package main

import (
	"fmt"
)

func main() {
	// Strings
	var username string = "Abhishek"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	// Boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	// small
	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	// small float
	var smallFloat float32 = 255.45557765675858
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T \n", anothervariable)

	// sprintf
	age := 35
	name := "Abhishek"
	var sprintf = fmt.Sprintf("age is %v , name is %v", age, name)
	fmt.Println("Saved string is:", sprintf)

}
