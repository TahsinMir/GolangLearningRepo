package main

import "fmt"

var outVar int = 1000

func main() {
	// printing
	fmt.Println("okay google")

	// variables and types
	var x int = 5
	fmt.Println(x)

	var someString = "test"
	fmt.Println(someString)

	// making alias of existing types
	type Celcius float64
	var myTemp Celcius = 42.345
	fmt.Println(myTemp)

	// short variable declaration
	y := 100

	// pointer, & returns address of a var, * returns the value of a pointer
	fmt.Println(&y)
	fmt.Println(*&y)

	// new() returns a pointer to a variable
	ptr := new(int)
	*ptr = 3
	fmt.Println(*ptr)

	someFunc()
}

func someFunc() {

	var outVar int = 2000
	fmt.Println(outVar)
}
