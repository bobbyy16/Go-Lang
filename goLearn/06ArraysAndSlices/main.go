package main

import "fmt"

func main() {
	// arrays
	var ages [3]int = [3]int{1, 2, 3}
	fmt.Println(ages, len(ages))

	var ages1 = [3]int{1, 2, 3}
	fmt.Println(ages1, len(ages1))

	ages3 := [3]int{1, 2, 3}
	fmt.Println(ages3, len(ages3))
	ages3[2] = 4
	fmt.Println(ages3, len(ages3))

	// slices
	var names = []string{"Abhishek", "Jason", "Errol"}
	fmt.Println(names, len(names))
	names = append(names, "Nanditha")
	fmt.Println(names, len(names))

	rangeOne := names[2:]
	fmt.Println(rangeOne)

	rangeTwo := ages[:2]
	fmt.Println(rangeTwo)
}
