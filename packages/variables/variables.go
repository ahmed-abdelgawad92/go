package variables

import "fmt"

//Display variables
func Display() {
	var x, y, z int

	x = 10
	y = 20
	z = x * y

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var s string
	s = "Ahmed Ayman"
	fmt.Println(s)

	check := false
	fmt.Println(check)

	fmt.Printf("s is of type %T\n", s)

	var a, b, c int = 1, 2, 3

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	//constants
	const ASD int = 10
	fmt.Println(ASD)

}
