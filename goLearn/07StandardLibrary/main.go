package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// libraries

	greetings := "Hello guys, Welcome to go docs"

	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, ","))
	fmt.Println(strings.Split(greetings, " "))

	fmt.Println("Original greetings: ", greetings)

	// sort
	ages := []int{6, 21, 3, 5, 7, 1}
	sort.Ints(ages)
	fmt.Println(ages)

	names := []string{"Errol", "Abhishek", "Jason", "Nanditha"}
	sort.Strings(names)
	fmt.Println(names)

	// search
	search := sort.SearchInts(ages, 3)
	fmt.Println(search)

	search2 := sort.SearchStrings(names, "Abhishek")
	fmt.Println(search2)
}
