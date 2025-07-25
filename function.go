package main

import "fmt"

// important

//in go variable are passed by vale not by reference
//if you want to pass by reference then you have to pass the address of the variable
//and you have to use the * operator to get the value of the variable
//suppose x:=2 and y:=x then y=2 but if you want to pass by reference then you have to pass the address of x so basically y is copy of x
//and if you change the value of any for eg y=1 then x will not change because x is copy of y

//this is a return tyoe function here if the parameters are same then we can write them once
//if the parameters are different then we have to write them all
func function(x, y int) int {
	z := x + y
	fmt.Println(z)
	return z
}

// if there are more than one return type we wrap them in ()
func getName() (string, string) {
	return "John", "Doe"
}

//if you want to ignore a return type then you can use _
// x, _ := getName()
