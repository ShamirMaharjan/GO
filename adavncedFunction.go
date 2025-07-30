package main

// first class function is a function that can be treated as any other value/variable
func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

//higher order function is a function that takes a function as an argument or returns a function as a return value.
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

//defer keyword
//it allows  a function to be executed automatically just before its enclosing function returns, like you are copying a file
//after copying the file you want to close the file
