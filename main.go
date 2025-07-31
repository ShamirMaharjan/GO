package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("hello world")

	// variable()
	// function(12, 13)

	// structure()
	// r := rectangle{
	// 	length: 10,
	// 	width:  10,
	// }

	// interfaceShape(r)
	// errorHandling()
	// loop()
	// arrayandSlices()
	// maps()

	// fmt.Println(aggregate(2, 3, 4, add))

	pointers()

	fmt.Println("Starting...")

	// Call printNumbers normally (synchronously)
	// printNumbers("Main")

	// Start a goroutine
	go printNumbers("Goroutine-1")

	// Start another goroutine
	go printNumbers("Goroutine-2")

	// Call the function normally in the main goroutine
	printNumbers("Main")

	// Important: Give goroutines time to finish before main exits
	// Otherwise, the program might end before they print anything.
	time.Sleep(2 * time.Second) // Sleep longer than the goroutines take
	fmt.Println("Done!")
}
