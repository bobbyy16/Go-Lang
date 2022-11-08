package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Println("Good Morning!", n)
}

func sayBye(n string) {
	fmt.Println("Good Night!", n)
}
func cycleNames(number []string, function func(string)) {
	for _, value := range number {
		function(value)
	}
}

// return type

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreetings("Abhishek")
	sayBye("Abhishek")

	cycleNames([]string{"Naruto", "Hinata", "Boruto", "Himawari"}, sayGreetings)
	cycleNames([]string{"Naruto", "Hinata", "Boruto", "Himawari"}, sayBye)

	Area := circleArea(22)
	fmt.Println(Area)
}
