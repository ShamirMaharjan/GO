package main

import "fmt"

func variable() {
	//int, int8, int16, int32, int64  - whole numbers
	//uint, uint8, uint16, uint32, uint64 - whole numbers without negative values
	//float32, float64 - decimal numbers
	//bool - true or false
	//string - text
	//byte - single character
	//rune - single character
	//complex64, complex128 - complex numbers

	//there are 2 ways to declare a variable
	// 1. var a int
	// 2. a := 10

	a := 10

	//can declar multiple variables in one line
	one, two := 1, 2

	fmt.Println(a)

	fmt.Println(one, two)

	//can convert a variable to another type
	isint := 2

	isfloat := float64(isint)

	fmt.Println(isfloat)

	//const - constant variable

	const pi = 3.14
	fmt.Println(pi)

	//conditional statements

	if a > 10 {
		fmt.Println("a is greater than 10")
	} else {
		fmt.Println("a is less than 10")
	}

}
