package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Println(x)
		x++
	}
	fmt.Println("===============================")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("===============================")

	names := []string{"Abhishek", "Errol", "Jason", "Nans"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println("===============================")

	for index, value := range names {
		fmt.Printf("The value of index %v is %v \n", index, value)
	}

	fmt.Println("===============================")

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
}
