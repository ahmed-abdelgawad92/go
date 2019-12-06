package main

import (
	"fmt"
	"packages/arrays"
	"packages/functions"
	"packages/variables"
	"strings"
)

func main() {
	/* HELLO WORLD PROGRAM*/
	//Hello World
	fmt.Println("Hello World")

	//call the function variables
	variables.Display()

	//functions
	fmt.Println(functions.GreaterThan(10, 20))

	var a, b string = "Ahmed", "Abdelgawad"
	fmt.Println("a = "+a, "b = "+b)

	functions.Swap(&a, &b)

	fmt.Println("a = "+a, "b = "+b)

	fmt.Printf("a = %s\n", a)

	count := functions.GetSequence()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	count2 := functions.GetSequence()

	fmt.Println(count())
	fmt.Println(count2())
	fmt.Println(count2())
	fmt.Println(count2())
	fmt.Println(count2())

	var circle functions.Circle
	circle.X = 0
	circle.Y = 0
	circle.Radius = 5
	fmt.Println(circle.Area())

	//Array And String
	// var variable_name [SIZE] variable_type
	var numbers = []int{11, 22, 33, 44, 55}
	functions.ArrayExample(numbers)

	var newstrings = []string{"Ahmed", "Ayman", "Mokhtar"}
	fmt.Println(strings.Join(newstrings, " "))

	//slices
	arrays.Slices()

	//Range
	arrays.Range()

	//maps
	arrays.Map()
}
