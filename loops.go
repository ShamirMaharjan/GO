package main

// initial; condition; after
import "fmt"

//go allows for the omission of a for loop, a while loop is just a for loop without the initial and after statements

func loop() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	height := 1

	for height < 5 {
		fmt.Println(height)
		height++
	}
}
