package main

import "fmt"

func pointers() {

	x := 10

	fmt.Println(x)

	xptr := &x

	*xptr = 20

	fmt.Println(x)
}
