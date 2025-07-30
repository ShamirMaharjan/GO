package main

import "fmt"

func maps() {
	ages := make(map[string]int) //map of string keys to int values
	ages["alice"] = 31           //insert or update a key value pair
	ages["charlie"] = 34
	fmt.Println(ages["alice"]) // "31"
	fmt.Println(ages["bob"])   // "0"

	fmt.Println(len(ages)) // "2"

	getValue := ages["alice"]

	fmt.Println(getValue) // to get the value of a key

	delete(ages, "alice")
	fmt.Println(ages)

	//checking if a key is present
	elem, ok := ages["alice"]
	fmt.Println(elem, ok) // "0 false"
}
