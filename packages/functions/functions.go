package functions

import "math"

import "fmt"

//Circle struct
type Circle struct {
	Radius, X, Y float64
}

//Area export method
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

//GreaterThan return true if num1 > num2
func GreaterThan(num1 int, num2 int) string {
	if num1 > num2 {
		return "Yes"
	}
	return "No"
}

//Swap start with capital S as the first letter for the name of
//the type, which means this type is exported and accessible by other packages.
func Swap(a, b *string) {
	// return swap(a, b)
	temp := *a
	*a = *b
	*b = temp
}

//GetSequence return anonymous func
//An Example of function closures which are Anonymous Functions that can be used in dynamic programming
func GetSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//this swap func is not accessible outside of the package because starts with lowercase "s"
func swap(a, b string) (string, string) {
	return b, a
}

//ArrayExample Passing array to a function
func ArrayExample(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
