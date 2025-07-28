package main

import "fmt"

// arrays are fixed size groups of variables of the same type
// arrays are zero indexed
// arrays are passed by value
// arrays are mutable

//important if ur size of array or slice is 5 and u want to add more values to it,
// it will make copy of the array and add the values to the copy

//the ... operator is used to pass any number of values to a function
func sum(nums ...int) int {
	var total int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		total += num
	}
	return total
}

func arrayandSlices() {
	//declare array
	// var myInts [10]int //initialize array of 10 ints with default values of 0

	//declare and initialize array
	myInts2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//slices
	prime := [5]int{2, 3, 5, 7, 11}
	mySlice := prime[1:4] //slice from index 1 to 4

	//slices are mutable
	mySlice[1] = 42 //change the value of index 1 to 42
	fmt.Println(mySlice)
	fmt.Println(myInts2)

	//make
	mySlice2 := make([]int, 5, 10) //make a slice of ints with length 5 and capacity 10

	fmt.Println(mySlice2)

	//append
	mySlice2 = append(mySlice2, 1, 2, 3, 4, 5) //append 1, 2, 3, 4, 5 to mySlice2

	fmt.Println(mySlice2)

	//bubild in length function that returns the length of the slice
	fmt.Println(len(mySlice2))

	//built in capacity function that returns the capacity of the slice
	fmt.Println(cap(mySlice2))

	total := sum(1, 2, 3)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4, 5}
	//numbers ... is a spread operator that spreads the values of the slice into the function
	total = sum(numbers...)
	fmt.Println(total)

	//range key word
	for i, num := range numbers {
		fmt.Println(i, num)
	}
}
