package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

	// integers
	var one int8 = 5
	fmt.Println(one)

	var two int16 = 10
	fmt.Println(two)

	var three int32 = 15
	fmt.Println(three)

	var four int64 = 20
	fmt.Println(four)

	var five uint8 = 25
	fmt.Println(five)

	var six uint16 = 30
	fmt.Println(six)

	var seven uint32 = 35
	fmt.Println(seven)

	var eight uint64 = 40
	fmt.Println(eight)

	// type conversion
	var xx int16 = 5
	var yy int8 = 2
	// x = y; would not work
	xx = int16(yy)
	fmt.Println(xx)

	// the other way around
	xx = 10000
	yy = int8(xx)
	fmt.Println(yy) // works but gives wrong output

	// floats
	var float1 float32 = 1.5
	fmt.Println(float1)

	var float2 float64 = 2.5e2 // e2 == 10 to the 2
	fmt.Println(float2)

	// complex numbers
	var complex1 complex128 = complex(2, 3)
	fmt.Println(complex1)

	// string
	var string1 string = "some string"
	var string2 string = "some strinh"
	fmt.Println(string1)

	fmt.Println(strings.Compare(string1, string2))
	fmt.Println(strings.ToUpper(string1))

	var numberString string = "123.5"
	y2, e2 := strconv.Atoi(numberString)

	fmt.Println(y2)
	fmt.Println(e2)

	// constants
	const myConst int = 5
	fmt.Println(myConst)

	// enumaration
	type Grade int
	const (
		AA Grade = iota
		BB
		CC
		DD
		FF
	)

	var someGrade Grade = BB

	fmt.Println(someGrade)

	// if statement
	var ifStatementVar int = 5
	if ifStatementVar == 5 {
		fmt.Println(" it is 5")
	} else {
		fmt.Println("it is not 5")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("loop is working")
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	/*for {
		fmt.Println("hi")
	}*/

	// switch case
	var someOtherGrade Grade = CC
	// automatically breaks without a "break" statement
	switch someOtherGrade {
	case AA:
		fmt.Println("it is A")
	case BB:
		fmt.Println("it is B")
	case CC:
		fmt.Println("it is C")
	default:
		fmt.Println("default")

	}

	// tagless switch
	// have conditions
	switch {
	case someOtherGrade == AA:
		fmt.Println("it is A")
	case someOtherGrade == BB:
		fmt.Println("it is B")
	case someOtherGrade == CC:
		fmt.Println("it is C")
	default:
		fmt.Println("default")

	}

	// scanning user input
	var number int
	fmt.Println("enter a number...")
	fmt.Scan(&number)
	fmt.Println("you entered: ", number)

}

func someFunc() {

	var outVar int = 2000
	fmt.Println(outVar)
}
